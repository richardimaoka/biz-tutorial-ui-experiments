package processing

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_SourceCodeProcessor(t *testing.T) {
	type Operation struct {
		operation     FileSystemOperation
		expectSuccess bool
	}

	type Entry struct {
		name       string
		operations []Operation
		resultFile string
	}

	applyOperations := func(processor *SourceCodeProcessor, operations []Operation) error {
		for opNum, op := range operations {
			var opError error
			switch v := op.operation.(type) {
			case DirectoryAdd:
				opError = processor.AddDirectory(v)
			case DirectoryDelete:
				opError = processor.DeleteDirectory(v)
			case FileAdd:
				opError = processor.AddFile(v)
			case FileUpdate:
				opError = processor.UpdateFile(v)
			case FileDelete:
				opError = processor.DeleteFile(v)
			default:
				return fmt.Errorf("op %d faild:\nwrong operation type = %v", opNum, reflect.TypeOf(v))
			}

			opSuccess := opError == nil
			if opSuccess != op.expectSuccess { //if result is not expected
				if op.expectSuccess {
					return fmt.Errorf(
						"operations[%d] success is expected, but result is failure\nerror:     %s\noperation: %+v\n",
						opNum,
						opError.Error(),
						op)
				} else {
					return fmt.Errorf(
						"operations[%d] failure is expected, but result is success\noperation: %+v\n",
						opNum,
						op)
				}
			}
		}

		return nil
	}

	// accept testEntries, and run tests on them
	runEntries := func(t *testing.T, testEntries []Entry) {
		for _, e := range testEntries {
			t.Run(e.name, func(t *testing.T) {
				sourceCode := NewSourceCodeProcessor()
				if err := applyOperations(sourceCode, e.operations); err != nil {
					t.Errorf(err.Error()) // fail but continue running, as two ops can fail in sequence
					return
				}

				compareAfterMarshal(t, e.resultFile, sourceCode.ToGraphQLModel())
			})
		}
	}
	t.Run("new_source_code", func(t *testing.T) {
		runEntries(t, []Entry{
			{name: "new",
				operations: []Operation{}, resultFile: "testdata/source_code/new-source-code.json"},
		})
	})

	t.Run("add_directory", func(t *testing.T) {
		//reuse runEntries logic
		runEntries(t,
			[]Entry{
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

				{name: "error_add_dir_duplicate1",
					operations: []Operation{
						{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello"}},
						{expectSuccess: false, operation: DirectoryAdd{FilePath: "hello"}},
					}, resultFile: "testdata/source_code/add-directory1.json"},

				{name: "add_dir_nested2",
					operations: []Operation{
						{expectSuccess: true, operation: DirectoryAdd{FilePath: "hello/world"}},
					}, resultFile: "testdata/source_code/add-directory3.json"},
			})
	})

	t.Run("add_file", func(t *testing.T) {
		runEntries(t, []Entry{
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

			{name: "add_file_next_to",
				operations: []Operation{
					{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/japan.txt"}},
					{expectSuccess: true, operation: FileAdd{FilePath: "hello/world/america.txt"}},
				}, resultFile: "testdata/source_code/add-file4.json"},

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
				}, resultFile: "testdata/source_code/add-file5.json"},
		})
	})

	t.Run("delete_directory", func(t *testing.T) {
		runEntries(t, []Entry{
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
					// below "goodmorning.*" dirs are not affected
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
		})
	})

	t.Run("delete_file", func(t *testing.T) {
		runEntries(t, []Entry{
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
		})
	})
}

