package input

type TooltipTiming = string

const (
	START TooltipTiming = "START"
	END   TooltipTiming = "END"
)

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
