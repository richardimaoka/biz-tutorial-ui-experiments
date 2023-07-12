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
	FilePath  *string       `json:"filePath"`
	Offset    *int          `json:"offset"`
	IsUpdated *bool         `json:"isUpdated"`
}

type Markdown struct {
	Step      *string            `json:"step"`
	Contents  *string            `json:"contents"`
	Alignment *MarkdownAlignment `json:"alignment"`
}

type MarkdownOld struct {
	Step     *string `json:"step"`
	Contents *string `json:"contents"`
}

type NextAction struct {
	TerminalName    *string          `json:"terminalName"`
	TerminalCommand *TerminalCommand `json:"terminalCommand"`
	Markdown        *MarkdownOld     `json:"markdown"`
}

type OpenFile struct {
	FilePath      *string          `json:"filePath"`
	FileName      *string          `json:"fileName"`
	Content       *string          `json:"content"`
	IsFullContent *bool            `json:"isFullContent"`
	Language      *string          `json:"language"`
	Highlight     []*FileHighlight `json:"highlight"`
}

type PageState struct {
	Step       *string      `json:"step"`
	NextStep   *string      `json:"nextStep"`
	PrevStep   *string      `json:"prevStep"`
	SourceCode *SourceCode  `json:"sourceCode"`
	Terminals  []*Terminal  `json:"terminals"`
	Markdown   *MarkdownOld `json:"markdown"`
	NextAction *NextAction  `json:"nextAction"`
}

type Terminal struct {
	Step             *string         `json:"step"`
	Name             *string         `json:"name"`
	CurrentDirectory *string         `json:"currentDirectory"`
	Nodes            []*TerminalNode `json:"nodes"`
}

type TerminalCommand struct {
	BeforeExecution *bool   `json:"beforeExecution"`
	Command         *string `json:"command"`
}

func (TerminalCommand) IsTerminalElement() {}

type TerminalNode struct {
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

type MarkdownAlignment string

const (
	MarkdownAlignmentLeft   MarkdownAlignment = "LEFT"
	MarkdownAlignmentCenter MarkdownAlignment = "CENTER"
)

var AllMarkdownAlignment = []MarkdownAlignment{
	MarkdownAlignmentLeft,
	MarkdownAlignmentCenter,
}

func (e MarkdownAlignment) IsValid() bool {
	switch e {
	case MarkdownAlignmentLeft, MarkdownAlignmentCenter:
		return true
	}
	return false
}

func (e MarkdownAlignment) String() string {
	return string(e)
}

func (e *MarkdownAlignment) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MarkdownAlignment(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MarkdownAlignment", str)
	}
	return nil
}

func (e MarkdownAlignment) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
