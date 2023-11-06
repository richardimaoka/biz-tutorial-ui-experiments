package edits

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
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
		expected SingleLineChange
	}{
		{
			"import Editor from \"@monaco-editor/react\";\n",
			SingleLineChange{
				true,
				"import Editor from \"@monaco-editor/react\";",
			},
		},
		{
			"  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;\n",
			SingleLineChange{
				true,
				"  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;",
			},
		},
		{
			"  // pass-in a callback like below to manipulate editor instance\n",
			SingleLineChange{
				true,
				"  // pass-in a callback like below to manipulate editor instance",
			},
		},
		{
			" some word vvvv",
			SingleLineChange{
				false,
				" some word vvvv",
			},
		}, {
			"",
			SingleLineChange{
				false,
				"",
			},
		},
		{
			// if only "\n", then it 1) has '\n' as indicated by `true`, but 2) the content without '\n' is empty string
			"\n",
			SingleLineChange{
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

func TestSplitChunkToLines(t *testing.T) {
	cases := []struct {
		input    gitwrap.Chunk
		expected []SingleLineChange
	}{
		{
			input: gitwrap.Chunk{
				Content: "  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;\n  // pass-in a callback like below to manipulate editor instance\n",
				Type:    "Add",
			},
			expected: []SingleLineChange{
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

func TestLineToChunksToAdd(t *testing.T) {
	cases := []struct {
		inputPos    TypingPosition
		inputLine   SingleLineChange
		expected    []ChunkToAdd
		expectedPos TypingPosition
	}{
		{
			TypingPosition{LineNumber: 1, Column: 1},
			SingleLineChange{
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
			SingleLineChange{
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
			SingleLineChange{
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
			resultPos, result := lineToChunksToAdd(c.inputLine, c.inputPos)
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

func TestLineToChunksToDelete(t *testing.T) {
	cases := []struct {
		inputPos    TypingPosition
		inputLine   SingleLineChange
		expected    []ChunkToDelete
		expectedPos TypingPosition
	}{
		{
			TypingPosition{LineNumber: 1, Column: 1},
			SingleLineChange{
				ContentWithoutNewLine: "import Editor from \"@monaco-editor/react\";",
				NewLineAtEnd:          true,
			},
			[]ChunkToDelete{
				{
					Content:       "import Editor from \"@monaco-editor/react\";\n",
					RangeToDelete: RangeToDelete{StartLineNumber: 1, EndLineNumber: 2, StartColumn: 1, EndColumn: 0},
				},
			},
			TypingPosition{LineNumber: 1, Column: 1},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			resultPos, result := lineToChunksToDelete(c.inputLine, c.inputPos)
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

func TestToChunksToAdd(t *testing.T) {
	cases := []struct {
		inputPos    TypingPosition
		inputChunk  gitwrap.Chunk
		expected    []ChunkToAdd
		expectedPos TypingPosition
	}{
		{
			TypingPosition{LineNumber: 1, Column: 1},
			gitwrap.Chunk{
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
			gitwrap.Chunk{
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
			gitwrap.Chunk{
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
			resultPos, result := toChunksToAdd(c.inputChunk, c.inputPos)
			if !cmp.Equal(c.expected, result) {
				t.Errorf(cmp.Diff(c.expected, result))
			}
			if resultPos != c.expectedPos {
				t.Errorf(cmp.Diff(c.expectedPos, resultPos))
			}
		})
	}
}

func TestToChunksToDelete(t *testing.T) {
	cases := []struct {
		inputPos   TypingPosition
		inputChunk gitwrap.Chunk
		expected   []ChunkToDelete
	}{
		{
			TypingPosition{LineNumber: 1, Column: 1},
			gitwrap.Chunk{
				Content: "import Editor, { OnChange } from \"@monaco-editor/react\";",
				Type:    "Delete",
			},
			[]ChunkToDelete{
				{
					Content:       "import Editor, { OnChange } from \"@monaco-editor/react\";",
					RangeToDelete: RangeToDelete{StartLineNumber: 1, EndLineNumber: 1, StartColumn: 1, EndColumn: 56},
				},
			},
		},
		{
			TypingPosition{LineNumber: 1, Column: 1},
			gitwrap.Chunk{
				Content: "import Editor, { OnChange } from \"@monaco-editor/react\";\n",
				Type:    "Delete",
			},
			[]ChunkToDelete{
				{
					Content:       "import Editor, { OnChange } from \"@monaco-editor/react\";\n",
					RangeToDelete: RangeToDelete{StartLineNumber: 1, EndLineNumber: 2, StartColumn: 1, EndColumn: 0},
				},
			},
		},
		{
			TypingPosition{LineNumber: 1, Column: 1},
			gitwrap.Chunk{
				Content: "import { editor } from \"monaco-editor\";\n\ninterface Props {\n",
				Type:    "Delete",
			},
			[]ChunkToDelete{
				{
					Content:       "import { editor } from \"monaco-editor\";\n",
					RangeToDelete: RangeToDelete{StartLineNumber: 1, EndLineNumber: 2, StartColumn: 1, EndColumn: 0},
				},
				{
					Content:       "\n",
					RangeToDelete: RangeToDelete{StartLineNumber: 1, EndLineNumber: 2, StartColumn: 1, EndColumn: 0},
				},
				{
					Content:       "interface Props {\n",
					RangeToDelete: RangeToDelete{StartLineNumber: 1, EndLineNumber: 2, StartColumn: 1, EndColumn: 0},
				},
			},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := toChunksToDelete(c.inputChunk, c.inputPos)
			if !cmp.Equal(c.expected, result) {
				t.Errorf(cmp.Diff(c.expected, result))
			}
		})
	}
}

func TestProcessChunk(t *testing.T) {
	cases := []struct {
		inputPos    TypingPosition
		inputChunk  gitwrap.Chunk
		expected    []SingleEditOperation
		expectedPos TypingPosition
	}{
		{
			TypingPosition{LineNumber: 1, Column: 1},
			gitwrap.Chunk{
				Content: "\"use client\";\n\n",
				Type:    "Equal",
			},
			[]SingleEditOperation{},
			TypingPosition{LineNumber: 3, Column: 1},
		},
		{
			TypingPosition{LineNumber: 3, Column: 1},
			gitwrap.Chunk{
				Content: "import Editor from \"@monaco-editor/react\";\n",
				Type:    "Delete",
			},
			[]SingleEditOperation{
				{Range: Range{StartLineNumber: 3, StartColumn: 1, EndLineNumber: 4, EndColumn: 0}, Text: ""},
			},
			TypingPosition{LineNumber: 3, Column: 1},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			resultPos, resultOps := processChunk(c.inputChunk, c.inputPos)
			if !cmp.Equal(c.expected, resultOps) {
				t.Errorf(cmp.Diff(c.expected, resultOps))
			}
			if resultPos != c.expectedPos {
				t.Errorf(cmp.Diff(c.expectedPos, resultPos))
			}
		})
	}
}
