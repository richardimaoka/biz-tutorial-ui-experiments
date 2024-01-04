package input

import (
	"fmt"
	"strings"
)

type Mode string

const (
	// Lower cases since they are from manual entries
	SlideshowMode Mode = "slideshow"
	HandsonMode   Mode = "handson"
)

func toMode(s string) (Mode, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(SlideshowMode):
		return SlideshowMode, nil
	case string(HandsonMode):
		return HandsonMode, nil
	default:
		return "", fmt.Errorf("mode  = '%s' is invalid", s)
	}
}
