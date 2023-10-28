package input

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

func TestToTerminalCommand(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/cmd1.json", "testdata/cmd1-golden.json"},
		{"testdata/cmd2.json", "testdata/cmd2-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var abst Abstract
			err := internal.JsonRead2(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			result, err := toTerminalCommand(&abst)
			if err != nil {
				t.Fatal(err)
			}

			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestToTerminalCommandError(t *testing.T) {
	cases := []struct {
		inputFile string
	}{
		{"testdata/cmd-error1.json"},
		{"testdata/cmd-error2.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var abst Abstract
			err := internal.JsonRead2(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			_, err = toTerminalCommand(&abst)
			if err == nil {
				t.Fatal("expected to fail but succeeded")
			}
		})
	}
}

func TestToTerminalOutput(t *testing.T) {
	cases := []struct {
		inputFile  string
		goldenFile string
	}{
		{"testdata/output1.json", "testdata/output1-golden.json"},
		{"testdata/output2.json", "testdata/output2-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var abst Abstract
			err := internal.JsonRead2(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			result, err := toTerminalOutput(&abst)
			if err != nil {
				t.Fatal(err)
			}

			internal.CompareWitGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestToTerminalOutputError(t *testing.T) {
	cases := []struct {
		inputFile string
	}{
		// {"testdata/output-error1.json"},
		{"testdata/output-error2.json"},
	}

	for _, c := range cases {
		t.Run(c.inputFile, func(t *testing.T) {
			var abst Abstract
			err := internal.JsonRead2(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			_, err = toTerminalOutput(&abst)
			if err == nil {
				t.Fatal("expected to fail but succeeded")
			}
		})
	}
}
