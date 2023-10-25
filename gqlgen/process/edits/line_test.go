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
			result := chunkToLines(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}

func TestLineToPosChunks(t *testing.T) {
	cases := []struct {
		inputLineNumber int
		inputColumn     int
		inputLine       SingleLineToAdd
		expected        []PositionedChunk
	}{
		{
			0, 0, SingleLineToAdd{
				ContentWithoutNewLine: "import Editor from \"@monaco-editor/react\";",
				NewLineAtEnd:          true,
			},
			[]PositionedChunk{
				{LineNumber: 0, Column: 0, Type: "Add", Content: "\n"},
				{LineNumber: 0, Column: 0, Type: "Add", Content: "import "},
				{LineNumber: 0, Column: 7, Type: "Add", Content: "Editor "},
				{LineNumber: 0, Column: 14, Type: "Add", Content: "from "},
				{LineNumber: 0, Column: 19, Type: "Add", Content: "\"@monaco-editor/react\";"},
			},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := lineToPosChunks(c.inputLine, c.inputLineNumber, c.inputColumn)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %v", c.expected)
				t.Errorf("result  : %v", result)
				t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}
