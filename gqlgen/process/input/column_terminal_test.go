package input

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/testio"
)

func TestToTerminalRow(t *testing.T) {
	cases := []struct {
		name       string
		inputFile  string
		goldenFile string
	}{
		{"command with tooltip" /*****/, "testdata/terminal/cmd1.json", "testdata/terminal/cmd1-golden.json"},
		{"command without tooltip" /**/, "testdata/terminal/cmd2.json", "testdata/terminal/cmd2-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var abst Row
			err := jsonwrap.Read(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			result, err := toTerminalCommandRow(&abst)
			if err != nil {
				t.Fatal(err)
			}

			testio.CompareWithGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestToTerminalOutputRow(t *testing.T) {
	cases := []struct {
		name       string
		inputFile  string
		goldenFile string
	}{
		{"output without tooltip" /***/, "testdata/terminal/output1.json", "testdata/terminal/output1-golden.json"},
		{"output with tooltip" /******/, "testdata/terminal/output2.json", "testdata/terminal/output2-golden.json"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var abst Row
			err := jsonwrap.Read(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			result, err := toTerminalOutputRow(&abst)
			if err != nil {
				t.Fatal(err)
			}

			testio.CompareWithGoldenFile(t, *updateFlag, c.goldenFile, result)
		})
	}
}

func TestToTerminalCommandRowError(t *testing.T) {
	cases := []struct {
		name      string
		inputFile string
	}{
		{"empty contents" /********/, "testdata/terminal/cmd-error1.json"},
		{"tooltip timing is wrong" /**/, "testdata/terminal/cmd-error2.json"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var abst Row
			err := jsonwrap.Read(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			_, err = toTerminalCommandRow(&abst)
			if err == nil {
				t.Fatal("expected to fail but succeeded")
			}
		})
	}
}

func TestToTerminalOutputRowError(t *testing.T) {
	cases := []struct {
		name      string
		inputFile string
	}{
		{"empty contents" /********/, "testdata/terminal/output-error1.json"},
		{"tooltip timing is wrong" /**/, "testdata/terminal/output-error2.json"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var abst Row
			err := jsonwrap.Read(c.inputFile, &abst)
			if err != nil {
				t.Fatal(err)
			}

			_, err = toTerminalOutputRow(&abst)
			if err == nil {
				t.Fatal("expected to fail but succeeded")
			}
		})
	}
}
