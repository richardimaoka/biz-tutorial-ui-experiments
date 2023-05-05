package processing

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
)

func TestActionCommandMarshal(t *testing.T) {
	type Entry struct {
		name         string
		command      ActionCommand
		expectedFile string
	}

	entries := []Entry{
		{name: "command_only",
			expectedFile: "testdata/action/command/action_command_marshal1.json",
			command:      ActionCommand{TerminalName: "default", Command: "mkdir hello"}},
		{name: "another_terminal",
			expectedFile: "testdata/action/command/action_command_marshal2.json",
			command:      ActionCommand{TerminalName: "another", Command: "mkdir hello"}},
		{name: "command_output",
			expectedFile: "testdata/action/command/action_command_marshal3.json",
			command:      ActionCommand{TerminalName: "default", Command: "echo abc", Output: internal.Address("abc")}},
		{name: "command_cd",
			expectedFile: "testdata/action/command/action_command_marshal4.json",
			command:      ActionCommand{TerminalName: "another", Command: "cd hello/world", CurrentDirectory: internal.Address("hello/world")}},
		{name: "command_cd_output",
			expectedFile: "testdata/action/command/action_command_marshal5.json",
			command:      ActionCommand{TerminalName: "another", Command: "complex_command", Output: internal.Address("some output"), CurrentDirectory: internal.Address("hello/world")}},
		{name: "command_file_diff",
			expectedFile: "testdata/action/command/action_command_marshal6.json",
			command: ActionCommand{
				TerminalName: "default",
				Command:      "with_file_diff",
				Diff: Diff{
					FilesAdded:         []FileAdd{{FilePath: "a/b/c", Content: "file content", IsFullContent: true}},
					FilesDeleted:       []FileDelete{{FilePath: "a/b/d"}, {FilePath: "a/b/e"}},
					DirectoriesDeleted: []DirectoryDelete{{FilePath: "aa/b/d"}, {FilePath: "aa/b/e"}},
				},
			},
		},
		{name: "command_directory_effect",
			expectedFile: "testdata/action/command/action_command_marshal7.json",
			command: ActionCommand{
				TerminalName: "default",
				Command:      "with_dir_diff",
				Diff: Diff{
					DirectoriesDeleted: []DirectoryDelete{{FilePath: "a/b/d"}, {FilePath: "a/b/e"}},
				},
			},
		},
	}

	for _, e := range entries {
		t.Run("test_action_command_marshal", func(t *testing.T) {
			internal.CompareAfterMarshal(t, e.expectedFile, e.command)
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
			internal.CompareAfterMarshal(t, f, cmd)
		})
	}
}

func TestEnrichActionCommandDiff(t *testing.T) {
	type Operation struct {
		operation     FileSystemOperation
		expectSuccess bool
	}

	type Entry struct {
		name       string
		action     Action
		operations []Operation
		resultFile string
	}

	runEntries := func(t *testing.T, testEntries []Entry) {
		for _, e := range testEntries {
			t.Run(e.name, func(t *testing.T) {
				for _, op := range e.operations {
					e.action.Enrich(op.operation)
				}

				internal.CompareAfterMarshal(t, e.resultFile, e.action)
			})
		}
	}

	t.Run("action_diff", func(t *testing.T) {
		runEntries(t, []Entry{
			{name: "add_file_single",
				resultFile: "testdata/action/enrich/action1.json",
				action:     &ActionCommand{TerminalName: "default", Command: "git apply 324x435d"},
				operations: []Operation{
					{expectSuccess: true, operation: FileAdd{Content: "***", IsFullContent: false, FilePath: "protoc-go-experiments/helloworld/greeting.pb"}},
				},
			},
		})
	})
}
