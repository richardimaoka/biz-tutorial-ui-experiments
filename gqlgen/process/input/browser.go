package input

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type BrowserNumSeq struct {
	StepId          string
	Comment         string
	ImageBaseName   string
	ImageFileSuffix string
	NumImages       int
}

type BrowserSeq struct {
	StepId         string
	Comment        string
	ImageFileNames []string
}

func toBrowserNumSeq(ab *Abstract) (*BrowserNumSeq, error) {
	if ab.Instruction == "" {
		return nil, fmt.Errorf("failed to convert to BrowserNumSeq, 'instruction' was empty")
	}

	num, err := strconv.Atoi(ab.Instruction2)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to BrowserNumSeq, %s", err)
	}

	return &BrowserNumSeq{
		StepId:          ab.StepId,
		Comment:         ab.Comment,
		ImageBaseName:   ab.Instruction,
		ImageFileSuffix: "png",
		NumImages:       num,
	}, nil
}

func toBrowserSeq(ab *Abstract) (*BrowserSeq, error) {
	if ab.Instruction2 == "" {
		return nil, fmt.Errorf("failed to convert to BrowserSeq, 'instruction2' is empty")
	}
	imgFileSuffix := ab.Instruction2

	fileBaseNames := strings.Split(ab.Instruction, "\n")
	var imgFiles []string
	for i, v := range fileBaseNames {
		if v == "" {
			return nil, fmt.Errorf("failed to convert to BrowserSeq, empty string at %s of 'instruction'", internal.Ordinal(i))
		}
		imgFiles = append(imgFiles, fmt.Sprintf("%s.%s", v, imgFileSuffix))
	}

	return &BrowserSeq{
		StepId:         ab.StepId,
		Comment:        ab.Comment,
		ImageFileNames: imgFiles,
	}, nil
}
