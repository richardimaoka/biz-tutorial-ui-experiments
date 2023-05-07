package processing_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

// TODO: rewrite tests with table driven tests
// TODO: private-ise the methods exept Transition
// TODO: move the test to processing_test
//
// type TerminalEffect struct {
// 	SeqNo            int     `json:"seqNo"`
// 	TerminalName     string  `json:"terminalName"`
// 	Command          string  `json:"command"`
// 	Output           *string `json:"output"`           //if zero value, no output after execution
// 	CurrentDirectory *string `json:"currentDirectory"` //if zero value, current directory is not changed after execution
// }
// var cases := []TerminalEffect{
// 	{ SeqNo: 0, TerminalName: "default", Command: "echo abc", Output: nil, CurrentDirectory: nil },
// }

// test case for TerminalProcessor's WriteCommand method

func Test_Terminal(t *testing.T) {
	type TestCase struct {
		ExpectedFile   string
		TerminalEffect processing.TerminalEffect
	}

	testCases := []TestCase{
		{"testdata/terminal/terminal1-1.json", processing.TerminalEffect{SeqNo: 0, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: nil}},
		// {"testdata/terminal/terminal1-2.json", processing.TerminalEffect{SeqNo: 1, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: internal.Address("def")}},
		// {"testdata/terminal/terminal1-3.json", processing.TerminalEffect{SeqNo: 2, TerminalName: "default", Command: "echo abc", Output: internal.Address("abc"), CurrentDirectory: nil}},
		// {"testdata/terminal/terminal1-4.json", processing.TerminalEffect{SeqNo: 3, TerminalName: "default", Command: "mkdir def", Output: nil, CurrentDirectory: nil}},
		// {"testdata/terminal/terminal1-5.json", processing.TerminalEffect{SeqNo: 4, TerminalName: "default", Command: "mkdir abc", Output: nil, CurrentDirectory: nil}},
	}

	terminal := processing.NewTerminalProcessor("default")
	for _, c := range testCases {
		t.Run("tt.descriptio", func(t *testing.T) {
			step := fmt.Sprintf("%03d", c.TerminalEffect.SeqNo)
			terminal.Transition(c.TerminalEffect, step)

			if *update {
				internal.WriteGoldenFile(t, c.ExpectedFile, terminal.ToGraphQLTerminal())
			}

			internal.CompareAfterMarshal(t, c.ExpectedFile, terminal.ToGraphQLTerminal())
		})
	}
}
