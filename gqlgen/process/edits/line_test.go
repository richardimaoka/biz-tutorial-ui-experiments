package edits

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

func TestSplitAfterNewLine(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"import Editor from \"@monaco-editor/react\";\n",
			[]string{"import Editor from \"@monaco-editor/react\";\n"},
		},
		{
			"  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;\n  // pass-in a callback like below to manipulate editor instance\n",
			[]string{
				"  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;\n",
				"  // pass-in a callback like below to manipulate editor instance\n",
			},
		},
		{
			"",         // if it happends to be an empty change, ...
			[]string{}, // then it's safe to omit
		},
		{
			"\n",
			[]string{"\n"},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := splitAfterNewLine(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %s", c.expected)
				t.Errorf("result  : %s", result)
				// t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}

func TestDetectNewLines(t *testing.T) {
	type Expectation struct {
		HasNewLetter          bool
		ContentWithoutNewLine string
	}

	cases := []struct {
		input    string
		expected SingleLineToAdd
	}{
		{
			"import Editor from \"@monaco-editor/react\";\n",
			SingleLineToAdd{
				true,
				"import Editor from \"@monaco-editor/react\";",
			},
		},
		{
			"  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;\n",
			SingleLineToAdd{
				true,
				"  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;",
			},
		},
		{
			"  // pass-in a callback like below to manipulate editor instance\n",
			SingleLineToAdd{
				true,
				"  // pass-in a callback like below to manipulate editor instance",
			},
		},
		{
			" some word vvvv",
			SingleLineToAdd{
				false,
				" some word vvvv",
			},
		}, {
			"",
			SingleLineToAdd{
				false,
				"",
			},
		},
		{
			// if only "\n", then it 1) has '\n' as indicated by `true`, but 2) the content without '\n' is empty string
			"\n",
			SingleLineToAdd{
				true,
				"",
			},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := detectNewLine(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %t, '%s'", c.expected.NewLineAtEnd, c.expected.ContentWithoutNewLine)
				t.Errorf("result  : %t, '%s'", result.NewLineAtEnd, result.ContentWithoutNewLine)
			}
		})
	}
}

func TestSplitIntoSingleLines(t *testing.T) {
	cases := []struct {
		input    internal.Chunk
		expected []SingleLineToAdd
	}{
		{
			input: internal.Chunk{
				Content: "  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;\n  // pass-in a callback like below to manipulate editor instance\n",
				Type:    "Add",
			},
			expected: []SingleLineToAdd{
				{
					true,
					"  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;",
				},
				{
					true,
					"  // pass-in a callback like below to manipulate editor instance",
				},
			},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := splitChunkToLines(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}

func TestLineToPosChunks(t *testing.T) {
	cases := []struct {
		inputPos    TypingPosition
		inputLine   SingleLineToAdd
		expected    []ChunkToAdd
		expectedPos TypingPosition
	}{
		{
			TypingPosition{LineNumber: 1, Column: 1},
			SingleLineToAdd{
				ContentWithoutNewLine: "import Editor from \"@monaco-editor/react\";",
				NewLineAtEnd:          false,
			},
			[]ChunkToAdd{
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "import "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 8}, Type: "Add", Content: "Editor "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 15}, Type: "Add", Content: "from "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 20}, Type: "Add", Content: "\"@monaco-editor/react\";"},
			},
			TypingPosition{LineNumber: 1, Column: 43},
		},
		{
			TypingPosition{LineNumber: 1, Column: 1},
			SingleLineToAdd{
				ContentWithoutNewLine: "import Editor from \"@monaco-editor/react\";",
				NewLineAtEnd:          true, // + '\n' to the above content
			},
			[]ChunkToAdd{
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "\n"},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "import "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 8}, Type: "Add", Content: "Editor "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 15}, Type: "Add", Content: "from "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 20}, Type: "Add", Content: "\"@monaco-editor/react\";"},
			},
			TypingPosition{LineNumber: 2, Column: 1},
		},
		{
			TypingPosition{1, 1},
			SingleLineToAdd{
				ContentWithoutNewLine: "",
				NewLineAtEnd:          true, // '\n' only
			},
			[]ChunkToAdd{
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "\n"},
			},
			TypingPosition{LineNumber: 2, Column: 1},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			resultPos, result := lineToPosChunks(c.inputLine, c.inputPos)
			if !cmp.Equal(c.expected, result) {
				t.Errorf(cmp.Diff(c.expected, result))
			}
			if resultPos != c.expectedPos {
				t.Errorf("expected pos: %v", c.expectedPos)
				t.Errorf("result pos  : %v", resultPos)
			}
		})
	}
}

