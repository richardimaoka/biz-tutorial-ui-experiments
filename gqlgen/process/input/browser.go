package input

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var BrowserNumSeqPattern *regexp.Regexp = regexp.MustCompile(`\[[0-9]+\]`)

type BrowserSingle struct {
	StepId        string `json:"stepId"`
	Trivial       bool   `json:"trivial"`
	Comment       string `json:"comment"`
	ImageFileName string `json:"imageFileName"`
}

type BrowserNumSeq struct {
	StepId          string `json:"stepId"`
	Trivial         bool   `json:"trivial"`
	Comment         string `json:"comment"`
	ImageBaseName   string `json:"imageFileBaseName"`
	ImageFileSuffix string `json:"imageFileSuffix"`
	NumImages       int    `json:"numImages"`
}

type BrowserSequence struct {
	StepId         string   `json:"stepId"`
	Trivial        bool     `json:"trivial"`
	Comment        string   `json:"comment"`
	ImageFileNames []string `json:"imageFileNames"`
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

	//
	// Check trivial field
	//
	trivial, err := strToBool(ab.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &BrowserSingle{
		StepId:        ab.StepId,
		Trivial:       trivial,
		Comment:       ab.Comment,
		ImageFileName: imageFileName,
	}, nil
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

	//
	// Check trivial field
	//
	trivial, err := strToBool(ab.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &BrowserNumSeq{
		StepId:          ab.StepId,
		Trivial:         trivial,
		Comment:         ab.Comment,
		ImageBaseName:   baseName,
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

	//
	// Check trivial field
	//
	trivial, err := strToBool(ab.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &BrowserSequence{
		StepId:         ab.StepId,
		Trivial:        trivial,
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
// s to be form of (e.g.) 'filename[42]'
func positiveNumInSqBracket(s string) (int, error) {
	found := BrowserNumSeqPattern.FindString(s)
	if found == "" {
		return 0, fmt.Errorf("the string doesn't have the form '[${num}]'")
	}

	// found = (e.g.) '[42]'
	num, err := strconv.Atoi(found[1 : len(found)-1]) // remove '[' and ']' from 'found'
	if err != nil {
		return 0, err
	}
	if num < 1 {
		return 0, fmt.Errorf("number in '[${num}]' is a negative number = %d", num)
	}

	return num, nil
}
