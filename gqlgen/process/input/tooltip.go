package input

import (
	"fmt"
	"strings"
)

type TooltipTiming = string

const (
	START TooltipTiming = "start"
	END   TooltipTiming = "end"
)

func toTooltipTiming(s string) (TooltipTiming, error) {
	switch strings.ToLower(s) {
	case START:
		return START, nil
	case END:
		return END, nil
	case "":
		return END, nil
	default:
		return "", fmt.Errorf("TooltipTiming value = '%s' is invalid", s)
	}
}

type PositionPreference = string

const (
	ABOVE PositionPreference = "ABOVE"
	BELOW PositionPreference = "BELOW"
	EXACT PositionPreference = "EXACT"
)

func toPositionPreference(s string) (PositionPreference, error) {
	switch strings.ToUpper(s) {
	case ABOVE:
		return ABOVE, nil
	case BELOW:
		return BELOW, nil
	case "":
		return BELOW, nil
	default:
		return "", fmt.Errorf("PositionPreference value = '%s' is invalid", s)
	}
}
