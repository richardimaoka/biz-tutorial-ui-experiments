package input

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type ImageFileSuffix = string

const (
	JPG  ImageFileSuffix = "jpg"
	JPEG ImageFileSuffix = "jpeg"

	GIF ImageFileSuffix = "gif"

	PNG ImageFileSuffix = "png"
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

type BrowserNumSeq struct {
	StepId          string
	Comment         string
	ImageBaseName   string
	ImageFileSuffix string
	NumImages       int
}

type BrowserSequence struct {
	StepId         string
	Comment        string
	ImageFileNames []string
}

func toBrowserNumSeq(ab *Abstract) (*BrowserNumSeq, error) {
	errorPrefix := "failed to convert to BrowserNumSeq"

	//
	// Check column and type
	//
	if strings.ToLower(ab.Column) != "browser" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, ab.Column)
	}
	if strings.ToLower(ab.Type) != "numseq" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, ab.Type)
	}

	//
	// Check instruction fields
	//
	if ab.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' was empty", errorPrefix)
	}

	suffix, err := toImageFileSuffix(ab.Instruction2)
	if err != nil {
		return nil, fmt.Errorf("%s,'instruction2' is wrong, %s", errorPrefix, err)
	}

	num, err := strconv.Atoi(ab.Instruction3)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	return &BrowserNumSeq{
		StepId:          ab.StepId,
		Comment:         ab.Comment,
		ImageBaseName:   ab.Instruction,
		ImageFileSuffix: suffix,
		NumImages:       num,
	}, nil
}

func toBrowserSequence(ab *Abstract) (*BrowserSequence, error) {
	errorPrefix := "failed to convert to BrowserSequence"

	//
	// Check column and type
	//
	if strings.ToLower(ab.Column) != "browser" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, ab.Column)
	}
	if strings.ToLower(ab.Type) != "seq" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, ab.Type)
	}

	//
	// Check instruction fields
	//
	if ab.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' was empty", errorPrefix)
	}

	if ab.Instruction2 == "" {
		return nil, fmt.Errorf("%s, 'instruction2' is empty", errorPrefix)
	}
	imgFileSuffix := ab.Instruction2

	fileBaseNames := strings.Split(ab.Instruction, "\n")
	var imgFiles []string
	for i, v := range fileBaseNames {
		if v == "" {
			return nil, fmt.Errorf("%s, empty string at %s of 'instruction'", errorPrefix, internal.Ordinal(i))
		}
		imgFiles = append(imgFiles, fmt.Sprintf("%s.%s", v, imgFileSuffix))
	}

	return &BrowserSequence{
		StepId:         ab.StepId,
		Comment:        ab.Comment,
		ImageFileNames: imgFiles,
	}, nil
}
