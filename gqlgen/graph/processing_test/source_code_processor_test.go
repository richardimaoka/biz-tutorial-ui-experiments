package processing_test

import (
	"reflect"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

func Test_SourceCodeProcessor(t *testing.T) {
	type TestCase struct {
		ExpectSuccess bool
		ExpectedFile  string
		Operation     processing.FileSystemOperation
	}

	applyOperation := func(t *testing.T, processor *processing.SourceCodeProcessor, op processing.FileSystemOperation) error {
		switch v := op.(type) {
		case processing.DirectoryAdd:
			return processor.AddDirectory(v)
		case processing.DirectoryDelete:
			return processor.DeleteDirectory(v)
		case processing.FileAdd:
			return processor.AddFile(v)
		case processing.FileUpdate:
			return processor.UpdateFile(v)
		case processing.FileDelete:
			return processor.DeleteFile(v)
		default:
			t.Fatalf("operation faild:\nwrong operation type = %v", reflect.TypeOf(v))
			return nil
		}
	}

	checkResult := func(t *testing.T, index int, op processing.FileSystemOperation, expectSuccess bool, err error) {
		resultSuccess := err == nil
		if resultSuccess != expectSuccess { //if result is not expected
			if expectSuccess {
				t.Fatalf("operation[%d] success is expected, but result is failure\noperation: %+v\nerror: %s\n", index, op, err)
			} else {
				t.Fatalf("operation[%d] failure is expected, but result is success\noperation: %+v\n", index, op)
			}
		}
	}

	// accept testEntries, and run tests on them
	runEntries := func(t *testing.T, testCases []TestCase) {
		for i, c := range testCases {
			sourceCode := processing.NewSourceCodeProcessor()
			err := applyOperation(t, sourceCode, c.Operation)
			checkResult(t, i, c.Operation, c.ExpectSuccess, err)

			// if `-update` flag is passed from command line, update the golden file
			if *update {
				internal.WriteGoldenFile(t, c.ExpectedFile, sourceCode.ToGraphQLModel())
			}

			internal.CompareAfterMarshal(t, c.ExpectedFile, sourceCode.ToGraphQLModel())
		}
	}

	t.Run("add_directory", func(t *testing.T) {
		runEntries(t, []TestCase{
			{true, "testdata/source_code/add-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
		})
	})

	t.Run("add_dir_nested", func(t *testing.T) {
		runEntries(t, []TestCase{
			{true, "testdata/source_code/add-dir-nested1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
			{true, "testdata/source_code/add-dir-nested1-2.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		})
	})

	t.Run("add_dir_nested2", func(t *testing.T) {
		runEntries(t, []TestCase{
			{true, "testdata/source_code/add-directory3.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		})
	})

	t.Run("add_dir_multiple", func(t *testing.T) {
		runEntries(t, []TestCase{
			{true, "testdata/source_code/add_dir_multiple1-1.json", processing.DirectoryAdd{FilePath: "hello/world"}},
			{true, "testdata/source_code/add_dir_multiple1-2.json", processing.DirectoryAdd{FilePath: "aloha"}},
		})
	})

	// 			{name: "add_dir_multiple",
	// 				operations: []Operation{
	// 					{true, processing.DirectoryAdd{FilePath: "hello/world"}},
	// 					{true, processing.DirectoryAdd{FilePath: "aloha"}},
	// 				}, resultFile: "testdata/source_code/add-directory4.json"},

	// 			{name: "error_add_dir_empty",
	// 				operations: []Operation{
	// 					{false, processing.DirectoryAdd{FilePath: ""}}, // "" is a wrong file path
	// 				}, resultFile: "testdata/source_code/new-source-code.json"}, // json should be same as initial state

	// 			{name: "error_add_dir_duplicate1",
	// 				operations: []Operation{
	// 					{true, processing.DirectoryAdd{FilePath: "hello"}},
	// 					{false, processing.DirectoryAdd{FilePath: "hello"}},
	// 				}, resultFile: "testdata/source_code/add-directory1.json"},

	// 			{name: "add_dir_nested2",
	// 				operations: []Operation{
	// 					{true, processing.DirectoryAdd{FilePath: "hello/world"}},
	// 				}, resultFile: "testdata/source_code/add-directory3.json"},
	// 		})
	// })

	// t.Run("add_file", func(t *testing.T) {
	// 	runEntries(t, []Entry{
	// 		{name: "add_file_single",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}},
	// 			}, resultFile: "testdata/source_code/add-file1.json"},

	// 		{name: "add_file_nested",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello/world/new.txt", Content: "hello new world", IsFullContent: true}},
	// 			}, resultFile: "testdata/source_code/add-file2.json"},

	// 		{name: "add_file_nested2",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello/world/japan.txt"}},
	// 			}, resultFile: "testdata/source_code/add-file3.json"},

	// 		{name: "add_file_next_to",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello/world/japan.txt"}},
	// 				{true, FileAdd{FilePath: "hello/world/america.txt"}},
	// 			}, resultFile: "testdata/source_code/add-file4.json"},

	// 		{name: "error_add_file_duplicate1",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello.txt", Content: "hello new world"}},
	// 				{false, FileAdd{FilePath: "hello.txt"}},
	// 			}, resultFile: "testdata/source_code/add-file1.json"},

	// 		{name: "error_add_file_duplicate2",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello/world.txt"}},
	// 				{false, FileAdd{FilePath: "hello/world.txt"}},
	// 				{false, FileAdd{FilePath: "hello"}},
	// 			}, resultFile: "testdata/source_code/add-file5.json"},
	// 	})
	// })

	// t.Run("delete_directory", func(t *testing.T) {
	// 	runEntries(t, []Entry{
	// 		{name: "delete_dir_single",
	// 			operations: []Operation{
	// 				{true, processing.DirectoryAdd{FilePath: "hello"}},
	// 				{true, DirectoryDelete{FilePath: "hello"}},
	// 			}, resultFile: "testdata/source_code/new-source-code.json"}, // json should be same as initial state

	// 		{name: "delete_dir_nested_leaf",
	// 			operations: []Operation{
	// 				{true, processing.DirectoryAdd{FilePath: "hello"}},
	// 				{true, processing.DirectoryAdd{FilePath: "hello/world"}},
	// 				{true, DirectoryDelete{FilePath: "hello/world"}},
	// 			}, resultFile: "testdata/source_code/delete-directory1.json"},

	// 		{name: "delete_dir_nested_middle",
	// 			operations: []Operation{
	// 				{true, processing.DirectoryAdd{FilePath: "hello"}},
	// 				{true, processing.DirectoryAdd{FilePath: "hello/world"}},
	// 				{true, processing.DirectoryAdd{FilePath: "hello/world/japan"}},
	// 				// below "goodmorning.*" dirs are not affected
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning"}},
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning/hello"}},
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
	// 				{true, DirectoryDelete{FilePath: "hello/world"}},
	// 			}, resultFile: "testdata/source_code/delete-directory2.json"},

	// 		{name: "delete_dir_nested_parent",
	// 			operations: []Operation{
	// 				{true, processing.DirectoryAdd{FilePath: "hello"}},
	// 				{true, processing.DirectoryAdd{FilePath: "hello/world"}},
	// 				{true, processing.DirectoryAdd{FilePath: "hello/world/japan"}},
	// 				// below "goodmorning.*" dirs are note affected
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning"}},
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning/hello"}},
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
	// 				{true, DirectoryDelete{FilePath: "hello"}},
	// 			}, resultFile: "testdata/source_code/delete-directory3.json"},

	// 		{name: "error_delete_dir_non_existent",
	// 			operations: []Operation{
	// 				// below "goodmorning.*" dirs are note affected
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning"}},
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning/hello"}},
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
	// 				{false, DirectoryDelete{FilePath: "goodmorning/hello/universe"}},
	// 				{false, DirectoryDelete{FilePath: "goodmorning/vonjour/world"}},
	// 			}, resultFile: "testdata/source_code/delete-directory4.json"},

	// 		{name: "error_delete_dir_twice",
	// 			operations: []Operation{
	// 				// below "goodmorning.*" dirs are note affected
	// 				{true, processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
	// 				{true, DirectoryDelete{FilePath: "goodmorning/hello/world"}},
	// 				{false, DirectoryDelete{FilePath: "goodmorning/hello/world"}},
	// 			}, resultFile: "testdata/source_code/delete-directory5.json"},
	// 	})
	// })

	// t.Run("delete_file", func(t *testing.T) {
	// 	runEntries(t, []Entry{
	// 		{name: "delete_file_single",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello.txt"}},
	// 				{true, FileDelete{FilePath: "hello.txt"}},
	// 			}, resultFile: "testdata/source_code/new-source-code.json"},

	// 		{name: "delete_file_nested",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello/world/japan.txt"}},
	// 				{true, FileDelete{FilePath: "hello/world/japan.txt"}},
	// 			}, resultFile: "testdata/source_code/delete-file1.json"},

	// 		{name: "delete_file_next_to",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello/world/japan.txt"}},
	// 				{true, FileAdd{FilePath: "hello/world/america.txt"}},
	// 				{true, FileDelete{FilePath: "hello/world/japan.txt"}},
	// 			}, resultFile: "testdata/source_code/delete-file2.json"},

	// 		{name: "error_delete_file_non_existent",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello/world/japan.txt"}},
	// 				{true, FileAdd{FilePath: "hello/world/america.txt"}},
	// 				{false, FileDelete{FilePath: "hello/world/france.txt"}},
	// 			}, resultFile: "testdata/source_code/delete-file3.json"},

	// 		{name: "error_delete_file_twice",
	// 			operations: []Operation{
	// 				{true, FileAdd{FilePath: "hello/world/japan.txt"}},
	// 				{true, FileAdd{FilePath: "hello/world/america.txt"}},
	// 				{true, FileDelete{FilePath: "hello/world/japan.txt"}},
	// 				{false, FileDelete{FilePath: "hello/world/japan.txt"}},
	// 			}, resultFile: "testdata/source_code/delete-file2.json"},
	// 	})
	// })
}

// func TestSourceCode_Diff(t *testing.T) {
// 	type Operation struct {
// 		diff          Diff
// 		expectSuccess bool
// 	}

// 	type Entry struct {
// 		name       string
// 		operations []Operation
// 		resultFile string
// 	}

// 	applyOperations := func(processor *SourceCodeProcessor, operations []Operation) error {
// 		for opNum, op := range operations {
// 			opError := processor.ApplyDiff(op.diff)

// 			opSuccess := opError == nil
// 			if opSuccess != op.expectSuccess { //if result is not expected
// 				if op.expectSuccess {
// 					return fmt.Errorf(
// 						"operations[%d] success is expected, but result is failure\nerror:     %s\noperation: %+v\n",
// 						opNum,
// 						opError.Error(),
// 						op)
// 				} else {
// 					return fmt.Errorf(
// 						"operations[%d] failure is expected, but result is success\noperation: %+v\n",
// 						opNum,
// 						op)
// 				}
// 			}
// 		}

// 		return nil
// 	}

// 	// accept testEntries, and run tests on them
// 	runEntries := func(t *testing.T, testEntries []Entry) {
// 		for _, e := range testEntries {
// 			t.Run(e.name, func(t *testing.T) {
// 				sourceCode := NewSourceCodeProcessor()
// 				if err := applyOperations(sourceCode, e.operations); err != nil {
// 					t.Errorf(err.Error()) // fail but continue running, as two ops can fail in sequence
// 					return
// 				}

// 				internal.CompareAfterMarshal(t, e.resultFile, sourceCode.ToGraphQLModel())
// 			})
// 		}
// 	}

// 	t.Run("diff", func(t *testing.T) {
// 		runEntries(t, []Entry{
// 			{name: "add_file_single",
// 				operations: []Operation{
// 					{true, diff: Diff{FilesAdded: []FileAdd{{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}}}},
// 				}, resultFile: "testdata/source_code/diff-file1.json",
// 			},

// 			{name: "empty",
// 				operations: []Operation{
// 					{true, diff: Diff{}},
// 				}, resultFile: "testdata/source_code/new-source-code.json",
// 			},

// 			{name: "error_dupe",
// 				operations: []Operation{
// 					{false, diff: Diff{FilesAdded: []FileAdd{
// 						{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true},
// 						{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true},
// 					}}},
// 				}, resultFile: "testdata/source_code/new-source-code.json",
// 			},
// 		})
// 	})
// }

// func TestSourceCode_Mutation(t *testing.T) {
// 	sourceCode := NewSourceCodeProcessor()
// 	sourceCode.AddFile(FileAdd{FilePath: "hello/world/japan.txt"})
// 	sourceCode.AddFile(FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
// 	sourceCode.AddDirectory(processing.DirectoryAdd{FilePath: "goodmorning/hello/world"})

// 	// once materialized GraphQL model...
// 	materialized := sourceCode.ToGraphQLModel()
// 	internal.CompareAfterMarshal(t, "testdata/source_code/materialized.json", materialized)

// 	// ...mutation afterwards should have no effect
// 	sourceCode.step = "mutated step"
// 	sourceCode.defaultOpenFilePath = "mutated/file/path"
// 	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).content = "mutated content"
// 	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).filePath = "mutated/path/to/file"
// 	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).isUpdated = true
// 	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).filePath = "mutated/path/to/dir"
// 	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).isUpdated = false

// 	// materialized GraphQL model should not be affected
// 	internal.CompareAfterMarshal(t,
// 		"testdata/source_code/materialized.json",
// 		materialized) // by changing this to sourceCode.ToGraphQLModel(), you can confirm mutation actually updated sourceCode
// }

// func TestSourceCode_Clone(t *testing.T) {
// 	sourceCode := NewSourceCodeProcessor()
// 	sourceCode.AddFile(FileAdd{FilePath: "hello/world/japan.txt"})
// 	sourceCode.AddFile(FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
// 	sourceCode.AddDirectory(processing.DirectoryAdd{FilePath: "goodmorning/hello/world"})

// 	// once cloned ...
// 	sourceCodeCloned := sourceCode.Clone()
// 	internal.CompareAfterMarshal(t, "testdata/source_code/cloned.json", sourceCodeCloned.ToGraphQLModel())

// 	// ... mutation after cloning should have no effect
// 	sourceCode.step = "mutated step"
// 	sourceCode.defaultOpenFilePath = "mutated/file/path"
// 	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).content = "mutated content"
// 	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).filePath = "mutated/path/to/file"
// 	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).isUpdated = true
// 	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).filePath = "mutated/path/to/dir"
// 	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).isUpdated = false

// 	// cloned source code is not affected
// 	internal.CompareAfterMarshal(t,
// 		"testdata/source_code/cloned.json",
// 		sourceCodeCloned.ToGraphQLModel()) // by changing this to sourceCode.ToGraphQLModel(), you can confirm mutation actually updated sourceCode
// }
