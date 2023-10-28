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
		{"testdata/terminal/cmd1.json", "testdata/terminal/cmd1-golden.json"},
		{"testdata/terminal/cmd2.json", "testdata/terminal/cmd2-golden.json"},
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
		{"testdata/terminal/cmd-error1.json"},
		{"testdata/terminal/cmd-error2.json"},
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
		{"testdata/terminal/output1.json", "testdata/terminal/output1-golden.json"},
		{"testdata/terminal/output2.json", "testdata/terminal/output2-golden.json"},
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
		{"testdata/terminal/output-error1.json"},
		{"testdata/terminal/output-error2.json"},
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
