package model2

import (
	"fmt"
	"reflect"
	"testing"
)

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

	entries = []Entry{
		{name: "new_terminal",
			operations: []Operation{}, // no operation
			resultFile: "testdata/terminal/new-terminal.json"},

		{name: "cd1",
			operations: []Operation{
				{expectSuccess: true, operation: ChangeDirectory{FilePath: "hello"}},
			},
			resultFile: "testdata/terminal/cd1.json"},

		{name: "cd2",
			operations: []Operation{
				{expectSuccess: true, operation: ChangeDirectory{FilePath: "hello/world/thunder"}},
			},
			resultFile: "testdata/terminal/cd2.json"},
	}

	t.Run("cd", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "type_in_single",
			operations: []Operation{
				{expectSuccess: true, operation: TypeInCommand{Command: "mkdir abc"}},
			},
			resultFile: "testdata/terminal/type-in-command1.json"},

		{name: "error_type_in_continuous",
			operations: []Operation{
				{expectSuccess: true, operation: TypeInCommand{Command: "mkdir abc"}},
				{expectSuccess: false, operation: TypeInCommand{Command: "mkdir efg"}},
			},
			resultFile: "testdata/terminal/type-in-command1.json"},
	}

	t.Run("type_in", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "mark_single",
			operations: []Operation{
				{expectSuccess: true, operation: TypeInCommand{Command: "mkdir abc"}},
				{expectSuccess: true, operation: MarkLastCommandExecuted{Command: "mkdir abc"}},
			},
			resultFile: "testdata/terminal/mark-last-command-executed1.json"},

		{name: "mark_and_typein",
			operations: []Operation{
				{expectSuccess: true, operation: TypeInCommand{Command: "mkdir abc"}},
				{expectSuccess: true, operation: MarkLastCommandExecuted{Command: "mkdir abc"}},
				{expectSuccess: true, operation: TypeInCommand{Command: "mkdir efg"}},
			},
			resultFile: "testdata/terminal/mark-last-command-executed2.json"},

		{name: "mark_two",
			operations: []Operation{
				{expectSuccess: true, operation: TypeInCommand{Command: "mkdir abc"}},
				{expectSuccess: true, operation: MarkLastCommandExecuted{Command: "mkdir abc"}},
				{expectSuccess: true, operation: TypeInCommand{Command: "mkdir efg"}},
				{expectSuccess: true, operation: MarkLastCommandExecuted{Command: "mkdir efg"}},
			},
			resultFile: "testdata/terminal/mark-last-command-executed3.json"},
	}

	t.Run("mark_last_command_executed", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "output1",
			operations: []Operation{
				{expectSuccess: true, operation: TypeInCommand{Command: "echo abc"}},
				{expectSuccess: true, operation: MarkLastCommandExecuted{Command: "echo abc"}},
				{expectSuccess: true, operation: WriteOutput{Output: "abc"}},
			},
			resultFile: "testdata/terminal/write-output1.json"},

		{name: "write_output2",
			operations: []Operation{
				{expectSuccess: true, operation: TypeInCommand{Command: "echo abc"}},
				{expectSuccess: true, operation: MarkLastCommandExecuted{Command: "echo abc"}},
				{expectSuccess: true, operation: WriteOutput{Output: "abc"}},
				{expectSuccess: true, operation: TypeInCommand{Command: "echo efg"}},
				{expectSuccess: true, operation: MarkLastCommandExecuted{Command: "echo efg"}},
				{expectSuccess: true, operation: WriteOutput{Output: "efg"}},
			},
			resultFile: "testdata/terminal/write-output2.json"},

		{name: "error_output_should_follow_execution",
			operations: []Operation{
				{expectSuccess: true, operation: TypeInCommand{Command: "echo abc"}},
				{expectSuccess: false, operation: WriteOutput{Output: "abc"}},
			},
			resultFile: "testdata/terminal/write-output3.json"},
	}
	t.Run("write_output", func(t *testing.T) { runEntries(t, entries) })
}
