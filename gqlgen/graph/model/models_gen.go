// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Column interface {
	IsColumn()
	GetPlaceholder() *string
}

type Browser struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Path   string `json:"path"`
}

type BrowserColumn2 struct {
	Placeholder *string  `json:"_placeholder"`
	Browser     *Browser `json:"browser"`
}

func (BrowserColumn2) IsColumn()                    {}
func (this BrowserColumn2) GetPlaceholder() *string { return this.Placeholder }

type ColumnWrapper struct {
	Column            Column  `json:"column"`
	ColumnName        string  `json:"columnName"`
	ColumnDisplayName *string `json:"columnDisplayName"`
}

type EditSequence struct {
	ID    string                 `json:"id"`
	Edits []*MonacoEditOperation `json:"edits"`
}

type FileHighlight struct {
	FromLine *int `json:"fromLine"`
	ToLine   *int `json:"toLine"`
}

type FileNode struct {
	NodeType  FileNodeType `json:"nodeType"`
	Name      *string      `json:"name"`
	FilePath  string       `json:"filePath"`
	Offset    *int         `json:"offset"`
	IsUpdated *bool        `json:"isUpdated"`
	IsDeleted *bool        `json:"isDeleted"`
}

type Modal struct {
	MarkdownBody *string `json:"markdownBody"`
}

type MonacoEditOperation struct {
	Text  string           `json:"text"`
	Range *MonacoEditRange `json:"range"`
}

type MonacoEditRange struct {
	StartLineNumber int `json:"startLineNumber"`
	StartColumn     int `json:"startColumn"`
	EndLineNumber   int `json:"endLineNumber"`
	EndColumn       int `json:"endColumn"`
}

type OpenFile struct {
	FilePath      *string            `json:"filePath"`
	FileName      *string            `json:"fileName"`
	Content       *string            `json:"content"`
	OldContent    *string            `json:"oldContent"`
	IsFullContent *bool              `json:"isFullContent"`
	Language      *string            `json:"language"`
	Size          *float64           `json:"size"`
	EditSequence  *EditSequence      `json:"editSequence"`
	Tooltip       *SourceCodeTooltip `json:"tooltip"`
	Highlight     []*FileHighlight   `json:"highlight"`
}

type Page struct {
	Step        *string          `json:"step"`
	NextStep    *string          `json:"nextStep"`
	PrevStep    *string          `json:"prevStep"`
	IsTrivial   *bool            `json:"isTrivial"`
	Columns     []*ColumnWrapper `json:"columns"`
	FocusColumn *string          `json:"focusColumn"`
	Modal       *Modal           `json:"modal"`
}

type SourceCode2 struct {
	Step           *string     `json:"step"`
	ProjectDir     *string     `json:"projectDir"`
	FileTree       []*FileNode `json:"fileTree"`
	IsFoldFileTree *bool       `json:"isFoldFileTree"`
	OpenFile       *OpenFile   `json:"openFile"`
}

type SourceCodeColumn2 struct {
	Placeholder *string      `json:"_placeholder"`
	SourceCode  *SourceCode2 `json:"sourceCode"`
}

func (SourceCodeColumn2) IsColumn()                    {}
func (this SourceCodeColumn2) GetPlaceholder() *string { return this.Placeholder }

type SourceCodeTooltip struct {
	MarkdownBody string                   `json:"markdownBody"`
	LineNumber   int                      `json:"lineNumber"`
	Timing       *SourceCodeTooltipTiming `json:"timing"`
}

type Terminal struct {
	Name             *string          `json:"name"`
	CurrentDirectory string           `json:"currentDirectory"`
	Entries          []*TerminalEntry `json:"entries"`
	Tooltip          *TerminalTooltip `json:"tooltip"`
}

type TerminalColumn struct {
	Placeholder *string     `json:"_placeholder"`
	Terminals   []*Terminal `json:"terminals"`
}

func (TerminalColumn) IsColumn()                    {}
func (this TerminalColumn) GetPlaceholder() *string { return this.Placeholder }

type TerminalEntry struct {
	ID        string            `json:"id"`
	EntryType TerminalEntryType `json:"entryType"`
	Text      string            `json:"text"`
}

type TerminalTooltip struct {
	MarkdownBody string                 `json:"markdownBody"`
	Timing       *TerminalTooltipTiming `json:"timing"`
}

type TestObjs struct {
	AppTestTerminalPage              *TerminalColumn `json:"appTestTerminalPage"`
	AppTestTutorialColumnsPage       *Page           `json:"appTestTutorialColumnsPage"`
	AppTestTutorialTutorialPage      *Page           `json:"appTestTutorialTutorialPage"`
	AppTestSourcecodeFilecontentPage *OpenFile       `json:"appTestSourcecodeFilecontentPage"`
}

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

type SourceCodeTooltipTiming string

const (
	SourceCodeTooltipTimingStart SourceCodeTooltipTiming = "START"
	SourceCodeTooltipTimingEnd   SourceCodeTooltipTiming = "END"
)

var AllSourceCodeTooltipTiming = []SourceCodeTooltipTiming{
	SourceCodeTooltipTimingStart,
	SourceCodeTooltipTimingEnd,
}

func (e SourceCodeTooltipTiming) IsValid() bool {
	switch e {
	case SourceCodeTooltipTimingStart, SourceCodeTooltipTimingEnd:
		return true
	}
	return false
}

func (e SourceCodeTooltipTiming) String() string {
	return string(e)
}

func (e *SourceCodeTooltipTiming) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SourceCodeTooltipTiming(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SourceCodeTooltipTiming", str)
	}
	return nil
}

func (e SourceCodeTooltipTiming) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TerminalEntryType string

const (
	TerminalEntryTypeCommand TerminalEntryType = "COMMAND"
	TerminalEntryTypeOutput  TerminalEntryType = "OUTPUT"
)

var AllTerminalEntryType = []TerminalEntryType{
	TerminalEntryTypeCommand,
	TerminalEntryTypeOutput,
}

func (e TerminalEntryType) IsValid() bool {
	switch e {
	case TerminalEntryTypeCommand, TerminalEntryTypeOutput:
		return true
	}
	return false
}

func (e TerminalEntryType) String() string {
	return string(e)
}

func (e *TerminalEntryType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TerminalEntryType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TerminalEntryType", str)
	}
	return nil
}

func (e TerminalEntryType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TerminalTooltipTiming string

const (
	TerminalTooltipTimingStart TerminalTooltipTiming = "START"
	TerminalTooltipTimingEnd   TerminalTooltipTiming = "END"
)

var AllTerminalTooltipTiming = []TerminalTooltipTiming{
	TerminalTooltipTimingStart,
	TerminalTooltipTimingEnd,
}

func (e TerminalTooltipTiming) IsValid() bool {
	switch e {
	case TerminalTooltipTimingStart, TerminalTooltipTimingEnd:
		return true
	}
	return false
}

func (e TerminalTooltipTiming) String() string {
	return string(e)
}

func (e *TerminalTooltipTiming) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TerminalTooltipTiming(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TerminalTooltipTiming", str)
	}
	return nil
}

func (e TerminalTooltipTiming) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
