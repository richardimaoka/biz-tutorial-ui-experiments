package input

import (
	"fmt"
	"regexp"
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

type BrowserSingle struct {
	StepId        string
	Comment       string
	ImageFileName string
}

func toBrowserSingle(ab *Abstract) (*BrowserSingle, error) {
	errorPrefix := "failed to convert to BrowserNumSeq"

	//
	// Check column and type
	//
	if strings.ToLower(ab.Column) != "browser" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, ab.Column)
	}
	if ab.Type != "" && strings.ToLower(ab.Type) != "single" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, ab.Type)
	}

	//
	// Check instruction fields
	//
	if ab.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	// fileBaseName := ab.Instruction

	// suffix, err := toImageFileSuffix(ab.Instruction2)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s,'instruction2' is wrong, %s", errorPrefix, err)
	// }

	// imageFileNme := fmt.Sprintf("%s.%s", fileBaseName, suffix)

	var imageFileNme string
	split := strings.Split(ab.Instruction, ".")
	if len(split) == 1 {
		fileBaseName := split[0] // confirmed non-empty already

		suffix, err := toImageFileSuffix(ab.Instruction2)
		if err != nil {
			return nil, fmt.Errorf("%s,'instruction2' is wrong, %s", errorPrefix, err)
		}

		imageFileNme = fmt.Sprintf("%s.%s", fileBaseName, suffix)
	} else if len(split) == 2 {
		_, err := toImageFileSuffix(split[1])
		if err != nil {
			return nil, fmt.Errorf("%s,'instruction' has wrong file suffix, %s", errorPrefix, err)
		}
		imageFileNme = ab.Instruction
	} else {
		return nil, fmt.Errorf("%s,'instruction' has too many dots", errorPrefix)
	}

	return &BrowserSingle{
		StepId:        ab.StepId,
		Comment:       ab.Comment,
		ImageFileName: imageFileNme,
	}, nil
}

type BrowserNumSeq struct {
	StepId          string
	Comment         string
	ImageBaseName   string
	ImageFileSuffix string
	NumImages       int
}

var BrowserNumSeqPattern *regexp.Regexp = regexp.MustCompile(`\[[0-9]+\]`)

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
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	found := BrowserNumSeqPattern.FindString(ab.Instruction)
	if found == "" {
		return nil, fmt.Errorf("%s, 'instruction' should be the form of 'filebase[${num}]'", errorPrefix)
	}

	// found = (e.g.) [30]
	num, err := strconv.Atoi(found[1 : len(found)-1]) // remove '[' and ']' from 'found'
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}
	if num < 1 {
		return nil, fmt.Errorf("%s, 'instruction' must have a positive number in 'filebase[${num}]'", errorPrefix)
	}

	var imageBaseName string
	var suffix string
	splitByDot := strings.Split(ab.Instruction, ".")
	if len(splitByDot) == 1 {
		splitBySquareBracket := strings.Split(ab.Instruction, "[")
		imageBaseName = splitBySquareBracket[0] // confirmed non-empty already

		suffix, err = toImageFileSuffix(ab.Instruction2)
		if err != nil {
			return nil, fmt.Errorf("%s,'instruction2' is empty, %s", errorPrefix, err)
		}

	} else if len(splitByDot) == 2 {
		splitBySquareBracket := strings.Split(ab.Instruction, "[")
		imageBaseName = splitBySquareBracket[0] // confirmed non-empty already

		suffix, err = toImageFileSuffix(splitByDot[1])
		if err != nil {
			return nil, fmt.Errorf("%s, 'instruction' has wrong file suffix, %s", errorPrefix, err)
		}
	} else {
		return nil, fmt.Errorf("%s, 'instruction' has too many dots", errorPrefix)
	}

	return &BrowserNumSeq{
		StepId:          ab.StepId,
		Comment:         ab.Comment,
		ImageBaseName:   imageBaseName,
		ImageFileSuffix: suffix,
		NumImages:       num,
	}, nil
}

type BrowserSequence struct {
	StepId         string
	Comment        string
	ImageFileNames []string
}

func toBrowserSequence(ab *Abstract) (*BrowserSequence, error) {
	errorPrefix := "failed to convert to BrowserSequence"

	//
	// Check column and type
	//
	if strings.ToLower(ab.Column) != "browser" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, ab.Column)
	}
	if strings.ToLower(ab.Type) != "sequence" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, ab.Type)
	}

	//
	// Check instruction fields
	//
	if ab.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	fileBaseNames := strings.Split(ab.Instruction, ",")
	var imgFiles []string
	for i, v := range fileBaseNames {
		if v == "" {
			return nil, fmt.Errorf("%s, %s comma-separated element in 'instruction' is empty", errorPrefix, internal.Ordinal(i))
		}

		split := strings.Split(v, ".")
		if len(split) == 1 {
			fileBaseName := split[0] // confirmed non-empty already

			suffix, err := toImageFileSuffix(ab.Instruction2)
			if err != nil {
				return nil, fmt.Errorf("%s,'instruction2' is wrong, %s", errorPrefix, err)
			}

			imgFiles = append(imgFiles, fmt.Sprintf("%s.%s", fileBaseName, suffix))
		} else if len(split) == 2 {
			_, err := toImageFileSuffix(split[1])
			if err != nil {
				return nil, fmt.Errorf("%s,'instruction' has wrong file suffix, %s", errorPrefix, err)
			}
			imgFiles = append(imgFiles, v)
		} else {
			return nil, fmt.Errorf("%s,'instruction' has too many dots", errorPrefix)
		}

	}

	return &BrowserSequence{
		StepId:         ab.StepId,
		Comment:        ab.Comment,
		ImageFileNames: imgFiles,
	}, nil
}
