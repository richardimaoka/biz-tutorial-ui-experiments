package model2

import (
	"fmt"
	"testing"
)

type OperationType string

const (
	OpAddDirectory OperationType = "Add Directory"
	OpAddFile      OperationType = "Add File"
)

type Operation struct {
	filePath      string
	operationType OperationType
	expectSuccess bool
}

type Entry struct {
	name       string
	operations []Operation
	resultFile string
}

func statusString(expectSuccess bool) string {
	if expectSuccess {
		return "success"
	} else {
		return "failure"
	}
}

func TestDirectoryCases(t *testing.T) {
	var entries []Entry = []Entry{
		{name: "create SourceCode",
			operations: []Operation{}, // no operation
			resultFile: "testdata/new-source-code.json"},

		{name: "error on adding dir with empty file path",
			operations: []Operation{
				{filePath: "", operationType: OpAddDirectory, expectSuccess: false}, // "" is a wrong file path
			}, resultFile: "testdata/new-source-code.json"}, // json should be same as initial state

		{name: "add a single dir",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
			}, resultFile: "testdata/add-directory1.json"},

		{name: "add a dir and its child dir",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world", operationType: OpAddDirectory, expectSuccess: true},
			}, resultFile: "testdata/add-directory2.json"},

		{name: "add a dir and its child dir and another dir",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "aloha", operationType: OpAddDirectory, expectSuccess: true},
			}, resultFile: "testdata/add-directory3.json"},
	}

	runEntries(t, entries)
}

func TestFileCases(t *testing.T) {
	var entries []Entry = []Entry{
		{name: "add a file",
			operations: []Operation{
				{filePath: "hello.txt", operationType: OpAddFile, expectSuccess: true},
			}, resultFile: "testdata/add-file1.json"},

		{name: "error adding a file",
			operations: []Operation{
				{filePath: "hello/world.txt", operationType: OpAddFile, expectSuccess: false},
			}, resultFile: "testdata/new-source-code.json"},

		{name: "add two files",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world.txt", operationType: OpAddFile, expectSuccess: true},
			}, resultFile: "testdata/add-file2.json"},

		{name: "add two files",
			operations: []Operation{
				{filePath: "hello", operationType: OpAddDirectory, expectSuccess: true},
				{filePath: "hello/world.txt", operationType: OpAddFile, expectSuccess: true},
			}, resultFile: "testdata/add-file2.json"},
	}

	runEntries(t, entries)
}

func runEntries(t *testing.T, entries []Entry) {
	for i, e := range entries {
		t.Run(e.name, func(t *testing.T) {
			sc := newSourceCode()
			for j, op := range e.operations {
				var err error
				switch op.operationType {
				case OpAddDirectory:
					err = sc.addDirectory(op.filePath)
				case OpAddFile:
					err = sc.addFile(op.filePath)
				default:
					t.Fatalf("entry %d, op %d faild:\nwrong op.nodeType = %s", i, j, op.operationType)
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
