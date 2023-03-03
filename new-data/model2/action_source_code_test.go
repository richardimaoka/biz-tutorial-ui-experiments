package model2

import (
	"testing"
)

func TestActionSourceCodeAddFile1(t *testing.T) {
	add := FileAdd{FilePath: "hello.txt"}
	sourceCode := NewSourceCode()
	if err := sourceCode.AddFileNode(add.FilePath); err != nil {
		t.Fatal(err)
	}
	if err := sourceCode.AddFileContent(add.FilePath, add.Content, add.IsFullContent); err != nil {
		t.Fatal(err)
	}
	compareAfterMarshal(t, "testdata/action/command/file-add1.json", sourceCode)
}

//TODO: directry translate AddFile, UpdateFile, ... etc to SourceCode by invoking SourceCode's methods
