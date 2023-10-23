package edits

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplitSingleLineAdd(t *testing.T) {
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
			result := splitSingleLineAdd(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %s", c.expected)
				t.Errorf("result  : %s", result)
				// t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}

func TestMovesNewLineToHead(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"import Editor from \"@monaco-editor/react\";\n",
			[]string{"\n", "import Editor from \"@monaco-editor/react\";"},
		},
		{
			"  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;\n",
			[]string{
				"\n",
				"  onDidMount?: (editorInstance: editor.IStandaloneCodeEditor) =\u003e void;",
			},
		},
		{
			"  // pass-in a callback like below to manipulate editor instance\n",
			[]string{
				"\n",
				"  // pass-in a callback like below to manipulate editor instance",
			},
		},
		{
			"", //even if it's an empty string, we don't care, just return what's given as it has no "\n|
			[]string{""},
		},
		{
			// if only "\n", then only return "\n"
			"\n",
			[]string{"\n"},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := moveNewLineToHead(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %s", c.expected)
				t.Errorf("result  : %s", result)
				t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}

func TestOneLetterAdditions(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"import Editor from \"@monaco-editor/react\";",
			[]string{"i", "m", "p", "o", "r", "t", " ", "E", "d", "i", "t", "o", "r", " ", "f", "r", "o", "m", " ", "\"", "@", "m", "o", "n", "a", "c", "o", "-", "e", "d", "i", "t", "o", "r", "/", "r", "e", "a", "c", "t", `"`, ";"},
		},
		{
			"", //even if it's an empty string, we don't care, just return what's given as it has no "\n|
			[]string{""},
		},
		{
			// if only "\n", then only return "\n"
			"\n",
			[]string{"\n"},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := eachCharacterAdditions(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %s", c.expected)
				t.Errorf("result  : %s", result)
				t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}

func TestWordByWordAdditions(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"import Editor from \"@monaco-editor/react\";",
			[]string{"import ", "Editor ", "from ", "\"@monaco-editor/react\";"},
		},
		{
			"", //even if it's an empty string, we don't care, just return what's given as it has no "\n|
			[]string{""},
		},
		{
			// if only "\n", then only return "\n"
			"\n",
			[]string{"\n"},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := wordByWordAdditions(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %s", c.expected)
				t.Errorf("result  : %s", result)
				t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}

func TestWholeLineAddition(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			"import Editor from \"@monaco-editor/react\";",
			[]string{"import Editor from \"@monaco-editor/react\";"},
		},
		{
			"", //even if it's an empty string, we don't care, just return what's given as it has no "\n|
			[]string{""},
		},
		{
			// if only "\n", then only return "\n"
			"\n",
			[]string{"\n"},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := wholeLineAddition(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %s", c.expected)
				t.Errorf("result  : %s", result)
				t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}

func TestConditionalAdditions(t *testing.T) {
	cases := []struct {
		input        string
		additionType string
		expected     []string
	}{
		{
			"import Editor from \"@monaco-editor/react\";",
			WORD_BY_WORD_ADDITIONS,
			[]string{"import ", "Editor ", "from ", "\"@monaco-editor/react\";"},
		},
		{
			"", //even if it's an empty string, we don't care, just return what's given as it has no "\n|
			EACH_CHARACTER_ADDITIONS,
			[]string{""},
		},
		{
			`		return nil, fmt.Errorf("failed in gitFilesForCommit, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", commitHashStr, commitHash.String())`,
			WHOLE_LINE_ADDITIONS,
			[]string{`		return nil, fmt.Errorf("failed in gitFilesForCommit, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", commitHashStr, commitHash.String())`},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := conditionalAdditions(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %s", c.expected)
				t.Errorf("result  : %s", result)
				t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}
