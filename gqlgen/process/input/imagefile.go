package input

import (
	"fmt"
	"strings"
)

type ImageFileSuffix = string

const (
	JPG  ImageFileSuffix = "jpg"
	JPEG ImageFileSuffix = "jpeg"
	GIF  ImageFileSuffix = "gif"
	PNG  ImageFileSuffix = "png"
)

func toImageFileSuffix(s string) (ImageFileSuffix, error) {
	switch strings.ToLower(s) {
	case JPG:
		return JPG, nil
	case JPEG:
		return JPEG, nil
	case GIF:
		return GIF, nil
	case PNG:
		return PNG, nil
	case "":
		return PNG, nil
	default:
		return "", fmt.Errorf("ImageFileSuffix value = '%s' is invalid", s)
	}
}
