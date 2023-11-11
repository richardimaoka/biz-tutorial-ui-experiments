package state2

type ColumnType string

const (
	NoColumnType       ColumnType = ""
	SourceColumnType   ColumnType = "SourceCode"
	TerminalColumnType ColumnType = "Terminal"
	BrowserColumnType  ColumnType = "Browser"
)

type ColumnFields struct {
	FocusColumn ColumnType `json:"focusColumn"`
	Column1     ColumnType `json:"column1"`
	Column2     ColumnType `json:"column2"`
	Column3     ColumnType `json:"column3"`
	Column4     ColumnType `json:"column4"`
	Column5     ColumnType `json:"column5"`
	Column6     ColumnType `json:"column6"`
	Column7     ColumnType `json:"column7"`
	Column8     ColumnType `json:"column8"`
	Column9     ColumnType `json:"column9"`
	Column10    ColumnType `json:"column10"`
}
