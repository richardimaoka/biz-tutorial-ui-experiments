package input

import (
	"fmt"
	"strings"
)

type ColumnType string

const (
	// Lower cases since they are from manual entries
	SourceColumn   ColumnType = "source"
	TerminalColumn ColumnType = "terminal"
	BrowserColumn  ColumnType = "browser"
)

func toColumnType(s string) (ColumnType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(SourceColumn):
		return SourceColumn, nil
	case string(TerminalColumn):
		return TerminalColumn, nil
	case string(BrowserColumn):
		return BrowserColumn, nil
	default:
		return "", fmt.Errorf("'%s' is an invalid column type", s)
	}
}
