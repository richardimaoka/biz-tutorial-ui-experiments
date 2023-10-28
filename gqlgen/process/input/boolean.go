package input

import (
	"fmt"
	"strings"
)

func strToBool(s string) (bool, error) {
	switch strings.ToUpper(s) {
	case "TRUE":
		return true, nil
	case "FALSE":
		return false, nil
	case "":
		return false, nil
	default:
		return false, fmt.Errorf("'%s' cannnot be converted to bool", s)
	}
}