func TestToPositionedChunks(t *testing.T) {
	cases := []struct {
		inputPos    TypingPosition
		inputChunk  internal.Chunk
		expected    []ChunkToAdd
		expectedPos TypingPosition
	}{
		{
			TypingPosition{LineNumber: 1, Column: 1},
			internal.Chunk{
				Content: "import { editor } from \"monaco-editor\";\n",
				Type:    "Add",
			},
			[]ChunkToAdd{
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "\n"},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "import "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 8}, Type: "Add", Content: "{ "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 10}, Type: "Add", Content: "editor "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 17}, Type: "Add", Content: "} "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 19}, Type: "Add", Content: "from "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 24}, Type: "Add", Content: "\"monaco-editor\";"},
			},
			TypingPosition{LineNumber: 2, Column: 1},
		},
		{
			TypingPosition{LineNumber: 1, Column: 1},
			internal.Chunk{
				Content: "import { editor } from \"monaco-editor\";\n\ninterface Props {\n",
				Type:    "Add",
			},
			[]ChunkToAdd{
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "\n"},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "import "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 8}, Type: "Add", Content: "{ "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 10}, Type: "Add", Content: "editor "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 17}, Type: "Add", Content: "} "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 19}, Type: "Add", Content: "from "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 24}, Type: "Add", Content: "\"monaco-editor\";"},
				{TypingPosition: TypingPosition{LineNumber: 2, Column: 1}, Type: "Add", Content: "\n"},
				{TypingPosition: TypingPosition{LineNumber: 3, Column: 1}, Type: "Add", Content: "\n"},
				{TypingPosition: TypingPosition{LineNumber: 3, Column: 1}, Type: "Add", Content: "interface "},
				{TypingPosition: TypingPosition{LineNumber: 3, Column: 11}, Type: "Add", Content: "Props "},
				{TypingPosition: TypingPosition{LineNumber: 3, Column: 17}, Type: "Add", Content: "{"},
			},
			TypingPosition{LineNumber: 4, Column: 1},
		},
		{
			TypingPosition{LineNumber: 1, Column: 1},
			internal.Chunk{
				Content: "import { editor } from \"monaco-editor\";\n\ninterface Props {",
				Type:    "Add",
			},
			[]ChunkToAdd{
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "\n"},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 1}, Type: "Add", Content: "import "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 8}, Type: "Add", Content: "{ "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 10}, Type: "Add", Content: "editor "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 17}, Type: "Add", Content: "} "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 19}, Type: "Add", Content: "from "},
				{TypingPosition: TypingPosition{LineNumber: 1, Column: 24}, Type: "Add", Content: "\"monaco-editor\";"},
				{TypingPosition: TypingPosition{LineNumber: 2, Column: 1}, Type: "Add", Content: "\n"},
				{TypingPosition: TypingPosition{LineNumber: 3, Column: 1}, Type: "Add", Content: "interface "},
				{TypingPosition: TypingPosition{LineNumber: 3, Column: 11}, Type: "Add", Content: "Props "},
				{TypingPosition: TypingPosition{LineNumber: 3, Column: 17}, Type: "Add", Content: "{"},
			},
			TypingPosition{LineNumber: 3, Column: 18},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			resultPos, result := toPositionedChunks(c.inputChunk, c.inputPos)
			if !cmp.Equal(c.expected, result) {
				t.Errorf(cmp.Diff(c.expected, result))
			}
			if resultPos != c.expectedPos {
				t.Errorf(cmp.Diff(c.expectedPos, resultPos))
			}
		})
	}
}
