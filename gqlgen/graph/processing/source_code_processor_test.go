package processing

import (
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

	// accept testEntries, and run tests on them
	testEntries := func(t *testing.T, testEntries []Entry) {
		for _, e := range testEntries {
			t.Run(e.name, func(t *testing.T) {
				processor := NewSourceCodeProcessor()
				for opNum, op := range e.operations {
					var err error
					switch v := op.operation.(type) {
					case model.DirectoryAdd:
						err = processor.AddDirectory(v)
					case model.DirectoryDelete:
						err = processor.DeleteDirectory(v)
					case model.FileAdd:
						err = processor.AddFile(v)
					case model.FileUpdate:
						err = processor.UpdateFile(v)
					case model.FileDelete:
						err = processor.DeleteFile(v)
					default:
						t.Fatalf("op %d faild:\nwrong operation type = %v", opNum, reflect.TypeOf(v))
						return
					}

					resultSuccess := err == nil
					if resultSuccess != op.expectSuccess { //if result is not expected
						if op.expectSuccess {
							t.Errorf("operation success is expected, but result is failure\nerror:     %s\noperation: %+v\n", err.Error(), op)
						} else {
							t.Errorf("operation failure is expected, but result is success\noperation: %+v\n", op)
						}
						return
					}
				}
			})
		}
	}

	t.Run("add_directory", func(t *testing.T) {
		//call testEntries inside func(t *testing.T) to reuse testEntries logic
		testEntries(t,
			[]Entry{
				{name: "delete_dir_single",
					operations: []Operation{
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
						{expectSuccess: true, operation: model.DirectoryDelete{FilePath: "hello"}},
						{expectSuccess: true, operation: model.DirectoryAdd{FilePath: "hello"}},
					}, resultFile: "testdata/source_code/new-source-code.json"}, // json should be same as initial state
			})
	})
}
