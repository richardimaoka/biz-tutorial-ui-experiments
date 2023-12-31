package input

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/csvfield"

type Row struct {
	RowId   string        `json:"rowId,omitempty,omitempty"`
	Phase   string        `json:"phase,omitempty"`
	Comment string        `json:"comment,omitempty"`
	Trivial csvfield.Bool `json:"trivial,omitempty"`

	// Mode and type fileds
	Mode    string `json:"mode,omitempty"`
	RowType string `json:"rowType,omitempty"`
	SubType string `json:"subtype,omitempty"`

	// Modal fileds
	ModalContents string `json:"modalContents,omitempty"`
	ModalPosition string `json:"modalPosition,omitempty"`

	// Contents fileds
	Contents        string        `json:"contents,omitempty"`
	TypingAnimation csvfield.Bool `json:"typingAnimation,omitempty"`
	TerminalName    string        `json:"terminalName,omitempty"`

	// FilePath fileds
	FilePath     string            `json:"filePath,omitempty"`
	ImageSize    string            `json:"imageSize,omitempty"`
	ImageWidth   csvfield.MultiInt `json:"imageWidth,omitempty"`
	ImageHeight  csvfield.MultiInt `json:"imageHeight,omitempty"`
	ImageCaption string            `json:"imageCaption,omitempty"`

	// Tooltip fileds
	Tooltip       string      `json:"tooltip,omitempty"`
	TooltipTiming string      `json:"tooltipTiming,omitempty"`
	TooltipLine   IntOrString `json:"tooltipLine,omitempty"`   // IntOrString because an empty cell from CSV becomes empty string = ""
	TooltipAppend string      `json:"tooltipAppend,omitempty"` // string (i.e. 'TRUE', 'FALSE', or empty string), not bool, as it is a value from CSV
}
