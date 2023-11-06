package column

type ColumnType = string

const (
	// Lower cases since they are from manual entries
	Source   ColumnType = "source"
	Terminal ColumnType = "terminal"
	Browser  ColumnType = "browser"
)
