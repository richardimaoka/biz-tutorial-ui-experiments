package model2

import (
	"testing"
)

func TestActionSourceCodeAddFile1(t *testing.T) {
	add := FileAdd{FilePath: "hello.txt"}
	sourceCode := NewSourceCode()
	sourceCode.AddFileNode(add.FilePath)
	sourceCode.AddFileContent(add.FilePath, add.Content, add.IsFullContent)
}

//TODO: directry translate AddFile, UpdateFile, ... etc to SourceCode by invoking SourceCode's methods
