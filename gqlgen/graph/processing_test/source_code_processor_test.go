package processing_test

import (
	"fmt"
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

func Test_SourceCodeGitEffect(t *testing.T) {
	type TestCase struct {
		ExpectSuccess bool
		Effect        processing.SourceCodeGitEffect
		GoldenFile    string
	}

	checkResult := func(t *testing.T, index int, effect processing.SourceCodeGitEffect, expectSuccess bool, err error) {
		resultSuccess := err == nil
		if resultSuccess != expectSuccess { //if result is not expected
			if expectSuccess {
				t.Fatalf("effect[%d] = %+v success is expected, but result is failure\nerror: %s\n", index, effect, err)
			} else {
				t.Fatalf("effect[%d] = %+v failure is expected, but result is success\n", index, effect)
			}
		}
	}

	runTestCases := func(t *testing.T, repoUrl string, testCases []TestCase) {
		t.Run(repoUrl, func(t *testing.T) {
			sourceCode, err := processing.SourceCodeProcessorFromGit(repoUrl)
			if err != nil {
				t.Fatalf("failed to create source code processor: %s", err)
			}

			for i, c := range testCases {
				err := sourceCode.TransitionGit(fmt.Sprintf("%03d", i), c.Effect)
				checkResult(t, i, c.Effect, c.ExpectSuccess, err)

				internal.CompareWitGoldenFile(t, *updateFlag, c.GoldenFile, sourceCode.ToGraphQLModel())
			}
		})
	}

	runTestCases(t, "https://github.com/richardimaoka/gqlgensandbox", []TestCase{
		{true, processing.SourceCodeGitEffect{SeqNo: 0, CommitHash: "91a99d0c0558d2fc03c930d19afa97fc141f0c2e"}, "testdata/source_code_git/gqlgensandbox1.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 1, CommitHash: "83f8ad84dea56e8e5549832fb98eb8b5b9db4912"}, "testdata/source_code_git/gqlgensandbox2.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 2, CommitHash: "86a03f4f18b081b07e058f0e9f96503772a50cf0"}, "testdata/source_code_git/gqlgensandbox3.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 3, CommitHash: "490808086bded6b27f3651b095aefb7bb6708da2"}, "testdata/source_code_git/gqlgensandbox4.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 4, CommitHash: "9f835b8aaafdfc55933f52aae5a7c6e9864432aa"}, "testdata/source_code_git/gqlgensandbox5.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 5, CommitHash: "99a2c7f129cbebab3b17034fa93ad579d0fe29f6"}, "testdata/source_code_git/gqlgensandbox6.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 6, CommitHash: "20c5ef14fc6a0deae8a528beee3ed0f984da9ae1"}, "testdata/source_code_git/gqlgensandbox7.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 7, CommitHash: "4bc48072066d6e9fe339fae1341c196d4be03286"}, "testdata/source_code_git/gqlgensandbox8.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 8, CommitHash: "8d08178cb98df959288f2c4f8d0aff1bb20d6fc9"}, "testdata/source_code_git/gqlgensandbox9.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 9, CommitHash: "813c7822a3232c43edd9cc02286f063851ff2b54"}, "testdata/source_code_git/gqlgensandbox10.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 10, CommitHash: "a234864d58a12d50458ee563ba59c628c6861286"}, "testdata/source_code_git/gqlgensandbox11.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 11, CommitHash: "18c23ac5d49428845afe91f2d189968265afc19f"}, "testdata/source_code_git/gqlgensandbox12.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 12, CommitHash: "e02dc3bbdf21a533f1812507134cf1484a971f5b"}, "testdata/source_code_git/gqlgensandbox13.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 13, CommitHash: "929e04606a6eb7619f0e0949c2bdc2a1688a2d25"}, "testdata/source_code_git/gqlgensandbox14.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 14, CommitHash: "b08a390257a68951b2cf05a463655c852de7a4de"}, "testdata/source_code_git/gqlgensandbox15.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 15, CommitHash: "f745b8e233b2adfd11a63e7855f18a1986c7c084"}, "testdata/source_code_git/gqlgensandbox16.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 16, CommitHash: "700a1d749f1d3e86330ebe163d64a9fe58e25fd2"}, "testdata/source_code_git/gqlgensandbox17.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 17, CommitHash: "8c62836cfbbf9a9d0ce957d0edc4084e4bc88e60"}, "testdata/source_code_git/gqlgensandbox18.json"},
		{true, processing.SourceCodeGitEffect{SeqNo: 18, CommitHash: "4dd8f51d6acbee9d61b24dc26715ecc48a5d2456"}, "testdata/source_code_git/gqlgensandbox19.json"},
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
