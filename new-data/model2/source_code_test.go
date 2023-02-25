package model2

import (
	"fmt"
	"strings"
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

func filePathPtrSlice2(filePath string) []*string {
	split := strings.Split(filePath, "/")
	fmt.Printf("filePathPtrSlice2 split = %v\n", split)

	var filePathSlice []*string
	for i, _ := range split {
		filePathSlice = append(filePathSlice, &split[i])
		fmt.Printf("filePathSlice = %v\n", filePathSlice)
	}

	return filePathSlice
}

func TestAddDirectory2(t *testing.T) {
	filePathPtrSlice2("hello/world/goodmorning")

	sc := newSourceCode()
	if err := sc.addDirectory("hello"); err != nil {
		t.Error(err)
		return
	}
	if err := sc.addDirectory("hello/world"); err != nil {
		t.Error(err)
		return
	}

	// json should be same as initial state
	compareAfterMarshal(t, "testdata/add-directory2.json", sc)
}

func TestAddDirectory3(t *testing.T) {
	filePathPtrSlice2("hello/world/goodmorning")

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

	// json should be same as initial state
	compareAfterMarshal(t, "testdata/add-directory3.json", sc)
}
