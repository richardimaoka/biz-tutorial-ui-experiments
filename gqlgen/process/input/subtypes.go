package input

import (
	"fmt"
	"strings"
)

type SubType string

const (
	// Lower cases since they are from manual entries
	CommandSubType SubType = "command"
	OutputSubType  SubType = "output"
)

func toCommandSubType(s string) (SubType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(CommandSubType):
		return CommandSubType, nil
	case string(OutputSubType):
		return OutputSubType, nil
	default:
		return "", fmt.Errorf("'%s' is an invalid sub type", s)
	}
}
