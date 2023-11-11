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

type Column2 interface {
	IsColumn2()
	GetPlaceholder() *string
}

type TerminalElement interface {
	IsTerminalElement()
}

type BackgroundImageColumn struct {
	Placeholder *string `json:"_placeholder"`
	Width       *int    `json:"width"`
	Height      *int    `json:"height"`
	Path        *string `json:"path"`
	URL         *string `json:"url"`
	Modal       *Modal  `json:"modal"`
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

func (BrowserColumn2) IsColumn2()                   {}
func (this BrowserColumn2) GetPlaceholder() *string { return this.Placeholder }

type ColumnWrapper struct {
	Index  *int    `json:"index"`
	Column Column  `json:"column"`
	Name   *string `json:"name"`
}

type ColumnWrapper2 struct {
	Column            Column2 `json:"column"`
	ColumnName        string  `json:"columnName"`
	ColumnDisplayName *string `json:"columnDisplayName"`
}

type DevToolsColumn struct {
	Placeholder *string `json:"_placeholder"`
	Width       *int    `json:"width"`
	Height      *int    `json:"height"`
	Path        *string `json:"path"`
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

type ImageCentered struct {
	Width  *int    `json:"width"`
	Height *int    `json:"height"`
	Path   *string `json:"path"`
	URL    *string `json:"url"`
}

type ImageDescriptionColumn struct {
	Placeholder      *string                 `json:"_placeholder"`
	Description      *Markdown               `json:"description"`
	Image            *ImageCentered          `json:"image"`
	Order            *ImageDescriptionOrder  `json:"order"`
	ContentsPosition *ColumnVerticalPosition `json:"contentsPosition"`
}

type Markdown struct {
	Step      *string            `json:"step"`
	Contents  *string            `json:"contents"`
	Alignment *MarkdownAlignment `json:"alignment"`
}

type MarkdownColumn struct {
	Placeholder      *string                 `json:"_placeholder"`
	Description      *Markdown               `json:"description"`
	ContentsPosition *ColumnVerticalPosition `json:"contentsPosition"`
}

type MarkdownOld struct {
	Step     *string `json:"step"`
	Contents *string `json:"contents"`
}

type Modal struct {
	Text     *string        `json:"text"`
	Position *ModalPosition `json:"position"`
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
	Step            *string          `json:"step"`
	NextStep        *string          `json:"nextStep"`
	PrevStep        *string          `json:"prevStep"`
	AutoNextSeconds *int             `json:"autoNextSeconds"`
	DurationSeconds *int             `json:"durationSeconds"`
	IsTrivialStep   *bool            `json:"isTrivialStep"`
	Columns         []*ColumnWrapper `json:"columns"`
	FocusColumn     *string          `json:"focusColumn"`
	Modal           *Modal           `json:"modal"`
}

type Page2 struct {
	Step        *string           `json:"step"`
	NextStep    *string           `json:"nextStep"`
	PrevStep    *string           `json:"prevStep"`
	IsTrivial   *bool             `json:"isTrivial"`
	Columns     []*ColumnWrapper2 `json:"columns"`
	FocusColumn *string           `json:"focusColumn"`
	Modal       *Modal            `json:"modal"`
}

type SourceCode2 struct {
	Step           *string     `json:"step"`
	ProjectDir     *string     `json:"projectDir"`
	FileTree       []*FileNode `json:"fileTree"`
	IsFoldFileTree *bool       `json:"isFoldFileTree"`
	OpenFile       *OpenFile   `json:"openFile"`
}

type SourceCodeColumn struct {
	Placeholder *string     `json:"_placeholder"`
	SourceCode  *SourceCode `json:"sourceCode"`
}

func (SourceCodeColumn) IsColumn()                    {}
func (this SourceCodeColumn) GetPlaceholder() *string { return this.Placeholder }

type SourceCodeColumn2 struct {
	Placeholder *string      `json:"_placeholder"`
	SourceCode  *SourceCode2 `json:"sourceCode"`
}

func (SourceCodeColumn2) IsColumn2()                   {}
func (this SourceCodeColumn2) GetPlaceholder() *string { return this.Placeholder }

type SourceCodeTooltip struct {
	MarkdownBody string                   `json:"markdownBody"`
	LineNumber   int                      `json:"lineNumber"`
	Timing       *SourceCodeTooltipTiming `json:"timing"`
}

type Terminal struct {
	Step             *string         `json:"step"`
	Name             *string         `json:"name"`
	CurrentDirectory *string         `json:"currentDirectory"`
	Nodes            []*TerminalNode `json:"nodes"`
}

type Terminal2 struct {
	Name             *string          `json:"name"`
	CurrentDirectory string           `json:"currentDirectory"`
	Entries          []*TerminalEntry `json:"entries"`
	Tooltip          *TerminalTooltip `json:"tooltip"`
}

type TerminalColumn struct {
	Placeholder *string   `json:"_placeholder"`
	Terminal    *Terminal `json:"terminal"`
}

func (TerminalColumn) IsColumn()                    {}
func (this TerminalColumn) GetPlaceholder() *string { return this.Placeholder }

type TerminalColumn2 struct {
	Placeholder *string      `json:"_placeholder"`
	Terminals   []*Terminal2 `json:"terminals"`
}

func (TerminalColumn2) IsColumn2()                   {}
func (this TerminalColumn2) GetPlaceholder() *string { return this.Placeholder }

type TerminalCommand struct {
	BeforeExecution *bool   `json:"beforeExecution"`
	Command         *string `json:"command"`
	Tooltip         *string `json:"tooltip"`
}

func (TerminalCommand) IsTerminalElement() {}

type TerminalEntry struct {
	ID        string            `json:"id"`
	EntryType TerminalEntryType `json:"entryType"`
	Text      string            `json:"text"`
}

type TerminalNode struct {
	Content TerminalElement `json:"content"`
}

type TerminalOutput struct {
	Output  *string `json:"output"`
	Tooltip *string `json:"tooltip"`
}

func (TerminalOutput) IsTerminalElement() {}

type TerminalTooltip struct {
	MarkdownBody string                 `json:"markdownBody"`
	Timing       *TerminalTooltipTiming `json:"timing"`
}

type TestObjs struct {
	AppTestTerminalPage              *TerminalColumn2 `json:"appTestTerminalPage"`
	AppTestTutorialColumnsPage       *Page2           `json:"appTestTutorialColumnsPage"`
	AppTestTutorialTutorialPage      *Page2           `json:"appTestTutorialTutorialPage"`
	AppTestSourcecodeFilecontentPage *OpenFile        `json:"appTestSourcecodeFilecontentPage"`
}

type YouTubeColumn struct {
	Placeholder *string       `json:"_placeholder"`
	Youtube     *YouTubeEmbed `json:"youtube"`
}

type YouTubeEmbed struct {
	Width    *int    `json:"width"`
	Height   *int    `json:"height"`
	EmbedURL *string `json:"embedUrl"`
}

type ColumnVerticalPosition string

const (
	ColumnVerticalPositionTop    ColumnVerticalPosition = "TOP"
	ColumnVerticalPositionCenter ColumnVerticalPosition = "CENTER"
	ColumnVerticalPositionBottom ColumnVerticalPosition = "BOTTOM"
)

var AllColumnVerticalPosition = []ColumnVerticalPosition{
	ColumnVerticalPositionTop,
	ColumnVerticalPositionCenter,
	ColumnVerticalPositionBottom,
}

func (e ColumnVerticalPosition) IsValid() bool {
	switch e {
	case ColumnVerticalPositionTop, ColumnVerticalPositionCenter, ColumnVerticalPositionBottom:
		return true
	}
	return false
}

func (e ColumnVerticalPosition) String() string {
	return string(e)
}

func (e *ColumnVerticalPosition) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ColumnVerticalPosition(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ColumnVerticalPosition", str)
	}
	return nil
}

func (e ColumnVerticalPosition) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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

type ImageDescriptionOrder string

const (
	ImageDescriptionOrderImageThenDescription ImageDescriptionOrder = "IMAGE_THEN_DESCRIPTION"
	ImageDescriptionOrderDescriptionThenImage ImageDescriptionOrder = "DESCRIPTION_THEN_IMAGE"
)

var AllImageDescriptionOrder = []ImageDescriptionOrder{
	ImageDescriptionOrderImageThenDescription,
	ImageDescriptionOrderDescriptionThenImage,
}

func (e ImageDescriptionOrder) IsValid() bool {
	switch e {
	case ImageDescriptionOrderImageThenDescription, ImageDescriptionOrderDescriptionThenImage:
		return true
	}
	return false
}

func (e ImageDescriptionOrder) String() string {
	return string(e)
}

func (e *ImageDescriptionOrder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ImageDescriptionOrder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ImageDescriptionOrder", str)
	}
	return nil
}

func (e ImageDescriptionOrder) MarshalGQL(w io.Writer) {
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
