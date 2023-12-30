package input

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/csvfield"

type Row struct {
	RowId   string        `json:"rowId"`
	Phase   string        `json:"phase"`
	Comment string        `json:"comment"`
	Trivial csvfield.Bool `json:"trivial"`

	// Mode and type fileds
	Mode    string `json:"mode"`
	RowType string `json:"rowType"`
	SubType string `json:"subtype"`

	// Modal fileds
	ModalContents string `json:"modalContents"`
	ModalPosition string `json:"modalPosition"`

	// Contents fileds
	Contents        string        `json:"contents"`
	TypingAnimation csvfield.Bool `json:"typingAnimation"`
	TerminalName    string        `json:"terminalName"`

	// FilePath fileds
	FilePath     string            `json:"filePath"`
	ImageSize    string            `json:"imageSize"`
	ImageWidth   csvfield.MultiInt `json:"imageWidth"`
	ImageHeight  csvfield.MultiInt `json:"imageHeight"`
	ImageCaption string            `json:"imageCaption"`

	// Tooltip fileds
	Tooltip       string      `json:"tooltip"`
	TooltipTiming string      `json:"tooltipTiming"`
	TooltipLine   IntOrString `json:"tooltipLine"`   // IntOrString because an empty cell from CSV becomes empty string = ""
	TooltipAppend string      `json:"tooltipAppend"` // string (i.e. 'TRUE', 'FALSE', or empty string), not bool, as it is a value from CSV
}
