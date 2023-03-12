package model2

import (
	"fmt"
	"reflect"
	"testing"
)

func statusString(expectSuccess bool) string {
	if expectSuccess {
		return "success"
	} else {
		return "failure"
	}
}

func TestSourceCode_FileNodes(t *testing.T) {
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

	var entries []Entry

	runEntries := func(t *testing.T, testEntries []Entry) {
		for i, e := range testEntries {
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
							errMsg2 = "error:     " + err.Error() + "\n"
						}

						errMsg3 := fmt.Sprintf("operation: %+v\n", op)
						t.Errorf("%s%s%s", errMsg1, errMsg2, errMsg3)
						return
					}
				}

				compareAfterMarshal(t, e.resultFile, sc)
			})
		}
	}

	entries = []Entry{
		{name: "dir_create_SourceCode",
			operations: []Operation{}, // no operation
			resultFile: "testdata/source_code/new-source-code.json"},

		{name: "add_dir_single",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
			}, resultFile: "testdata/source_code/nodes/add-directory1.json"},

		{name: "add_dir_nested",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
			}, resultFile: "testdata/source_code/nodes/add-directory2.json"},

		{name: "add_dir_nested2",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
			}, resultFile: "testdata/source_code/nodes/add-directory3.json"},

		{name: "add_dir_multiple",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "aloha"},
			}, resultFile: "testdata/source_code/nodes/add-directory4.json"},

		{name: "error_add_dir_empty",
			operations: []Operation{
				{expectSuccess: false, operationType: OpAddDirectory, filePath: ""}, // "" is a wrong file path
			}, resultFile: "testdata/source_code/new-source-code.json"}, // json should be same as initial state

		{name: "error_add_dir_duplicate",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/japan"},
				{expectSuccess: false, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: false, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: false, operationType: OpAddDirectory, filePath: "hello/world/japan"},
			}, resultFile: "testdata/source_code/nodes/add-directory5.json"},
	}
	t.Run("add_directory", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "delete_dir_single",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpDeleteDirectory, filePath: "hello"},
			}, resultFile: "testdata/source_code/new-source-code.json"}, // json should be same as initial state

		{name: "delete_dir_nested_leaf",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpDeleteDirectory, filePath: "hello/world"},
			}, resultFile: "testdata/source_code/nodes/delete-directory1.json"},

		{name: "delete_dir_nested_middle",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world/japan"},
				// below "goodmorning.*" dirs are note affected
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning/hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning/hello/world"},
				{expectSuccess: true, operationType: OpDeleteDirectory, filePath: "hello/world"},
			}, resultFile: "testdata/source_code/nodes/delete-directory2.json"},

		{name: "delete_dir_nested_parent",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world/japan"},
				// below "goodmorning.*" dirs are note affected
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning/hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning/hello/world"},
				{expectSuccess: true, operationType: OpDeleteDirectory, filePath: "hello"},
			}, resultFile: "testdata/source_code/nodes/delete-directory3.json"},

		{name: "error_delete_dir_non_existent",
			operations: []Operation{
				// below "goodmorning.*" dirs are note affected
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning/hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning/hello/world"},
				{expectSuccess: false, operationType: OpDeleteDirectory, filePath: "goodmorning/hello/universe"},
				{expectSuccess: false, operationType: OpDeleteDirectory, filePath: "goodmorning/vonjour/world"},
			}, resultFile: "testdata/source_code/nodes/delete-directory4.json"},

		{name: "error_delete_dir_twice",
			operations: []Operation{
				// below "goodmorning.*" dirs are note affected
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning/hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "goodmorning/hello/world"},
				{expectSuccess: true, operationType: OpDeleteDirectory, filePath: "goodmorning/hello/world"},
				{expectSuccess: false, operationType: OpDeleteDirectory, filePath: "goodmorning/hello/world"},
			}, resultFile: "testdata/source_code/nodes/delete-directory5.json"},
	}
	t.Run("delete_directory", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "add_file_single",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello.txt"},
			}, resultFile: "testdata/source_code/nodes/add-file1.json"},

		{name: "add_file_nested",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world.txt"},
			}, resultFile: "testdata/source_code/nodes/add-file2.json"},

		{name: "add_file_nested2",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/japan.txt"},
			}, resultFile: "testdata/source_code/nodes/add-file3.json"},

		{name: "add_file_nested3",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/japan.txt"},
			}, resultFile: "testdata/source_code/nodes/add-file4.json"},

		{name: "add_file_next_to",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/japan.txt"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/america.txt"},
			}, resultFile: "testdata/source_code/nodes/add-file5.json"},

		{name: "error_add_file_duplicate1",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello.txt"},
				{expectSuccess: false, operationType: OpAddFile, filePath: "hello.txt"},
			}, resultFile: "testdata/source_code/nodes/add-file1.json"},

		{name: "error_add_file_duplicate2",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world.txt"},
				{expectSuccess: false, operationType: OpAddFile, filePath: "hello/world.txt"},
				{expectSuccess: false, operationType: OpAddFile, filePath: "hello"},
			}, resultFile: "testdata/source_code/nodes/add-file2.json"},
	}
	t.Run("add_file", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "delete_file_single",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello.txt"},
				{expectSuccess: true, operationType: OpDeleteFile, filePath: "hello.txt"},
			}, resultFile: "testdata/source_code/new-source-code.json"},

		{name: "delete_file_nested",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/japan.txt"},
				{expectSuccess: true, operationType: OpDeleteFile, filePath: "hello/world/japan.txt"},
			}, resultFile: "testdata/source_code/nodes/delete-file1.json"},

		{name: "delete_file_next_to",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/japan.txt"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/america.txt"},
				{expectSuccess: true, operationType: OpDeleteFile, filePath: "hello/world/japan.txt"},
			}, resultFile: "testdata/source_code/nodes/delete-file2.json"},

		{name: "error_delete_file_non_existent",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/japan.txt"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/america.txt"},
				{expectSuccess: false, operationType: OpDeleteFile, filePath: "hello/world/france.txt"},
			}, resultFile: "testdata/source_code/nodes/delete-file3.json"},

		{name: "error_delete_file_twice",
			operations: []Operation{
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello"},
				{expectSuccess: true, operationType: OpAddDirectory, filePath: "hello/world"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/japan.txt"},
				{expectSuccess: true, operationType: OpAddFile, filePath: "hello/world/america.txt"},
				{expectSuccess: true, operationType: OpDeleteFile, filePath: "hello/world/japan.txt"},
				{expectSuccess: false, operationType: OpDeleteFile, filePath: "hello/world/japan.txt"},
			}, resultFile: "testdata/source_code/nodes/delete-file2.json"},
	}
	t.Run("delete_file", func(t *testing.T) { runEntries(t, entries) })
}

func TestSourceCode_Contents(t *testing.T) {
	type Operation struct {
		operation     FileSystemOperation
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
				sc := NewSourceCode()
				for j, op := range e.operations {
					var err error
					switch v := op.operation.(type) {
					case FileAdd:
						if err := sc.AddFile(v); err != nil {
							t.Fatal(err)
						}
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

				compareAfterMarshal(t, e.resultFile, sc)
			})
		}
	}

	entries = []Entry{
		{name: "add_file_single",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}},
			}, resultFile: "testdata/source_code/contents/add-file1.json"},

		{name: "add_file_nested",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/new.txt", Content: "hello new world", IsFullContent: true}},
			}, resultFile: "testdata/source_code/contents/add-file2.json"},

		//TODO: fill the same cases as node tests
		//add file next_to
		//error add file duplicate1
		//error add file duplicate2

		//delete dir nested leaf
		//delete dir nested middle
		//delete dir nested parent

		//delete file single
		//delete file nested
		//error delete file dir
		//error delete file non_existent
		//error delete file twice
	}

	t.Run("add_file", func(t *testing.T) { runEntries(t, entries) })
}
