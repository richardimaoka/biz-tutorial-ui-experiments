package model

import (
	"encoding/json"
	"os"
	"testing"
)

func TestActionCommandMarshal(t *testing.T) {
	type Entry struct {
		command      ActionCommand
		expectedFile string
	}

	entries := []Entry{
		{expectedFile: "testdata/action/command/action_command1.json", command: ActionCommand{TerminalName: "default", Command: "mkdir hello"}},
		{expectedFile: "testdata/action/command/action_command2.json", command: ActionCommand{TerminalName: "default", Command: "echo abc", Output: address("abc")}},
		{expectedFile: "testdata/action/command/action_command3.json", command: ActionCommand{TerminalName: "another", Command: "cd hello/world", CurrentDirectory: address("hello/world")}},
	}

	for _, e := range entries {
		t.Run("test_action_command_marshal", func(t *testing.T) {
			compareAfterMarshal(t, e.expectedFile, e.command)
		})
	}
}

func TestActionCommandUnmarshal(t *testing.T) {
	files := []string{
		"testdata/action/command/action_command1.json",
		"testdata/action/command/action_command2.json",
		"testdata/action/command/action_command3.json",
		"testdata/action/command/action_command4.json",
	}

	for _, f := range files {
		t.Run("test_action_command_unmarshal", func(t *testing.T) {
			jsonBytes, err := os.ReadFile(f)
			if err != nil {
				t.Fatalf("failed to read %s", err)
			}

			var cmd ActionCommand
			if err := json.Unmarshal(jsonBytes, &cmd); err != nil {
				t.Fatalf("failed to unmarshal %s", err)
			}
			compareAfterMarshal(t, f, cmd)
		})
	}
}

//generate table-based tese cases for actions.enrich() method with all the operations
func TestEnrichActionCommand(t *testing.T) {
	type Operation struct {
		operation     FileSystemOperation
		expectSuccess bool
	}

	type Entry struct {
		name       string
		command    ActionCommand
		operations []Operation
		resultFile string
	}

	var entries []Entry

	runEntries := func(t *testing.T, testEntries []Entry) {
		for _, e := range testEntries {
			t.Run(e.name, func(t *testing.T) {
				var action Action = e.command
				var err error
				for _, op := range e.operations {
					action, err = action.Enrich(op.operation)

					resultSuccess := err == nil
					if resultSuccess != op.expectSuccess {
						t.Errorf("operation %t is expected, but result is %t\n", op.expectSuccess, resultSuccess)
						if op.expectSuccess {
							t.Errorf("error:     %s\n", err.Error())
						}
						t.Errorf("operation: %+v\n", op)
						return
					}
				}

				compareAfterMarshal(t, e.resultFile, action)
			})
		}
	}

	entries = []Entry{
		{name: "add_file_single",
			resultFile: "testdata/action/enrich/action1.json",
			command:    ActionCommand{TerminalName: "default", Command: "git apply 324x435d"},
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{Content: "***", IsFullContent: false, FilePath: "protoc-go-experiments/helloworld/greeting.pb"}},
			},
		},
	}
	t.Run("add_files", func(t *testing.T) { runEntries(t, entries) })
}
