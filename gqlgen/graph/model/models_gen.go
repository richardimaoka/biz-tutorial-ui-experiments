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

type Slide interface {
	IsSlide()
	GetPlaceholder() *string
}

type Browser struct {
	Image *Image `json:"image"`
}

type BrowserColumn struct {
	Placeholder *string  `json:"_placeholder"`
	Browser     *Browser `json:"browser"`
}

func (BrowserColumn) IsColumn()                    {}
func (this BrowserColumn) GetPlaceholder() *string { return this.Placeholder }

type ColumnWrapper struct {
	ColumnName        string  `json:"columnName"`
	ColumnDisplayName *string `json:"columnDisplayName"`
	Column            Column  `json:"column"`
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

type Image struct {
	Src     string  `json:"src"`
	Width   int     `json:"width"`
	Height  int     `json:"height"`
	Caption *string `json:"caption"`
}

type ImageSlide struct {
	Placeholder *string `json:"_placeholder"`
	Image       *Image  `json:"image"`
}

func (ImageSlide) IsSlide()                     {}
func (this ImageSlide) GetPlaceholder() *string { return this.Placeholder }

type MarkdownSlide struct {
	Placeholder  *string `json:"_placeholder"`
	MarkdownBody string  `json:"markdownBody"`
}

func (MarkdownSlide) IsSlide()                     {}
func (this MarkdownSlide) GetPlaceholder() *string { return this.Placeholder }

type Modal struct {
	MarkdownBody string         `json:"markdownBody"`
	Position     *ModalPosition `json:"position"`
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
	Modal       *Modal           `json:"modal"`
	Mode        PageMode         `json:"mode"`
	Slide       *SlideWrapper    `json:"slide"`
	FocusColumn *string          `json:"focusColumn"`
	Columns     []*ColumnWrapper `json:"columns"`
}

type SectionTitleSlide struct {
	Placeholder *string `json:"_placeholder"`
	SectionNum  int     `json:"sectionNum"`
	Title       string  `json:"title"`
}

func (SectionTitleSlide) IsSlide()                     {}
func (this SectionTitleSlide) GetPlaceholder() *string { return this.Placeholder }

type SlideWrapper struct {
	Slide Slide `json:"slide"`
}

type SourceCodeColumn struct {
	Placeholder *string     `json:"_placeholder"`
	SourceCode  *SourceCode `json:"sourceCode"`
}

func (SourceCodeColumn) IsColumn()                    {}
func (this SourceCodeColumn) GetPlaceholder() *string { return this.Placeholder }

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

type TutorialTitleSlide struct {
	Placeholder *string  `json:"_placeholder"`
	Title       string   `json:"title"`
	Images      []*Image `json:"images"`
}

func (TutorialTitleSlide) IsSlide()                     {}
func (this TutorialTitleSlide) GetPlaceholder() *string { return this.Placeholder }

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

type ModalPosition string

const (
	ModalPositionTop    ModalPosition = "TOP"
	ModalPositionCenter ModalPosition = "CENTER"
	ModalPositionBottom ModalPosition = "BOTTOM"
)

var AllModalPosition = []ModalPosition{
	ModalPositionTop,
	ModalPositionCenter,
	ModalPositionBottom,
}

func (e ModalPosition) IsValid() bool {
	switch e {
	case ModalPositionTop, ModalPositionCenter, ModalPositionBottom:
		return true
	}
	return false
}

func (e ModalPosition) String() string {
	return string(e)
}

func (e *ModalPosition) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ModalPosition(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ModalPosition", str)
	}
	return nil
}

func (e ModalPosition) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PageMode string

const (
	PageModeSlideshow PageMode = "SLIDESHOW"
	PageModeHandson   PageMode = "HANDSON"
)

var AllPageMode = []PageMode{
	PageModeSlideshow,
	PageModeHandson,
}

func (e PageMode) IsValid() bool {
	switch e {
	case PageModeSlideshow, PageModeHandson:
		return true
	}
	return false
}

func (e PageMode) String() string {
	return string(e)
}

func (e *PageMode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PageMode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PageMode", str)
	}
	return nil
}

func (e PageMode) MarshalGQL(w io.Writer) {
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
