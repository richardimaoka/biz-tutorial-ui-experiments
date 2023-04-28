package processing

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

func TestActionCommandMarshal(t *testing.T) {
	type Entry struct {
		name         string
		command      model.ActionCommand
		expectedFile string
	}

	entries := []Entry{
		{name: "command_only",
			expectedFile: "testdata/action/command/action_command_marshal1.json",
			command:      model.ActionCommand{TerminalName: "default", Command: "mkdir hello"}},
		{name: "another_terminal",
			expectedFile: "testdata/action/command/action_command_marshal2.json",
			command:      model.ActionCommand{TerminalName: "another", Command: "mkdir hello"}},
		{name: "command_output",
			expectedFile: "testdata/action/command/action_command_marshal3.json",
			command:      model.ActionCommand{TerminalName: "default", Command: "echo abc", Output: address("abc")}},
		{name: "command_cd",
			expectedFile: "testdata/action/command/action_command_marshal4.json",
			command:      model.ActionCommand{TerminalName: "another", Command: "cd hello/world", CurrentDirectory: address("hello/world")}},
		{name: "command_cd_output",
			expectedFile: "testdata/action/command/action_command_marshal5.json",
			command:      model.ActionCommand{TerminalName: "another", Command: "complex_command", Output: address("some output"), CurrentDirectory: address("hello/world")}},
		{name: "command_file_effect",
			expectedFile: "testdata/action/command/action_command_marshal6.json",
			command: model.ActionCommand{
				TerminalName: "default",
				Command:      "with_file_effect",
				Effect: model.GitDiff{
					Added:   []model.FileAdd{{FilePath: "a/b/c", Content: "file content", IsFullContent: true}},
					Deleted: []model.FileDelete{{FilePath: "a/b/d"}, {FilePath: "a/b/e"}},
				},
			},
		},
		{name: "command_directory_effect",
			expectedFile: "testdata/action/command/action_command_marshal7.json",
			command: model.ActionCommand{
				TerminalName: "default",
				Command:      "with_dir_effect",
				Effect: model.DirectoryDiff{
					Deleted: []model.DirectoryDelete{{FilePath: "a/b/d"}, {FilePath: "a/b/e"}},
				},
			},
		},
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

			var cmd model.ActionCommand
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
		operation     model.FileSystemOperation
		expectSuccess bool
	}

	type Entry struct {
		name       string
		command    model.ActionCommand
		operations []Operation
		resultFile string
	}

	var entries []Entry

	runEntries := func(t *testing.T, testEntries []Entry) {
		for _, e := range testEntries {
			t.Run(e.name, func(t *testing.T) {
				var action model.Action = e.command
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
			command:    model.ActionCommand{TerminalName: "default", Command: "git apply 324x435d"},
			operations: []Operation{
				{expectSuccess: true, operation: model.FileAdd{Content: "***", IsFullContent: false, FilePath: "protoc-go-experiments/helloworld/greeting.pb"}},
			},
		},
	}
	t.Run("add_files", func(t *testing.T) { runEntries(t, entries) })
}
