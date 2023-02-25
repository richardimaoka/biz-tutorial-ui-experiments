package model2

import (
	"fmt"
	"testing"
)

func statusString(expectSuccess bool) string {
	if expectSuccess {
		return "success"
	} else {
		return "failure"
	}
}

func TestAll(t *testing.T) {
	type Operation struct {
		filePath      string
		nodeType      FileNodeType
		expectSuccess bool
	}

	type Entry struct {
		name       string
		operations []Operation
		resultFile string
	}

	var entries []Entry = []Entry{
		{name: "create SourceCode",
			operations: []Operation{}, // no operation
			resultFile: "testdata/new-source-code.json"},

		{name: "error on adding dir with empty file path",
			operations: []Operation{
				{filePath: "", nodeType: FileNodeTypeDirectory, expectSuccess: false}, // "" is a wrong file path
			}, resultFile: "testdata/new-source-code.json"}, // json should be same as initial state

		{name: "add a single dir",
			operations: []Operation{
				{filePath: "hello", nodeType: FileNodeTypeDirectory, expectSuccess: true},
			}, resultFile: "testdata/add-directory1.json"},

		{name: "add a dir and its child dir",
			operations: []Operation{
				{filePath: "hello", nodeType: FileNodeTypeDirectory, expectSuccess: true},
				{filePath: "hello/world", nodeType: FileNodeTypeDirectory, expectSuccess: true},
			}, resultFile: "testdata/add-directory2.json"},

		{name: "add a dir and its child dir and another dir",
			operations: []Operation{
				{filePath: "hello", nodeType: FileNodeTypeDirectory, expectSuccess: true},
				{filePath: "hello/world", nodeType: FileNodeTypeDirectory, expectSuccess: true},
				{filePath: "aloha", nodeType: FileNodeTypeDirectory, expectSuccess: true},
			}, resultFile: "testdata/add-directory3.json"},

		{name: "add a file",
			operations: []Operation{
				{filePath: "hello.txt", nodeType: FileNodeTypeFile, expectSuccess: true},
			}, resultFile: "testdata/add-file1.json"},

		{name: "error adding a file",
			operations: []Operation{
				{filePath: "hello/world.txt", nodeType: FileNodeTypeFile, expectSuccess: false},
			}, resultFile: "testdata/new-source-code.json"},

		{name: "add two files",
			operations: []Operation{
				{filePath: "hello", nodeType: FileNodeTypeDirectory, expectSuccess: true},
				{filePath: "hello/world.txt", nodeType: FileNodeTypeFile, expectSuccess: true},
			}, resultFile: "testdata/add-file2.json"},
	}

	for i, e := range entries {
		t.Run(e.name, func(t *testing.T) {
			sc := newSourceCode()
			for j, op := range e.operations {
				var err error
				switch op.nodeType {
				case FileNodeTypeDirectory:
					err = sc.addDirectory(op.filePath)
				case FileNodeTypeFile:
					err = sc.addFile(op.filePath)
				default:
					t.Fatalf("entry %d, op %d faild:\nwrong op.nodeType = %s", i, j, op.nodeType)
					return
				}

				resultSuccess := err == nil
				if resultSuccess != op.expectSuccess {
					errMsg1 := fmt.Sprintf("operation %s is expected, but result is %s", statusString(op.expectSuccess), statusString(resultSuccess))
					errMsg2 := fmt.Sprintf("operation = %+v", op)
					errMsg3 := fmt.Sprintf("entry = %+v", e)
					t.Errorf("entry %d, op %d faild:\n%s\n%s\n%s", i, j, errMsg1, errMsg2, errMsg3)
					return
				}
			}

			compareAfterMarshal(t, e.resultFile, sc)
		})
	}
}
