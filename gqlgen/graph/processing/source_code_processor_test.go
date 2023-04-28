package processing

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

func Test_SourceCodeProcessor(t *testing.T) {
	type Operation struct {
		operation     model.FileSystemOperation
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
			case model.DirectoryAdd:
				opError = processor.AddDirectory(v)
			case model.DirectoryDelete:
				opError = processor.DeleteDirectory(v)
			case model.FileAdd:
				opError = processor.AddFile(v)
			case model.FileUpdate:
				opError = processor.UpdateFile(v)
			case model.FileDelete:
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
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
					}, resultFile: "testdata/source_code/add-directory1.json"},

				{name: "add_dir_nested",
					operations: []Operation{
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world"}},
					}, resultFile: "testdata/source_code/add-directory2.json"},

				{name: "add_dir_nested2",
					operations: []Operation{
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world"}},
					}, resultFile: "testdata/source_code/add-directory3.json"},

				{name: "add_dir_multiple",
					operations: []Operation{
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world"}},
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "aloha"}},
					}, resultFile: "testdata/source_code/add-directory4.json"},

				{name: "error_add_dir_empty",
					operations: []Operation{
						{expectSuccess: false, operation: model.DirectoryAdd{FilePath: ""}}, // "" is a wrong file path
					}, resultFile: "testdata/source_code/new-source-code.json"}, // json should be same as initial state

				{name: "error_add_dir_duplicate1",
					operations: []Operation{
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
						{expectSuccess: false, operation: model.DirectoryAdd{FilePath: "hello"}},
					}, resultFile: "testdata/source_code/add-directory1.json"},

				{name: "add_dir_nested2",
					operations: []Operation{
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world"}},
					}, resultFile: "testdata/source_code/add-directory3.json"},
			})
	})

	t.Run("add_file", func(t *testing.T) {
		runEntries(t, []Entry{
			{name: "add_file_single",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello.txt", Content: "hello new world", IsFullContent: true}},
				}, resultFile: "testdata/source_code/add-file1.json"},

			{name: "add_file_nested",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/new.txt", Content: "hello new world", IsFullContent: true}},
				}, resultFile: "testdata/source_code/add-file2.json"},

			{name: "add_file_nested2",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/japan.txt"}},
				}, resultFile: "testdata/source_code/add-file3.json"},

			{name: "add_file_next_to",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/japan.txt"}},
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/america.txt"}},
				}, resultFile: "testdata/source_code/add-file4.json"},

			{name: "error_add_file_duplicate1",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello.txt", Content: "hello new world"}},
					{expectSuccess: false, operation: model.FileAdd{FilePath: "hello.txt"}},
				}, resultFile: "testdata/source_code/add-file1.json"},

			{name: "error_add_file_duplicate2",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world.txt"}},
					{expectSuccess: false, operation: model.FileAdd{FilePath: "hello/world.txt"}},
					{expectSuccess: false, operation: model.FileAdd{FilePath: "hello"}},
				}, resultFile: "testdata/source_code/add-file5.json"},
		})
	})

	t.Run("delete_directory", func(t *testing.T) {
		runEntries(t, []Entry{
			{name: "delete_dir_single",
				operations: []Operation{
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
					{expectSuccess: true, operation: model.DirectoryDelete{FilePath: "hello"}},
				}, resultFile: "testdata/source_code/new-source-code.json"}, // json should be same as initial state

			{name: "delete_dir_nested_leaf",
				operations: []Operation{
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world"}},
					{expectSuccess: true, operation: model.DirectoryDelete{FilePath: "hello/world"}},
				}, resultFile: "testdata/source_code/delete-directory1.json"},

			{name: "delete_dir_nested_middle",
				operations: []Operation{
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world/japan"}},
					// below "goodmorning.*" dirs are not affected
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning/hello"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
					{expectSuccess: true, operation: model.DirectoryDelete{FilePath: "hello/world"}},
				}, resultFile: "testdata/source_code/delete-directory2.json"},

			{name: "delete_dir_nested_parent",
				operations: []Operation{
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world/japan"}},
					// below "goodmorning.*" dirs are note affected
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning/hello"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
					{expectSuccess: true, operation: model.DirectoryDelete{FilePath: "hello"}},
				}, resultFile: "testdata/source_code/delete-directory3.json"},

			{name: "error_delete_dir_non_existent",
				operations: []Operation{
					// below "goodmorning.*" dirs are note affected
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning/hello"}},
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
					{expectSuccess: false, operation: model.DirectoryDelete{FilePath: "goodmorning/hello/universe"}},
					{expectSuccess: false, operation: model.DirectoryDelete{FilePath: "goodmorning/vonjour/world"}},
				}, resultFile: "testdata/source_code/delete-directory4.json"},

			{name: "error_delete_dir_twice",
				operations: []Operation{
					// below "goodmorning.*" dirs are note affected
					{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "goodmorning/hello/world"}},
					{expectSuccess: true, operation: model.DirectoryDelete{FilePath: "goodmorning/hello/world"}},
					{expectSuccess: false, operation: model.DirectoryDelete{FilePath: "goodmorning/hello/world"}},
				}, resultFile: "testdata/source_code/delete-directory5.json"},
		})
	})

	t.Run("delete_file", func(t *testing.T) {
		runEntries(t, []Entry{
			{name: "delete_file_single",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello.txt"}},
					{expectSuccess: true, operation: model.FileDelete{FilePath: "hello.txt"}},
				}, resultFile: "testdata/source_code/new-source-code.json"},

			{name: "delete_file_nested",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/japan.txt"}},
					{expectSuccess: true, operation: model.FileDelete{FilePath: "hello/world/japan.txt"}},
				}, resultFile: "testdata/source_code/delete-file1.json"},

			{name: "delete_file_next_to",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/japan.txt"}},
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/america.txt"}},
					{expectSuccess: true, operation: model.FileDelete{FilePath: "hello/world/japan.txt"}},
				}, resultFile: "testdata/source_code/delete-file2.json"},

			{name: "error_delete_file_non_existent",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/japan.txt"}},
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/america.txt"}},
					{expectSuccess: false, operation: model.FileDelete{FilePath: "hello/world/france.txt"}},
				}, resultFile: "testdata/source_code/delete-file3.json"},

			{name: "error_delete_file_twice",
				operations: []Operation{
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/japan.txt"}},
					{expectSuccess: true, operation: model.FileAdd{FilePath: "hello/world/america.txt"}},
					{expectSuccess: true, operation: model.FileDelete{FilePath: "hello/world/japan.txt"}},
					{expectSuccess: false, operation: model.FileDelete{FilePath: "hello/world/japan.txt"}},
				}, resultFile: "testdata/source_code/delete-file2.json"},
		})
	})
}

