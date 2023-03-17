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
					case DirectoryAdd:
						err = sc.AddDirectory(v)
					case DirectoryDelete:
						err = sc.DeleteDirectory(v)
					case FileAdd:
						err = sc.AddFile(v)
					case FileUpdate:
						err = sc.UpdateFile(v)
					case FileDelete:
						err = sc.DeleteFile(v)
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
		{name: "new",
			operations: []Operation{}, resultFile: "testdata/source_code/new-source-code.json"},
	}

	t.Run("new_source_code", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "add_dir_single",
			operations: []Operation{
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello"}},
			}, resultFile: "testdata/source_code/add-directory1.json"},

		{name: "add_dir_nested",
			operations: []Operation{
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world"}},
			}, resultFile: "testdata/source_code/add-directory2.json"},

		{name: "add_dir_nested2",
			operations: []Operation{
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world"}},
			}, resultFile: "testdata/source_code/add-directory3.json"},

		{name: "add_dir_multiple",
			operations: []Operation{
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "aloha"}},
			}, resultFile: "testdata/source_code/add-directory4.json"},

		{name: "error_add_dir_empty",
			operations: []Operation{
				{expectSuccess: false, operation: DirectoryAdd{FilePath: ""}}, // "" is a wrong file path
			}, resultFile: "testdata/source_code/new-source-code.json"}, // json should be same as initial state

		{name: "error_add_dir_duplicate",
			operations: []Operation{
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world"}},
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/japan", Content: "hello", IsFullContent: true}},
				{expectSuccess: false, operation: DirectoryAdd{FilePath: "hello"}},
				{expectSuccess: false, operation: DirectoryAdd{FilePath: "hello/world"}},
				{expectSuccess: false, operation: DirectoryAdd{FilePath: "hello/world/japan"}},
			}, resultFile: "testdata/source_code/add-directory5.json"},
	}

	t.Run("add_directory", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "delete_dir_single",
			operations: []Operation{
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello"}},
				{expectSuccess: true, operation: DirectoryDelete{FilePath: "hello"}},
			}, resultFile: "testdata/source_code/new-source-code.json"}, // json should be same as initial state

		{name: "delete_dir_nested_leaf",
			operations: []Operation{
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world"}},
				{expectSuccess: true, operation: DirectoryDelete{FilePath: "hello/world"}},
			}, resultFile: "testdata/source_code/delete-directory1.json"},

		{name: "delete_dir_nested_middle",
			operations: []Operation{
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world/japan"}},
				// below "goodmorning.*" dirs are note affected
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning/hello"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning/hello/world"}},
				{expectSuccess: true, operation: DirectoryDelete{FilePath: "hello/world"}},
			}, resultFile: "testdata/source_code/delete-directory2.json"},

		{name: "delete_dir_nested_parent",
			operations: []Operation{
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world/japan"}},
				// below "goodmorning.*" dirs are note affected
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning/hello"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning/hello/world"}},
				{expectSuccess: true, operation: DirectoryDelete{FilePath: "hello"}},
			}, resultFile: "testdata/source_code/delete-directory3.json"},

		{name: "error_delete_dir_non_existent",
			operations: []Operation{
				// below "goodmorning.*" dirs are note affected
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning/hello"}},
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning/hello/world"}},
				{expectSuccess: false, operation: DirectoryDelete{FilePath: "goodmorning/hello/universe"}},
				{expectSuccess: false, operation: DirectoryDelete{FilePath: "goodmorning/vonjour/world"}},
			}, resultFile: "testdata/source_code/delete-directory4.json"},

		{name: "error_delete_dir_twice",
			operations: []Operation{
				// below "goodmorning.*" dirs are note affected
				{expectSuccess: true, operation: DirectoryAdd{FilePath: "goodmorning/hello/world"}},
				{expectSuccess: true, operation: DirectoryDelete{FilePath: "goodmorning/hello/world"}},
				{expectSuccess: false, operation: DirectoryDelete{FilePath: "goodmorning/hello/world"}},
			}, resultFile: "testdata/source_code/delete-directory5.json"},
	}
	t.Run("delete_directory", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "add_file_single",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}},
			}, resultFile: "testdata/source_code/add-file1.json"},

		{name: "add_file_nested",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/new.txt", Content: "hello new world", IsFullContent: true}},
			}, resultFile: "testdata/source_code/add-file2.json"},

		{name: "add_file_nested2",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/japan.txt"}},
			}, resultFile: "testdata/source_code/add-file3.json"},

		{name: "add_file_nested3",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/japan.txt"}},
			}, resultFile: "testdata/source_code/add-file4.json"},

		{name: "add_file_next_to",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/japan.txt"}},
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/america.txt"}},
			}, resultFile: "testdata/source_code/add-file5.json"},

		{name: "error_add_file_duplicate1",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello.txt", Content: "hello new world"}},
				{expectSuccess: false, operation: FileAdd{FilePath: "hello.txt"}},
			}, resultFile: "testdata/source_code/add-file1.json"},

		{name: "error_add_file_duplicate2",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world.txt"}},
				{expectSuccess: false, operation: FileAdd{FilePath: "hello/world.txt"}},
				{expectSuccess: false, operation: FileAdd{FilePath: "hello"}},
			}, resultFile: "testdata/source_code/add-file6.json"},
	}
	t.Run("add_file", func(t *testing.T) { runEntries(t, entries) })

	entries = []Entry{
		{name: "delete_file_single",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello.txt"}},
				{expectSuccess: true, operation: FileDelete{FilePath: "hello.txt"}},
			}, resultFile: "testdata/source_code/new-source-code.json"},

		{name: "delete_file_nested",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/japan.txt"}},
				{expectSuccess: true, operation: FileDelete{FilePath: "hello/world/japan.txt"}},
			}, resultFile: "testdata/source_code/delete-file1.json"},

		{name: "delete_file_next_to",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/japan.txt"}},
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/america.txt"}},
				{expectSuccess: true, operation: FileDelete{FilePath: "hello/world/japan.txt"}},
			}, resultFile: "testdata/source_code/delete-file2.json"},

		{name: "error_delete_file_non_existent",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/japan.txt"}},
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/america.txt"}},
				{expectSuccess: false, operation: FileDelete{FilePath: "hello/world/france.txt"}},
			}, resultFile: "testdata/source_code/delete-file3.json"},

		{name: "error_delete_file_twice",
			operations: []Operation{
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/japan.txt"}},
				{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/america.txt"}},
				{expectSuccess: true, operation: FileDelete{FilePath: "hello/world/japan.txt"}},
				{expectSuccess: false, operation: FileDelete{FilePath: "hello/world/japan.txt"}},
			}, resultFile: "testdata/source_code/delete-file2.json"},
	}
	t.Run("delete_file", func(t *testing.T) { runEntries(t, entries) })
}

func TestSourceCode_Diff(t *testing.T) {
	type Operation struct {
		diff          GitDiff
		expectSuccess bool
	}

	type Entry struct {
		name       string
		operations []Operation
		resultFile string
	}

	var entries []Entry

	runEntries := func(t *testing.T, testEntries []Entry) {
		for _, e := range testEntries {
			t.Run(e.name, func(t *testing.T) {
				sc := NewSourceCode()
				for _, op := range e.operations {
					err := sc.ApplyDiff(op.diff)

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
				{expectSuccess: true, diff: GitDiff{Added: []FileAdd{{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}}}},
			}, resultFile: "testdata/source_code/diff-file1.json",
		},

		{name: "empty",
			operations: []Operation{
				{expectSuccess: true, diff: GitDiff{}},
			}, resultFile: "testdata/source_code/new-source-code.json",
		},

		{name: "error_dupe",
			operations: []Operation{
				{expectSuccess: false, diff: GitDiff{Added: []FileAdd{
					{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true},
					{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true},
				}}},
			}, resultFile: "testdata/source_code/new-source-code.json",
		},
	}

	t.Run("diff", func(t *testing.T) { runEntries(t, entries) })
}
