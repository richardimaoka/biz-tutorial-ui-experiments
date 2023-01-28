// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type TerminalElement interface {
	IsTerminalElement()
}

type FileHighlight struct {
	FromLine *int `json:"fromLine"`
	ToLine   *int `json:"toLine"`
}

type FileNode struct {
	NodeType  *FileNodeType `json:"nodeType"`
	Name      *string       `json:"name"`
	FilePath  []*string     `json:"filePath"`
	Offset    *int          `json:"offset"`
	IsUpdated *bool         `json:"isUpdated"`
}

type OpenFile struct {
	FilePath      []*string        `json:"filePath"`
	FileName      *string          `json:"fileName"`
	Content       *string          `json:"content"`
	IsFullContent *bool            `json:"isFullContent"`
	Language      *string          `json:"language"`
	Highlight     []*FileHighlight `json:"highlight"`
}

type SourceCode struct {
	FileTree []*FileNode `json:"fileTree"`
	OpenFile *OpenFile   `json:"openFile"`
}

type Step struct {
	StepNum    *int        `json:"stepNum"`
	SourceCode *SourceCode `json:"sourceCode"`
	Terminals  []*Terminal `json:"terminals"`
	NextAction *string     `json:"nextAction"`
}

type Terminal struct {
	Name             *string         `json:"name"`
	CurrentDirectory []*string       `json:"currentDirectory"`
	Nodes            []*TerminalNode `json:"nodes"`
}

type TerminalCommand struct {
	Command *string `json:"command"`
}

func (TerminalCommand) IsTerminalElement() {}

type TerminalCommandSet struct {
	Commands []*TerminalCommand `json:"commands"`
}

func (TerminalCommandSet) IsTerminalElement() {}

type TerminalNode struct {
	Index   *int            `json:"index"`
	Content TerminalElement `json:"content"`
}

type TerminalOutput struct {
	Output *string `json:"output"`
}

func (TerminalOutput) IsTerminalElement() {}

type FileNodeType string

const (
	FileNodeTypeFile      FileNodeType = "FILE"
	FileNodeTypeDirectory FileNodeType = "DIRECTORY"
)

var AllFileNodeType = []FileNodeType{
	FileNodeTypeFile,
	FileNodeTypeDirectory,
}

func (e FileNodeType) IsValid() bool {
	switch e {
	case FileNodeTypeFile, FileNodeTypeDirectory:
		return true
	}
	return false
}

func (e FileNodeType) String() string {
	return string(e)
}

func (e *FileNodeType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FileNodeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FileNodeType", str)
	}
	return nil
}

func (e FileNodeType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