func TestSourceCode_Diff(t *testing.T) {
	type Operation struct {
		diff          Diff
		expectSuccess bool
	}

	type Entry struct {
		name       string
		operations []Operation
		resultFile string
	}

	applyOperations := func(processor *SourceCodeProcessor, operations []Operation) error {
		for opNum, op := range operations {
			opError := processor.ApplyDiff2(op.diff)

			opSuccess := opError == nil
			if opSuccess != op.expectSuccess { //if result is not expected
				if op.expectSuccess {
					return fmt.Errorf(
						"operations[%d] success is expected, but result is failure\nerror:     %s\noperation: %+v\n",
						opNum,
						opError.Error(),
						op)
				} else {
					return fmt.Errorf(
						"operations[%d] failure is expected, but result is success\noperation: %+v\n",
						opNum,
						op)
				}
			}
		}

		return nil
	}

	// accept testEntries, and run tests on them
	runEntries := func(t *testing.T, testEntries []Entry) {
		for _, e := range testEntries {
			t.Run(e.name, func(t *testing.T) {
				sourceCode := NewSourceCodeProcessor()
				if err := applyOperations(sourceCode, e.operations); err != nil {
					t.Errorf(err.Error()) // fail but continue running, as two ops can fail in sequence
					return
				}

				compareAfterMarshal(t, e.resultFile, sourceCode.ToGraphQLModel())
			})
		}
	}

	t.Run("diff", func(t *testing.T) {
		runEntries(t, []Entry{
			{name: "add_file_single",
				operations: []Operation{
					{expectSuccess: true, diff: Diff{FilesAdded: []FileAdd{{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}}}},
				}, resultFile: "testdata/source_code/diff-file1.json",
			},

			{name: "empty",
				operations: []Operation{
					{expectSuccess: true, diff: Diff{}},
				}, resultFile: "testdata/source_code/new-source-code.json",
			},

			{name: "error_dupe",
				operations: []Operation{
					{expectSuccess: false, diff: Diff{FilesAdded: []FileAdd{
						{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true},
						{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true},
					}}},
				}, resultFile: "testdata/source_code/new-source-code.json",
			},
		})
	})
}

func TestSourceCode_Mutation(t *testing.T) {
	sourceCode := NewSourceCodeProcessor()
	sourceCode.AddFile(FileAdd{FilePath: "hello/world/japan.txt"})
	sourceCode.AddFile(FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
	sourceCode.AddDirectory(DirectoryAdd{FilePath: "goodmorning/hello/world"})

	// once materialized GraphQL model...
	materialized := sourceCode.ToGraphQLModel()
	compareAfterMarshal(t, "testdata/source_code/materialized.json", materialized)

	// ...mutation afterwards should have no effect
	sourceCode.step = "mutated step"
	sourceCode.defaultOpenFilePath = "mutated/file/path"
	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).content = "mutated content"
	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).filePath = "mutated/path/to/file"
	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).isUpdated = true
	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).filePath = "mutated/path/to/dir"
	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).isUpdated = false

	// materialized GraphQL model should not be affected
	compareAfterMarshal(t,
		"testdata/source_code/materialized.json",
		materialized) // by changing this to sourceCode.ToGraphQLModel(), you can confirm mutation actually updated sourceCode
}

func TestSourceCode_Clone(t *testing.T) {
	sourceCode := NewSourceCodeProcessor()
	sourceCode.AddFile(FileAdd{FilePath: "hello/world/japan.txt"})
	sourceCode.AddFile(FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
	sourceCode.AddDirectory(DirectoryAdd{FilePath: "goodmorning/hello/world"})

	// once cloned ...
	sourceCodeCloned := sourceCode.Clone()
	compareAfterMarshal(t, "testdata/source_code/cloned.json", sourceCodeCloned.ToGraphQLModel())

	// ... mutation after cloning should have no effect
	sourceCode.step = "mutated step"
	sourceCode.defaultOpenFilePath = "mutated/file/path"
	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).content = "mutated content"
	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).filePath = "mutated/path/to/file"
	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).isUpdated = true
	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).filePath = "mutated/path/to/dir"
	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).isUpdated = false

	// cloned source code is not affected
	compareAfterMarshal(t,
		"testdata/source_code/cloned.json",
		sourceCodeCloned.ToGraphQLModel()) // by changing this to sourceCode.ToGraphQLModel(), you can confirm mutation actually updated sourceCode
}
