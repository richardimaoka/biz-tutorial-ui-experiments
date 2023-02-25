package model2

import (
	"testing"
)

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
