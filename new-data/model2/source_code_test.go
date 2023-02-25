package model2

import (
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
		operations []Operation
		resultFile string
	}

	var entries []Entry = []Entry{
		{operations: []Operation{}, resultFile: "testdata/new-source-code.json"},
	}

	for _, e := range entries {
		sc := newSourceCode()
		for _, op := range e.operations {
			switch op.nodeType {
			case FileNodeTypeDirectory:
				err := sc.addDirectory(op.filePath)
				resultSuccess := err == nil
				if resultSuccess != op.expectSuccess {
					t.Errorf("operation %s is expected, but result is %s", statusString(op.expectSuccess), statusString(resultSuccess))
					continue
				}
			case FileNodeTypeFile:
				err := sc.addFile(op.filePath)
				resultSuccess := err == nil
				if resultSuccess != op.expectSuccess {
					t.Errorf("operation %s is expected, but result is %s", statusString(op.expectSuccess), statusString(resultSuccess))
					continue
				}
			default:
				t.Errorf("wrong op.nodeType = %s", op.nodeType)
				continue
			}
		}
		compareAfterMarshal(t, e.resultFile, sc)
	}
}

func TestNewSourceCode(t *testing.T) {
	sc := newSourceCode()
	compareAfterMarshal(t, "testdata/new-source-code.json", sc)
}

func TestAddDirectoryFailed(t *testing.T) {
	sc := newSourceCode()
	if err := sc.addDirectory(""); err == nil {
		t.Error("error expected")
		return
	}

	if err := sc.addDirectory("abc/"); err == nil {
		t.Error("error expected")
		return
	}
	// json should be same as initial state
	compareAfterMarshal(t, "testdata/new-source-code.json", sc)
}

func TestAddDirectory1(t *testing.T) {
	sc := newSourceCode()
	if err := sc.addDirectory("hello"); err != nil {
		t.Error(err)
		return
	}
	compareAfterMarshal(t, "testdata/add-directory1.json", sc)
}

func TestAddDirectory2(t *testing.T) {
	sc := newSourceCode()
	if err := sc.addDirectory("hello"); err != nil {
		t.Error(err)
		return
	}
	if err := sc.addDirectory("hello/world"); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/add-directory2.json", sc)
}

func TestAddDirectory3(t *testing.T) {
	sc := newSourceCode()
	if err := sc.addDirectory("hello"); err != nil {
		t.Error(err)
		return
	}
	if err := sc.addDirectory("hello/world"); err != nil {
		t.Error(err)
		return
	}
	if err := sc.addDirectory("aloha"); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/add-directory3.json", sc)
}

func TestAddFile1(t *testing.T) {
	sc := newSourceCode()
	if err := sc.addFile("hello.txt"); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/add-file1.json", sc)
}

func TestAddFileFailure(t *testing.T) {
	sc := newSourceCode()
	if err := sc.addFile("hello/world.txt"); err == nil {
		t.Error("error expected")
		return
	}

	// json should be same as initial state
	compareAfterMarshal(t, "testdata/new-source-code.json", sc)
}

func TestAddFile2(t *testing.T) {
	sc := newSourceCode()
	if err := sc.addDirectory("hello"); err != nil {
		t.Error(err)
		return
	}
	if err := sc.addFile("hello/world.txt"); err != nil {
		t.Error(err)
		return
	}

	compareAfterMarshal(t, "testdata/add-file2.json", sc)
}
