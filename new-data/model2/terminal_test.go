package model2

import (
	"fmt"
	"reflect"
	"testing"
)

type TerminalOperation interface {
	IsTerminalOperation()
}

type ChangeDirectory struct {
	FilePath string `json:"filePath"`
}

func (o ChangeDirectory) IsTerminalOperation() {}

type TypeInCommand struct {
	Command string `json:"Command"`
}

func (o TypeInCommand) IsTerminalOperation() {}

type MarkLastCommandExecuted struct {
	Command string `json:"Command"`
}

func (o MarkLastCommandExecuted) IsTerminalOperation() {}

type WriteOutput struct {
	Output string `json:"Output"`
}

func (o WriteOutput) IsTerminalOperation() {}

func TestTerminal(t *testing.T) {
	type Operation struct {
		operation     TerminalOperation
		expectSuccess bool
	}

	type Entry struct {
		name       string
		operations []Operation
		resultFile string
	}

	var entries []Entry

	runEntries := func(t *testing.T, testEntries []Entry) {
		for i, e := range testEntries {
			t.Run(e.name, func(t *testing.T) {
				terminal := NewTerminal("default")
				for j, op := range e.operations {
					var err error
					switch v := op.operation.(type) {
					case ChangeDirectory:
						terminal.ChangeCurrentDirectory(v.FilePath)
					case TypeInCommand:
						err = terminal.TypeInCommand(v.Command)
					case MarkLastCommandExecuted:
						err = terminal.MarkLastCommandExecuted(v.Command)
					case WriteOutput:
						err = terminal.WriteOutput(v.Output)
					default:
						t.Fatalf("entry %d, op %d faild:\nwrong op.operation has type = %v", i, j, reflect.TypeOf(v))
						return
					}

					resultSuccess := err == nil
					if resultSuccess != op.expectSuccess {
						errMsg1 := fmt.Sprintf("operation %s is expected, but result is %s\n", statusString(op.expectSuccess), statusString(resultSuccess))

						var errMsg2 string = ""
						if op.expectSuccess {
							errMsg2 = "error:     " + err.Error() + "\n"
						}

						errMsg3 := fmt.Sprintf("operation: %+v\n", op)
						t.Errorf("%s%s%s", errMsg1, errMsg2, errMsg3)
						return
					}
				}

				compareAfterMarshal(t, e.resultFile, terminal)
			})
		}
	}

	entries = []Entry{}

	t.Run("add_file", func(t *testing.T) { runEntries(t, entries) })
}

func TestTerminalNew(t *testing.T) {
	terminal := NewTerminal("default")
	compareAfterMarshal(t, "testdata/terminal/new-terminal.json", terminal)
}

func TestTerminalCd1(t *testing.T) {
	terminal := NewTerminal("default")
	terminal.ChangeCurrentDirectory("hello")
	compareAfterMarshal(t, "testdata/terminal/cd1.json", terminal)
}

func TestTerminalCd2(t *testing.T) {
	terminal := NewTerminal("default")
	terminal.ChangeCurrentDirectory("hello/world/thunder")
	compareAfterMarshal(t, "testdata/terminal/cd2.json", terminal)
}

func TestTerminalTypein1(t *testing.T) {
	terminal := NewTerminal("default")
	if err := terminal.TypeInCommand("mkdir abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}

	compareAfterMarshal(t, "testdata/terminal/type-in-command1.json", terminal)
}

func TestTerminalTypein2(t *testing.T) {
	terminal := NewTerminal("default")
	if err := terminal.TypeInCommand("mkdir abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.MarkLastCommandExecuted("mkdir abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.TypeInCommand("mkdir cde"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}

	compareAfterMarshal(t, "testdata/terminal/type-in-command2.json", terminal)
}

func TestTerminalTypeinFail(t *testing.T) {

	terminal := NewTerminal("default")
	if err := terminal.TypeInCommand("mkdir abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.TypeInCommand("mkdir cde"); err == nil {
		t.Fatalf("error expected")
	}

	// not changed from the initial command
	compareAfterMarshal(t, "testdata/terminal/type-in-command1.json", terminal)
}

func TestTerminalWriteOutput1(t *testing.T) {
	terminal := NewTerminal("default")
	if err := terminal.TypeInCommand("echo abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.MarkLastCommandExecuted("echo abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}
	if err := terminal.WriteOutput("abc"); err != nil {
		t.Fatalf("no error expected, but %s", err)
	}

	compareAfterMarshal(t, "testdata/terminal/write-output.json", terminal)
}
