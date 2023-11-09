package input

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/result"
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

/**
 * BrowserSubType type(s) and functions
 */
type BrowserSubType string

const (
	// Lower cases since they are from manual entries
	BrowserSingle   BrowserSubType = "single"
	BrowserNumSeq   BrowserSubType = "numseq"
	BrowserSequence BrowserSubType = "seq"
)

func toBrowserSubType(s string) (BrowserSubType, error) {
	lower := strings.ToLower(s)

	switch lower {
	case string(BrowserSingle):
		return BrowserSingle, nil
	case string(BrowserNumSeq):
		return BrowserNumSeq, nil
	case string(BrowserSequence):
		return BrowserSequence, nil
	default:
		return "", fmt.Errorf("'%s' is an invalid browser sub type", s)
	}
}

/**
 * Browser row type(s) and functions
 */

type BrowserRow struct {
	StepId         string `json:"stepId"`
	Trivial        bool   `json:"trivial"`
	Comment        string `json:"comment"`
	ImageFileNames string `json:"imageFileNames"`
}

type BrowserSingleRow struct {
	StepId        string `json:"stepId"`
	Trivial       bool   `json:"trivial"`
	Comment       string `json:"comment"`
	ImageFileName string `json:"imageFileName"`
}

var BrowserNumSeqPattern *regexp.Regexp = regexp.MustCompile(`\[[0-9]+\]`)

type BrowserNumSeqRow struct {
	StepId          string `json:"stepId"`
	Trivial         bool   `json:"trivial"`
	Comment         string `json:"comment"`
	ImageBaseName   string `json:"imageFileBaseName"`
	ImageFileSuffix string `json:"imageFileSuffix"`
	NumImages       int    `json:"numImages"`
}

type BrowserSequenceRow struct {
	StepId         string   `json:"stepId"`
	Trivial        bool     `json:"trivial"`
	Comment        string   `json:"comment"`
	ImageFileNames []string `json:"imageFileNames"`
}

func toBrowserSingleRow(fromRow *Row) (*BrowserSingleRow, error) {
	errorPrefix := "failed to convert to BrowserNumSeq"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != BrowserType {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	subType, err := toBrowserSubType(fromRow.Type)
	if err != nil || subType != BrowserSingle {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	baseName, err := fileBaseName(fromRow.Instruction)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction' is invalid, %s", errorPrefix, err)
	}

	suffix, err := fileSuffix(fromRow.Instruction, fromRow)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	imageFileName := fmt.Sprintf("%s.%s", baseName, suffix)

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &BrowserSingleRow{
		StepId:        fromRow.StepId,
		Trivial:       trivial,
		Comment:       fromRow.Comment,
		ImageFileName: imageFileName,
	}, nil
}

func toBrowserNumSeqRow(fromRow *Row) (*BrowserNumSeqRow, error) {
	errorPrefix := "failed to convert to BrowserNumSeq"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != "browser" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != "numseq" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	// extract num from (e.g.) filename[30].png
	num, err := positiveNumInSqBracket(fromRow.Instruction)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction' is in wrong form, %s", errorPrefix, err)
	}

	baseName, err := fileBaseName(fromRow.Instruction)
	if err != nil {
		return nil, fmt.Errorf("%s, 'instruction' is invalid, %s", errorPrefix, err)
	}
	baseName = removeSqBrackets(baseName)

	suffix, err := fileSuffix(fromRow.Instruction, fromRow)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", errorPrefix, err)
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &BrowserNumSeqRow{
		StepId:          fromRow.StepId,
		Trivial:         trivial,
		Comment:         fromRow.Comment,
		ImageBaseName:   baseName,
		ImageFileSuffix: suffix,
		NumImages:       num,
	}, nil
}

func toBrowserSequenceRow(fromRow *Row) (*BrowserSequenceRow, error) {
	errorPrefix := "failed to convert to BrowserSequence"

	//
	// Check column and type
	//
	if strings.ToLower(fromRow.Column) != "browser" {
		return nil, fmt.Errorf("%s, called for wrong 'column' = %s", errorPrefix, fromRow.Column)
	}
	if strings.ToLower(fromRow.Type) != "sequence" {
		return nil, fmt.Errorf("%s, called for wrong 'type' = %s", errorPrefix, fromRow.Type)
	}

	//
	// Check instruction
	//
	if fromRow.Instruction == "" {
		return nil, fmt.Errorf("%s, 'instruction' is empty", errorPrefix)
	}

	splitByComma := strings.Split(fromRow.Instruction, ",")
	var imgFiles []string
	for _, v := range splitByComma {
		baseName, err := fileBaseName(v)
		if err != nil {
			return nil, fmt.Errorf("%s, 'instruction' is invalid, %s", errorPrefix, err)
		}
		suffix, err := fileSuffix(v, fromRow)
		if err != nil {
			return nil, fmt.Errorf("%s, %s", errorPrefix, err)
		}
		imgFiles = append(imgFiles, fmt.Sprintf("%s.%s", baseName, suffix))
	}

	//
	// Check trivial field
	//
	trivial, err := strToBool(fromRow.Trivial)
	if err != nil {
		return nil, fmt.Errorf("%s, 'trivial' is invalid, %s", errorPrefix, err)
	}

	return &BrowserSequenceRow{
		StepId:         fromRow.StepId,
		Trivial:        trivial,
		Comment:        fromRow.Comment,
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

// get file suffix from 1) given file-name candidate, 2) if not there, try finding from Row
func fileSuffix(fileNameCandidate string, fromRow *Row) (string, error) {
	splitByDot := strings.Split(fileNameCandidate, ".")
	if len(splitByDot) == 1 {
		// if 'instruction' doesn't have a '.', then the suffix must be in 'instruction2'
		suffix, err := toImageFileSuffix(fromRow.Instruction2)
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

/**
 * Function to turn a row into steps
 */

func toBrowserSteps(
	r *Row,
	finder *StepIdFinder,
	prevColumns ColumnInfo,
) ([]result.Step, ColumnInfo, error) {
	return nil, ColumnInfo{}, nil
}
