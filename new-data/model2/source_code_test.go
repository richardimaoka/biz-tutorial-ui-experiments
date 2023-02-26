package model2

import (
	"fmt"
	"testing"
)

type OperationType string

const (
	OpAddDirectory    OperationType = "Add Directory"
	OpDeleteDirectory OperationType = "Delete Directory"
	OpAddFile         OperationType = "Add File"
	OpDeleteFile      OperationType = "Delete File"
)

type Operation struct {
	filePath      string
	operationType OperationType
	expectSuccess bool
}

type Entry struct {
	name       string
	operations []Operation
	resultFile string
}

func statusString(expectSuccess bool) string {
	if expectSuccess {
		return "success"
	} else {
		return "failure"
	}
}

func TestAddDirectoryCases(t *testing.T) {
	var entries []Entry = []Entry{
		{name: "create_SourceCode",
			operations: []Operation{}, // no operation
			resultFile: "testdata/new-source-code.json"},

		{name: "error_add_dir_empty_file_path",
			operations: []Operation{
				{filePath: "", operationType: OpAddDirectory, expectSuccess: false}, // "" is a wrong file path
			}, resultFile: "testdata/new-source-code.json"}, // json should be same as initial state

		{name: "add_single_dir",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
			}, resultFile: "testdata/add-directory1.json"},

		{name: "add_dir_and_child_dir",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world", operationType: OpAddDirectory, expectSuccess: true},
			}, resultFile: "testdata/add-directory2.json"},

		{name: "add_dir_and_child_dir_and_another_dir",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "aloha", operationType: OpAddDirectory, expectSuccess: true},
			}, resultFile: "testdata/add-directory3.json"},
	}

	runEntries(t, entries)
}

func TestDeleteDirectoryCases(t *testing.T) {
	var entries []Entry = []Entry{
		{name: "add_and_delete_dir",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello", operationType: OpDeleteDirectory, expectSuccess: true},
			}, resultFile: "testdata/new-source-code.json"}, // json should be same as initial state

		{name: "add_nested_dirs_and_delete_dir",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world", operationType: OpDeleteDirectory, expectSuccess: true},
			}, resultFile: "testdata/delete-directory1.json"},
	}

	runEntries(t, entries)
}
func TestFileCases(t *testing.T) {
	var entries []Entry = []Entry{
		{name: "add_a_file",
			operations: []Operation{
				{filePath: "hello.txt", operationType: OpAddFile, expectSuccess: true},
			}, resultFile: "testdata/add-file1.json"},

		{name: "error_adding_a_file",
			operations: []Operation{
				{filePath: "hello/world.txt", operationType: OpAddFile, expectSuccess: false},
			}, resultFile: "testdata/new-source-code.json"},

		{name: "add_two_files",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world.txt", operationType: OpAddFile, expectSuccess: true},
			}, resultFile: "testdata/add-file2.json"},

		{name: "add_and_delete_a_file",
			operations: []Operation{
				{filePath: "hello.txt", operationType: OpAddFile, expectSuccess: true},
				{filePath: "hello.txt", operationType: OpDeleteFile, expectSuccess: true},
			}, resultFile: "testdata/new-source-code.json"},

		{name: "add_three_files_and_delete_one",
			operations: []Operation{
				{filePath: "goodmorning.txt", operationType: OpAddFile, expectSuccess: true},
				{filePath: "hello.txt", operationType: OpAddFile, expectSuccess: true},
				{filePath: "evening.txt", operationType: OpAddFile, expectSuccess: true},
				{filePath: "goodmorning.txt", operationType: OpDeleteFile, expectSuccess: true},
			}, resultFile: "testdata/add-file3.json"},

		{name: "add_and_delete_a_nested_file",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world.txt", operationType: OpAddFile, expectSuccess: true},
				{filePath: "hello/world.txt", operationType: OpDeleteFile, expectSuccess: true},
			}, resultFile: "testdata/add-delete.json"},
	}

	runEntries(t, entries)
}

func runEntries(t *testing.T, entries []Entry) {
	for i, e := range entries {
		t.Run(e.name, func(t *testing.T) {
			sc := NewSourceCode()
			for j, op := range e.operations {
				var err error
				switch op.operationType {
				case OpAddDirectory:
					err = sc.AddDirectoryNode(op.filePath)
				case OpDeleteDirectory:
					err = sc.DeleteDirectoryNode(op.filePath)
				case OpAddFile:
					err = sc.AddFileNode(op.filePath)
				case OpDeleteFile:
					err = sc.DeleteFileNode(op.filePath)
				default:
					t.Fatalf("entry %d, op %d faild:\nwrong op.nodeType = %s", i, j, op.operationType)
					return
				}

				resultSuccess := err == nil
				if resultSuccess != op.expectSuccess {
					errMsg1 := fmt.Sprintf("operation %s is expected, but result is %s\n", statusString(op.expectSuccess), statusString(resultSuccess))

					var errMsg2 string = ""
					if op.expectSuccess {
						errMsg2 = "error: " + err.Error() + "\n"
					}

					errMsg3 := fmt.Sprintf("operation = %+v\n", op)
					errMsg4 := fmt.Sprintf("entry = %+v", e)
					t.Errorf("%s%s%s%s", errMsg1, errMsg2, errMsg3, errMsg4)
					return
				}
			}
			compareAfterMarshal(t, e.resultFile, sc)
		})
	}
}
