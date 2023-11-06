package input

import (
	"fmt"
	"strings"
)

type TooltipTiming = string

const (
	START TooltipTiming = "START"
	END   TooltipTiming = "END"
)

func toTerminalTooltipTiming(s string) (TooltipTiming, error) {
	switch strings.ToUpper(s) {
	case START:
		return START, nil
	case END:
		return END, nil
	case "": // default value is different from source tooltip
		return START, nil
	default:
		return "", fmt.Errorf("TooltipTiming value = '%s' is invalid", s)
	}
}

func toSourceTooltipTiming(s string) (TooltipTiming, error) {
	switch strings.ToUpper(s) {
	case START:
		return START, nil
	case END:
		return END, nil
	case "": // default value is different from termianl tooltip
		return END, nil
	default:
		return "", fmt.Errorf("TooltipTiming value = '%s' is invalid", s)
	}
}

// type PositionPreference = string

// const (
// 	ABOVE PositionPreference = "ABOVE"
// 	BELOW PositionPreference = "BELOW"
// 	EXACT PositionPreference = "EXACT"
// )

// func toPositionPreference(s string) (PositionPreference, error) {
// 	switch strings.ToUpper(s) {
// 	case ABOVE:
// 		return ABOVE, nil
// 	case BELOW:
// 		return BELOW, nil
// 	case "":
// 		return BELOW, nil
// 	default:
// 		return "", fmt.Errorf("PositionPreference value = '%s' is invalid", s)
// 	}
// }

func toTerminalTooltip(fromRow *Row) (*TerminalTooltipRow, error) {
	// if tooltip is empty, then return no tooltip
	if fromRow.Tooltip == "" {
		return nil, nil
	}

	contents := fromRow.Tooltip

	tooltipTiming, err := toTerminalTooltipTiming(fromRow.TooltipTiming)
	if err != nil {
		return nil, fmt.Errorf("'tooltipTiming' field is wrong, %s", err)
	}

	return &TerminalTooltipRow{
		Contents: contents,
		Timing:   tooltipTiming,
	}, nil
}

func toSourceTooltip(fromRow *Row) (*SourceTooltipRow, error) {
	if fromRow.Tooltip == "" {
		return nil, nil
	}

	contents := fromRow.Tooltip

	tooltipTiming, err := toSourceTooltipTiming(fromRow.TooltipTiming)
	if err != nil {
		return nil, fmt.Errorf("'tooltipTiming' field is wrong, %s", err)
	}

	if fromRow.TooltipLine == 0 {
		return nil, fmt.Errorf("'tooltipLine' = %d, cannot be 0 nor empty", fromRow.TooltipLine)
	} else if fromRow.TooltipLine < 0 {
		return nil, fmt.Errorf("'tooltipLine' = %d, but cannot be a negative number", fromRow.TooltipLine)
	}

	return &SourceTooltipRow{
		Contents:   contents,
		Timing:     tooltipTiming,
		LineNumber: fromRow.TooltipLine,
	}, nil
}
