package edits

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
			result := breakdownToCharacters(c.input)
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
			result := breakdownToWords(c.input)
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
			result := breakdownToWholeLine(c.input)
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
			BREAKDOWN_TO_WORDS,
			[]string{"import ", "Editor ", "from ", "\"@monaco-editor/react\";"},
		},
		{
			"", //even if it's an empty string, we don't care, just return what's given as it has no "\n|
			BREAKDOWN_TO_CHARACTERS,
			[]string{""},
		},
		{
			`		return nil, fmt.Errorf("failed in gitFilesForCommit, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", commitHashStr, commitHash.String())`,
			BREAKDOWN_TO_WHOLE_LINE,
			[]string{`		return nil, fmt.Errorf("failed in gitFilesForCommit, commit hash = %s is invalid as its re-calculated hash is mismatched = %s", commitHashStr, commitHash.String())`},
		},
	}

	for index, c := range cases {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			result := breakdownLineToAdd(c.input)
			if !cmp.Equal(c.expected, result) {
				t.Errorf("expected: %s", c.expected)
				t.Errorf("result  : %s", result)
				t.Fatalf(cmp.Diff(c.expected, result))
			}
		})
	}
}