func TestSourceCode_Mutation(t *testing.T) {

	sourceCode := NewSourceCodeProcessor()
	sourceCode.AddFile(model.FileAdd{FilePath: "hello/world/japan.txt"})
	sourceCode.AddFile(model.FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
	sourceCode.AddDirectory(model.DirectoryAdd{FilePath: "goodmorning/hello/world"})

	// once materialized GraphQL model, mutation afterwards should have no effect
	result := sourceCode.ToGraphQLModel()
	compareAfterMarshal(t, "testdata/source_code/mutated.json", result)

	sourceCode.step = "mutated step"
	sourceCode.defaultOpenFilePath = "mutated/file/path"
	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).content = "mutated content"
	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).filePath = "mutated/path/to/file"
	sourceCode.fileMap["hello/world/japan.txt"].(*fileProcessorNode).isUpdated = true
	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).filePath = "mutated/path/to/dir"
	sourceCode.fileMap["goodmorning/hello/world"].(*directoryProcessorNode).isUpdated = false

	compareAfterMarshal(t, "testdata/source_code/mutated.json", result)
}

func testSourceCodeEqual(t *testing.T, sc1, sc2 *SourceCodeProcessor) {
	errors := []string{}

	if sc1.defaultOpenFilePath != sc2.defaultOpenFilePath {
		errors = append(errors, fmt.Sprintf("defaultOpenFilePath not equal: %s != %s", sc1.defaultOpenFilePath, sc2.defaultOpenFilePath))
	}

	if sc1.step != sc2.step {
		errors = append(errors, fmt.Sprintf("step not equal: %s != %s", sc1.step, sc2.step))
	}

	fileMapComparison := func(side string, fileMap1, fileMap2 map[string]fileTreeNode) {
		for k, v1 := range fileMap1 {
			v2, ok := fileMap2[k]
			if !ok {
				errors = append(errors, fmt.Sprintf("fileMap[%s] not found on %s", k, side))
				continue
			}

			switch vv1 := v1.(type) {
			case *fileProcessorNode:
				vv2, ok := v2.(*fileProcessorNode)
				if !ok {
					errors = append(errors, fmt.Sprintf("fileMap[%s] is not a fileProcessorNode on %s", k, side))
				} else if *vv1 != *vv2 {
					errors = append(errors, fmt.Sprintf("fileMap[%s] not equal: %v != %v", k, vv1, vv2))
				}
			case *directoryProcessorNode:
				vv2, ok := v2.(*directoryProcessorNode)
				if !ok {
					errors = append(errors, fmt.Sprintf("fileMap[%s] is not a directoryProcessorNode on %s", k, side))
				} else if *vv1 != *vv2 {
					errors = append(errors, fmt.Sprintf("fileMap[%s] not equal: %v != %v", k, vv1, vv2))
				}
			}
		}
	}

	fileMapComparison("sc2", sc1.fileMap, sc2.fileMap)
	fileMapComparison("sc1", sc2.fileMap, sc1.fileMap)

	if len(errors) > 0 {
		t.Fatalf("%s", strings.Join(errors, ", "))
	}
}

func TestSourceCode_Clone(t *testing.T) {
	sourceCode := NewSourceCodeProcessor()
	sourceCode.AddFile(model.FileAdd{FilePath: "hello/world/japan.txt"})
	sourceCode.AddFile(model.FileAdd{FilePath: "hello/world/america.txt", Content: "hello usa", IsFullContent: true})
	sourceCode.AddDirectory(model.DirectoryAdd{FilePath: "goodmorning/hello/world"})

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
