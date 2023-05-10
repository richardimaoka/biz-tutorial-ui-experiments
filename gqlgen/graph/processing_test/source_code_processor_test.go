package processing_test

import (
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

	runTestCases := func(t *testing.T, name string, testCases []TestCase) {
		t.Run(name, func(t *testing.T) {
			sourceCode := processing.NewSourceCodeProcessor()
			for i, c := range testCases {
				err := sourceCode.ApplyOperation(c.Operation)
				checkResult(t, i, c.Operation, c.ExpectSuccess, err)

				internal.CompareWitGoldenFile(t, *updateFlag, c.ExpectedFile, sourceCode.ToGraphQLModel())
			}
		})
	}

	runTestCases(t, "add_directory1", []TestCase{
		{true, "testdata/source_code/add-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
	})

	runTestCases(t, "add_dir_nested", []TestCase{
		{true, "testdata/source_code/add-dir-nested1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
		{true, "testdata/source_code/add-dir-nested1-2.json", processing.DirectoryAdd{FilePath: "hello/world"}},
	})

	runTestCases(t, "add_dir_nested2", []TestCase{
		{true, "testdata/source_code/add-directory3.json", processing.DirectoryAdd{FilePath: "hello/world"}},
	})

	runTestCases(t, "add_dir_multiple", []TestCase{
		{true, "testdata/source_code/add_dir_multiple1-1.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		{true, "testdata/source_code/add_dir_multiple1-2.json", processing.DirectoryAdd{FilePath: "aloha"}},
	})

	runTestCases(t, "add_dir_multiple", []TestCase{
		{true, "testdata/source_code/add-directory4-1.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		{true, "testdata/source_code/add-directory4-2.json", processing.DirectoryAdd{FilePath: "aloha"}},
	})

	runTestCases(t, "errro_empty_file_path", []TestCase{
		{false, "testdata/source_code/new-source-code.json", processing.DirectoryAdd{FilePath: ""}}, // "" is a wrong file path

	})

	runTestCases(t, "error_add_dir_duplicate1", []TestCase{
		{true, "testdata/source_code/add-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
		{false, "testdata/source_code/add-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
	})

	runTestCases(t, "add_dir_nested2", []TestCase{
		{true, "testdata/source_code/add-directory3.json", processing.DirectoryAdd{FilePath: "hello/world"}},
	})

	runTestCases(t, "add_file_single", []TestCase{
		{true, "testdata/source_code/add-file1.json", processing.FileAdd{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}},
	})

	runTestCases(t, "add_file_nested1", []TestCase{
		{true, "testdata/source_code/add-file2.json", processing.FileAdd{FilePath: "hello/world/new.txt", Content: "hello new world", IsFullContent: true}},
	})

	runTestCases(t, "add_file_nested2", []TestCase{
		{true, "testdata/source_code/add-file3.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
	})

	runTestCases(t, "add_file_next_to", []TestCase{
		{true, "testdata/source_code/add-file4-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
		{true, "testdata/source_code/add-file4-2.json", processing.FileAdd{FilePath: "hello/world/america.txt"}},
	})

	runTestCases(t, "error_add_file_duplicate1", []TestCase{
		{true, "testdata/source_code/add-file1.json", processing.FileAdd{FilePath: "hello.txt", Content: "hello new world"}},
		{false, "testdata/source_code/add-file1.json", processing.FileAdd{FilePath: "hello.txt"}},
	})

	runTestCases(t, "error_add_file_duplicate2", []TestCase{
		{true, "testdata/source_code/add-file5-1.json", processing.FileAdd{FilePath: "hello/world.txt"}},
		{false, "testdata/source_code/add-file5-1.json", processing.FileAdd{FilePath: "hello/world.txt"}},
		{false, "testdata/source_code/add-file5-1.json", processing.FileAdd{FilePath: "hello"}},
	})

	runTestCases(t, "delete_dir_single", []TestCase{
		{true, "testdata/source_code/delete-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
		{true, "testdata/source_code/new-source-code.json", processing.DirectoryDelete{FilePath: "hello"}},
	})

	runTestCases(t, "delete_dir_nested_leaf", []TestCase{
		{true, "testdata/source_code/delete-directory1-1.json", processing.DirectoryAdd{FilePath: "hello"}},
		{true, "testdata/source_code/delete-directory1-2.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		{true, "testdata/source_code/delete-directory1-3.json", processing.DirectoryDelete{FilePath: "hello/world"}},
	})

	runTestCases(t, "delete_dir_nested_middle", []TestCase{
		{true, "testdata/source_code/delete-directory2-1.json", processing.DirectoryAdd{FilePath: "hello"}},
		{true, "testdata/source_code/delete-directory2-2.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		{true, "testdata/source_code/delete-directory2-3.json", processing.DirectoryAdd{FilePath: "hello/world/japan"}},
		// below "goodmorning.*" dirs are not affected
		{true, "testdata/source_code/delete-directory2-4.json", processing.DirectoryAdd{FilePath: "goodmorning"}},
		{true, "testdata/source_code/delete-directory2-5.json", processing.DirectoryAdd{FilePath: "goodmorning/hello"}},
		{true, "testdata/source_code/delete-directory2-6.json", processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
		{true, "testdata/source_code/delete-directory2-7.json", processing.DirectoryDelete{FilePath: "hello/world"}},
	})

	runTestCases(t, "delete_dir_nested_parent", []TestCase{
		{true, "testdata/source_code/delete-directory3-1.json", processing.DirectoryAdd{FilePath: "hello"}},
		{true, "testdata/source_code/delete-directory3-2.json", processing.DirectoryAdd{FilePath: "hello/world"}},
		{true, "testdata/source_code/delete-directory3-3.json", processing.DirectoryAdd{FilePath: "hello/world/japan"}},
		// below "goodmorning.*" dirs are note affected
		{true, "testdata/source_code/delete-directory3-4.json", processing.DirectoryAdd{FilePath: "goodmorning"}},
		{true, "testdata/source_code/delete-directory3-5.json", processing.DirectoryAdd{FilePath: "goodmorning/hello"}},
		{true, "testdata/source_code/delete-directory3-6.json", processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
		{true, "testdata/source_code/delete-directory3-7.json", processing.DirectoryDelete{FilePath: "hello"}},
	})

	runTestCases(t, "error_delete_dir_non_existent", []TestCase{
		{true, "testdata/source_code/delete-directory4-1.json", processing.DirectoryAdd{FilePath: "goodmorning"}},
		{true, "testdata/source_code/delete-directory4-2.json", processing.DirectoryAdd{FilePath: "goodmorning/hello"}},
		{true, "testdata/source_code/delete-directory4-3.json", processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
		{false, "testdata/source_code/delete-directory4-3.json", processing.DirectoryDelete{FilePath: "goodmorning/hello/universe"}},
		{false, "testdata/source_code/delete-directory4-3.json", processing.DirectoryDelete{FilePath: "goodmorning/vonjour/world"}},
	})

	runTestCases(t, "error_delete_dir_twice", []TestCase{
		{true, "testdata/source_code/delete-directory5-1.json", processing.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
		{true, "testdata/source_code/delete-directory5-2.json", processing.DirectoryDelete{FilePath: "goodmorning/hello/world"}},
		{false, "testdata/source_code/delete-directory5-2.json", processing.DirectoryDelete{FilePath: "goodmorning/hello/world"}},
	})

	runTestCases(t, "delete_file_single", []TestCase{
		{true, "testdata/source_code/delete-file1.json", processing.FileAdd{FilePath: "hello.txt"}},
		{true, "testdata/source_code/new-source-code.json", processing.FileDelete{FilePath: "hello.txt"}},
	})

	runTestCases(t, "delete_file_nested", []TestCase{
		{true, "testdata/source_code/delete-file2-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
		{true, "testdata/source_code/delete-file2-2.json", processing.FileDelete{FilePath: "hello/world/japan.txt"}},
	})

	runTestCases(t, "delete_file_next_to", []TestCase{
		{true, "testdata/source_code/delete-file3-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
		{true, "testdata/source_code/delete-file3-2.json", processing.FileAdd{FilePath: "hello/world/america.txt"}},
		{true, "testdata/source_code/delete-file3-3.json", processing.FileDelete{FilePath: "hello/world/japan.txt"}},
	})

	runTestCases(t, "error_delete_file_non_existent", []TestCase{
		{true, "testdata/source_code/delete-file4-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
		{true, "testdata/source_code/delete-file4-2.json", processing.FileAdd{FilePath: "hello/world/america.txt"}},
		{false, "testdata/source_code/delete-file4-3.json", processing.FileDelete{FilePath: "hello/world/france.txt"}},
	})

	runTestCases(t, "error_delete_file_twice", []TestCase{
		{true, "testdata/source_code/delete-file5-1.json", processing.FileAdd{FilePath: "hello/world/japan.txt"}},
		{true, "testdata/source_code/delete-file5-2.json", processing.FileAdd{FilePath: "hello/world/america.txt"}},
		{true, "testdata/source_code/delete-file5-3.json", processing.FileDelete{FilePath: "hello/world/japan.txt"}},
		{false, "testdata/source_code/delete-file5-3.json", processing.FileDelete{FilePath: "hello/world/japan.txt"}},
	})
}

func TestSourceCode_Diff(t *testing.T) {
	checkResult := func(t *testing.T, index int, diff processing.Diff, expectSuccess bool, err error) {
		resultSuccess := err == nil
		if resultSuccess != expectSuccess { //if result is not expected
			if expectSuccess {
				t.Fatalf("diff[%d] = %+v success is expected, but result is failure\nerror: %s\n", index, diff, err)
			} else {
				t.Fatalf("diff[%d] = %+v failure is expected, but result is success\n", index, diff)
			}
		}
	}

	type TestCase struct {
		ExpectSuccess bool
		ExpectedFile  string
		Diff          processing.Diff
	}

	runTestCases := func(t *testing.T, name string, testCases []TestCase) {
		t.Run(name, func(t *testing.T) {
			sourceCode := processing.NewSourceCodeProcessor()
			for i, c := range testCases {
				err := sourceCode.ApplyDiff(c.Diff)
				checkResult(t, i, c.Diff, c.ExpectSuccess, err)
				internal.CompareWitGoldenFile(t, *updateFlag, c.ExpectedFile, sourceCode.ToGraphQLModel())
			}
		})
	}

	runTestCases(t, "add_file_single", []TestCase{
		{true, "testdata/source_code/diff-file1.json", processing.Diff{FilesAdded: []processing.FileAdd{{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}}}},
	})

	runTestCases(t, "empty_diff", []TestCase{
		{true, "testdata/source_code/new-source-code.json", processing.Diff{}},
	})

	runTestCases(t, "error_diff_dupe", []TestCase{
		{false, "testdata/source_code/new-source-code.json", processing.Diff{
			FilesAdded: []processing.FileAdd{
				{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true},
				{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true},
			},
		}},
	})
}

// Test mutation after sourceCode.ToGraphQLModel()
// Once a GraphQL model is materialized with sourceCode.ToGraphQLModel(), mutation to the sourceCode should have no effect on the GraphQL model
func TestSourceCode_Mutation1(t *testing.T) {
	sourceCode := processing.NewSourceCodeProcessor()
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "hello/world/japan.txt"})
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
	sourceCode.ApplyOperation(processing.DirectoryAdd{FilePath: "goodmorning/hello/world"})

	// once GraphQL model is materialized...
	materialized := sourceCode.ToGraphQLModel()
	internal.CompareWitGoldenFile(t, *updateFlag, "testdata/source_code/mutation1-1.json", materialized)

	// ...mutation to source code...
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "aloha/world/germany.txt"})
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "aloha/world/uk.txt", Content: "hello britain", IsFullContent: true})
	sourceCode.ApplyOperation(processing.FileDelete{FilePath: "hello/world/japan.txt"})

	// ...should of course have effect on re-materialized GraphQL model
	internal.CompareWitGoldenFile(t, *updateFlag, "testdata/source_code/mutation1-2.json", sourceCode.ToGraphQLModel())

	// ...but should have no effect on materialized GraphQL model
	internal.CompareAfterMarshal(t, "testdata/source_code/mutation1-1.json", materialized)
}

// Test mutation after sourceCode.ToGraphQLModel()
// Once a GraphQL model is materialized with sourceCode.ToGraphQLModel(), mutation to the GraphQL model should have no effect on the sourceCode
func TestSourceCode_Mutation2(t *testing.T) {
	sourceCode := processing.NewSourceCodeProcessor()
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "hello/world/japan.txt"})
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
	sourceCode.ApplyOperation(processing.DirectoryAdd{FilePath: "goodmorning/hello/world"})

	// once GraphQL model is materialized...
	materialized := sourceCode.ToGraphQLModel()
	internal.CompareWitGoldenFile(t, *updateFlag, "testdata/source_code/mutation2-1.json", materialized)

	// ...mutation to materialized GraphQL model...
	ptr1 := materialized.FileTree[0].Name
	*ptr1 = "mutation extra name"
	ptr2 := materialized.FileTree[0].FilePath
	*ptr2 = "mutation/extra/path"

	// ...should of course have effect on materialized GraphQL model
	internal.CompareWitGoldenFile(t, *updateFlag, "testdata/source_code/mutation2-2.json", materialized)

	// ...but should have no effect on source code
	internal.CompareAfterMarshal(t, "testdata/source_code/mutation2-1.json", sourceCode.ToGraphQLModel())
}

func TestSourceCode_Clone(t *testing.T) {
	sourceCode := processing.NewSourceCodeProcessor()
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "hello/world/japan.txt"})
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
	sourceCode.ApplyOperation(processing.DirectoryAdd{FilePath: "goodmorning/hello/world"})

	// once cloned ...
	sourceCodeCloned := sourceCode.Clone()
	internal.CompareWitGoldenFile(t, *updateFlag, "testdata/source_code/clone1-1.json", sourceCodeCloned.ToGraphQLModel())

	// ... mutation to source code
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "aloha/world/germany.txt"})
	sourceCode.ApplyOperation(processing.FileAdd{FilePath: "aloha/world/uk.txt", Content: "hello britain", IsFullContent: true})
	sourceCode.ApplyOperation(processing.FileDelete{FilePath: "hello/world/japan.txt"})

	// ...should of course have effect on terminal itself
	internal.CompareWitGoldenFile(t, *updateFlag, "testdata/source_code/clone1-2.json", sourceCode.ToGraphQLModel())

	// ...but should have no effect on sourceCode
	internal.CompareAfterMarshal(t, "testdata/source_code/clone1-1.json", sourceCodeCloned.ToGraphQLModel())
}
