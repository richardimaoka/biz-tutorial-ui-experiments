package model

import (
	"encoding/json"
	"os"
	"testing"
)

const (
	dataDir string = "../data"
	testDir string = "../data/test"
)

func address(s string) *string {
	return &s
}

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

func TestFilesInDir(t *testing.T) {
	dir := "testdata/action/input"
	prefix := "input"

	resultFiles, err := FilesInDir(dir, prefix)
	if err != nil {
		t.Fatalf("error reading files in %s with prefix = %s", dir, prefix)
	}

	if len(resultFiles) != 46 {
		t.Errorf("expected 46 but got %d", len(resultFiles))
	}
}

func TestSplitActoinList(t *testing.T) {
	targetPrefix := "input"
	actionListFile := "testdata/action/action_list.json"

	// the function to test

	if err := SplitActionList(actionListFile, testDir, targetPrefix); err != nil {
		t.Fatal(err)
	}

	// expectation and results

	expectedFiles, err := FilesInDir("testdata/action/input", targetPrefix)
	if err != nil {
		t.Fatalf("error reading files in testdata with prefix = %s", targetPrefix)
	}

	resultFiles, err := FilesInDir(testDir, targetPrefix)
	if err != nil {
		t.Fatalf("error reading files in %s with prefix = %s", testDir, targetPrefix)
	}

	// run check result

	if len(expectedFiles) != len(resultFiles) {
		t.Fatalf("expected %d files but result is %d files", len(expectedFiles), len(resultFiles))
	}

	for i := range expectedFiles {
		expectedBytes, err := os.ReadFile(expectedFiles[i])
		if err != nil {
			t.Errorf("failed to read %s", expectedFiles[i])
			continue
		}

		resultBytes, err := os.ReadFile(resultFiles[i])
		if err != nil {
			t.Errorf("failed to read %s", resultFiles[i])
			continue
		}

		if err := compareJsonBytes(expectedBytes, resultBytes); err != nil {
			t.Errorf("failed to compare files = %s vs. %s, %s", expectedFiles[i], resultFiles[i], err)
		}
	}
}

func TestReadOperationFromBytes(t *testing.T) {
	// file := fmt.Sprintf("testdata/action/source_code_ops", "file-add1.json")
	// op := FileAdd{FilePath: "protoc-go-experiments/helloworld/greeting.pb"}
	// seqNo := 26
	// bytes, err := os.ReadFile(file)
	// if err != nil {
	// 	t.Fatalf("failed to read %s", file)
	// }
	// resultSeqNo, resultOp, err := readOperationFromBytes(bytes)
	// if err != nil {
	// 	t.Fatalf("failed to read op from %s", file)
	// }
	// compareJsonBytes(t, by, resultOp)
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
				for _, op := range e.operations {
					err := e.command.Enrich(op.operation)

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

				compareAfterMarshal(t, e.resultFile, e.command)
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

func TestMain(m *testing.M) {
	// initial setup
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.Mkdir(dataDir, 0755); err != nil {
			panic("failed to create " + dataDir)
		}
	}
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		if err := os.Mkdir(testDir, 0755); err != nil {
			panic("failed to create " + testDir)
		}
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}
