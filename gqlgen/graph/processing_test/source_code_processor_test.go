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
				t.Fatalf("operation[%d] = %+v success is expected, but result is failure\nerror: %s\n", index, op, err)
			} else {
				t.Fatalf("operation[%d] = %+v failure is expected, but result is success\n", index, op)
			}
		}
	}

	// accept testEntries, and run tests on them
	runTestCases := func(t *testing.T, testCases []TestCase) {
		sourceCode := processing.NewSourceCodeProcessor()
		for i, c := range testCases {
			err := applyOperation(t, sourceCode, c.Operation)
			checkResult(t, i, c.Operation, c.ExpectSuccess, err)

			// if `-update` flag is passed from command line, update the golden file
			if *update {
				internal.WriteGoldenFile(t, c.ExpectedFile, sourceCode.ToGraphQLModel())
			}

			internal.CompareAfterMarshal(t, c.ExpectedFile, sourceCode.ToGraphQLModel())
		}
	}

	t.Run("add_directory1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
		})
	})

	t.Run("add_dir_nested", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-dir-nested1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
			{true, "testdata/source_code/add-dir-nested1-2.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		})
	})

	t.Run("add_dir_nested2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-directory3.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		})
	})

	t.Run("add_dir_multiple", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add_dir_multiple1-1.json", processing.DirectoryAdd{FilePath: "hello/world"}},
			{true, "testdata/source_code/add_dir_multiple1-2.json", processing.DirectoryAdd{FilePath: "aloha"}},
		})
	})

	t.Run("add_dir_multiple", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-directory4-1.json", processing.DirectoryAdd{FilePath: "hello/world"}},
			{true, "testdata/source_code/add-directory4-2.json", processing.DirectoryAdd{FilePath: "aloha"}},
		})
	})

	t.Run("", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{false, "testdata/source_code/new-source-code.json", processing.DirectoryAdd{FilePath: ""}}, // "" is a wrong file path

		})
	})

	t.Run("error_add_dir_duplicate1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
			{false, "testdata/source_code/add-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
		})
	})

	t.Run("add_dir_nested2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-directory3.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		})
	})

	t.Run("add_file_single", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-file1.json", processing.FileAdd{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}},
		})
	})

	t.Run("add_file_nested1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-file2.json", processing.FileAdd{FilePath: "hello/world/new.txt", Content: "hello new world", IsFullContent: true}},
		})
	})

	t.Run("add_file_nested2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-file3.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
		})
	})

	t.Run("add_file_next_to", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-file4-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
			{true, "testdata/source_code/add-file4-2.json", processing.FileAdd{FilePath: "hello/world/america.txt"}},
		})
	})

	t.Run("error_add_file_duplicate1", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-file1.json", processing.FileAdd{FilePath: "hello.txt", Content: "hello new world"}},
			{false, "testdata/source_code/add-file1.json", processing.FileAdd{FilePath: "hello.txt"}},
		})
	})

	t.Run("error_add_file_duplicate2", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/add-file5-1.json", processing.FileAdd{FilePath: "hello/world.txt"}},
			{false, "testdata/source_code/add-file5-1.json", processing.FileAdd{FilePath: "hello/world.txt"}},
			{false, "testdata/source_code/add-file5-1.json", processing.FileAdd{FilePath: "hello"}},
		})
	})

	t.Run("delete_dir_single", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
			{true, "testdata/source_code/new-source-code.json", processing.DirectoryDelete{FilePath: "hello"}},
		})
	})

	t.Run("delete_dir_nested_leaf", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
			{true, "testdata/source_code/delete-directory1-2.json", processing.DirectoryAdd{FilePath: "hello/world"}},
			{true, "testdata/source_code/delete-directory1-3.json", processing.DirectoryDelete{FilePath: "hello/world"}},
		})
	})

	t.Run("delete_dir_nested_middle", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-directory2-1.json", processing.DirectoryAdd{FilePath: "hello"}},
			{true, "testdata/source_code/delete-directory2-2.json", processing.DirectoryAdd{FilePath: "hello/world"}},
			{true, "testdata/source_code/delete-directory2-3.json", processing.DirectoryAdd{FilePath: "hello/world/japan"}},
			// below "goodmorning.*" dirs are not affected
			{true, "testdata/source_code/delete-directory2-4.json", processing.DirectoryAdd{FilePath: "goodmorning"}},
			{true, "testdata/source_code/delete-directory2-5.json", processing.DirectoryAdd{FilePath: "goodmorning/hello"}},
			{true, "testdata/source_code/delete-directory2-6.json", processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
			{true, "testdata/source_code/delete-directory2-7.json", processing.DirectoryDelete{FilePath: "hello/world"}},
		})
	})

	t.Run("delete_dir_nested_parent", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-directory3-1.json", processing.DirectoryAdd{FilePath: "hello"}},
			{true, "testdata/source_code/delete-directory3-2.json", processing.DirectoryAdd{FilePath: "hello/world"}},
			{true, "testdata/source_code/delete-directory3-3.json", processing.DirectoryAdd{FilePath: "hello/world/japan"}},
			// below "goodmorning.*" dirs are note affected
			{true, "testdata/source_code/delete-directory3-4.json", processing.DirectoryAdd{FilePath: "goodmorning"}},
			{true, "testdata/source_code/delete-directory3-5.json", processing.DirectoryAdd{FilePath: "goodmorning/hello"}},
			{true, "testdata/source_code/delete-directory3-6.json", processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
			{true, "testdata/source_code/delete-directory3-7.json", processing.DirectoryDelete{FilePath: "hello"}},
		})
	})

	t.Run("error_delete_dir_non_existent", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-directory4-1.json", processing.DirectoryAdd{FilePath: "goodmorning"}},
			{true, "testdata/source_code/delete-directory4-2.json", processing.DirectoryAdd{FilePath: "goodmorning/hello"}},
			{true, "testdata/source_code/delete-directory4-3.json", processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
			{false, "testdata/source_code/delete-directory4-3.json", processing.DirectoryDelete{FilePath: "goodmorning/hello/universe"}},
			{false, "testdata/source_code/delete-directory4-3.json", processing.DirectoryDelete{FilePath: "goodmorning/vonjour/world"}},
		})
	})

	t.Run("error_delete_dir_twice", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-directory5-1.json", processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
			{true, "testdata/source_code/delete-directory5-2.json", processing.DirectoryDelete{FilePath: "goodmorning/hello/world"}},
			{false, "testdata/source_code/delete-directory5-2.json", processing.DirectoryDelete{FilePath: "goodmorning/hello/world"}},
		})
	})

	t.Run("delete_file_single", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-file1.json", processing.FileAdd{FilePath: "hello.txt"}},
			{true, "testdata/source_code/new-source-code.json", processing.FileDelete{FilePath: "hello.txt"}},
		})
	})

	t.Run("delete_file_nested", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-file2-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
			{true, "testdata/source_code/delete-file2-2.json", processing.FileDelete{FilePath: "hello/world/japan.txt"}},
		})
	})

	t.Run("delete_file_next_to", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-file3-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
			{true, "testdata/source_code/delete-file3-2.json", processing.FileAdd{FilePath: "hello/world/america.txt"}},
			{true, "testdata/source_code/delete-file3-3.json", processing.FileDelete{FilePath: "hello/world/japan.txt"}},
		})
	})

	t.Run("error_delete_file_non_existent", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-file4-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
			{true, "testdata/source_code/delete-file4-2.json", processing.FileAdd{FilePath: "hello/world/america.txt"}},
			{false, "testdata/source_code/delete-file4-3.json", processing.FileDelete{FilePath: "hello/world/france.txt"}},
		})
	})

	t.Run("error_delete_file_twice", func(t *testing.T) {
		runTestCases(t, []TestCase{
			{true, "testdata/source_code/delete-file5-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
			{true, "testdata/source_code/delete-file5-2.json", processing.FileAdd{FilePath: "hello/world/america.txt"}},
			{true, "testdata/source_code/delete-file5-3.json", processing.FileDelete{FilePath: "hello/world/japan.txt"}},
			{false, "testdata/source_code/delete-file5-3.json", processing.FileDelete{FilePath: "hello/world/japan.txt"}},
		})
	})
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

func TestSourceCode_Mutation(t *testing.T) {
	sourceCode := processing.NewSourceCodeProcessor()
	sourceCode.AddFile(processing.FileAdd{FilePath: "hello/world/japan.txt"})
	sourceCode.AddFile(processing.FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
	sourceCode.AddDirectory(processing.DirectoryAdd{FilePath: "goodmorning/hello/world"})

	// once GraphQL model is materialized...
	materialized := sourceCode.ToGraphQLModel()
	if *update { // if `-update` flag is passed from command line, update the golden file
		internal.WriteGoldenFile(t, "testdata/source_code/mutation1-1.json", materialized)
	}
	internal.CompareAfterMarshal(t, "testdata/source_code/mutation1-1.json", materialized)

	// ...mutation to source code...
	sourceCode.AddFile(processing.FileAdd{FilePath: "aloha/world/germany.txt"})
	sourceCode.AddFile(processing.FileAdd{FilePath: "aloha/world/uk.txt", Content: "hello britain", IsFullContent: true})
	sourceCode.DeleteFile(processing.FileDelete{FilePath: "hello/world/japan.txt"})

	// ...should of course have effect on re-materialized GraphQL model
	if *update { // if `-update` flag is passed from command line, update the golden file
		internal.WriteGoldenFile(t, "testdata/source_code/mutation1-2.json", sourceCode.ToGraphQLModel())
	}
	internal.CompareAfterMarshal(t, "testdata/source_code/mutation1-2.json", sourceCode.ToGraphQLModel())

	// ...but should have no effect on materialized GraphQL model
	internal.CompareAfterMarshal(t, "testdata/source_code/mutation1-1.json", materialized)
}

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
