package processing

import (
	"fmt"
	"reflect"
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
	testEntries := func(t *testing.T, testEntries []Entry) {
		for _, e := range testEntries {
			t.Run(e.name, func(t *testing.T) {
				processor := NewSourceCodeProcessor()
				if err := applyOperations(processor, e.operations); err != nil {
					t.Errorf(err.Error())
					return
				}

				compareAfterMarshal(t, e.resultFile, processor.ToSourceCode())
			})
		}
	}

	t.Run("add_directory", func(t *testing.T) {
		//reuse testEntries logic
		testEntries(t,
			[]Entry{
				// {name: "add_dir_single",
				// 	operations: []Operation{
				// 		{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
				// 	}, resultFile: "testdata/source_code/add-directory1.json"},

				{name: "add_dir_nested",
					operations: []Operation{
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world"}},
					}, resultFile: "testdata/source_code/add-directory2.json"},

				// {name: "error_add_dir_duplicate1",
				// 	operations: []Operation{
				// 		{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
				// 		{expectSuccess: false, operation: model.DirectoryAdd{FilePath: "hello"}},
				// 	}, resultFile: "testdata/source_code/add-directory1.json"},

				// {name: "add_dir_nested2",
				// 	operations: []Operation{
				// 		{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello/world"}},
				// 	}, resultFile: "testdata/source_code/add-directory3.json"},
			})
	})
}
