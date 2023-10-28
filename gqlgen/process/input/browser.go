package input

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

	baseName, err := fileBaseName(ab.Instruction)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction' is invalid, %s", errorPrefix, err)
	}

	suffix, err := fileSuffix(ab.Instruction, ab)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	imageFileName := fmt.Sprintf("%s.%s", baseName, suffix)

	return &BrowserSingle{
		StepId:        ab.StepId,
		Comment:       ab.Comment,
		ImageFileName: imageFileName,
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

	// extract num from (e.g.) filename[30].png
	num, err := positiveNumInSqBracket(ab.Instruction)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction' is in wrong form, %s", errorPrefix, err)
	}

	baseName, err := fileBaseName(ab.Instruction)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction' is invalid, %s", errorPrefix, err)
	}
	baseName = removeSqBrackets(baseName)

	suffix, err := fileSuffix(ab.Instruction, ab)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	return &BrowserNumSeq{
		StepId:          ab.StepId,
		Comment:         ab.Comment,
		ImageBaseName:   baseName,
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

	splitByComma := strings.Split(ab.Instruction, ",")
	var imgFiles []string
	for _, v := range splitByComma {
		baseName, err := fileBaseName(v)
		if err != nil {
			return nil, fmt.Errorf("%s, 'instruction' is invalid, %s", errorPrefix, err)
		}
		suffix, err := fileSuffix(v, ab)
		if err != nil {
			return nil, fmt.Errorf("%s, %s", errorPrefix, err)
		}
		imgFiles = append(imgFiles, fmt.Sprintf("%s.%s", baseName, suffix))
	}

	return &BrowserSequence{
		StepId:         ab.StepId,
		Comment:        ab.Comment,
		ImageFileNames: imgFiles,
	}, nil
}

func removeSqBrackets(s string) string {
	split := strings.Split(s, "[")
	return split[0]
}

// get file base name from given file-name candidate
func fileBaseName(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("file name is empty")
	}

	splitByDot := strings.Split(s, ".")
	if len(splitByDot) == 1 {
		return s, nil
	} else if len(splitByDot) == 2 {
		return splitByDot[0], nil
	} else {
		return "", fmt.Errorf("file name has too many dots")
	}
}

// get file suffix from 1) given file-name candidate, 2) if not there, try finding from Abstract
func fileSuffix(fileNameCandidate string, ab *Abstract) (string, error) {
	splitByDot := strings.Split(fileNameCandidate, ".")
	if len(splitByDot) == 1 {
		// if 'instruction' doesn't have a '.', then the suffix must be in 'instruction2'
		suffix, err := toImageFileSuffix(ab.Instruction2)
		if err != nil {
			return "", fmt.Errorf("file name = '%s' has a wrong suffix, %s", fileNameCandidate, err)
		}
		return suffix, nil
	} else if len(splitByDot) == 2 {
		// if 'instruction' has a '.', then the suffix follows that .
		suffix, err := toImageFileSuffix(splitByDot[1])
		if err != nil {
			return "", fmt.Errorf("'instruction2' has a wrong suffix, %s", err)
		}
		return suffix, nil
	} else {
		return "", fmt.Errorf("file name = '%s' has too many dots", fileNameCandidate)
	}
}

// extract number from square bracket
func positiveNumInSqBracket(s string) (int, error) {
	found := BrowserNumSeqPattern.FindString(s)
	if found == "" {
		return 0, fmt.Errorf("the string doesn't have the form '[${num}]'")
	}

	// found = (e.g.) [30]
	num, err := strconv.Atoi(found[1 : len(found)-1]) // remove '[' and ']' from 'found'
	if err != nil {
		return 0, err
	}
	if num < 1 {
		return 0, fmt.Errorf("number in '[${num}]' is a negative number = %d", num)
	}

	return num, nil
}
