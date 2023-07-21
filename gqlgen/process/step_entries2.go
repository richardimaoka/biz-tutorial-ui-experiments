package process

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

type StepEntry2 struct {
	// Uppercase fields to allow json dump for testing

	// steps
	Step              string `json:"step"`
	AutoAnimateSecond int    `json:"autoAnimateSecond,omitempty"`

	// columns
	FocusColumn string `json:"focusColumn,omitempty"`
	Column1     string `json:"column1,omitempty"`
	Column2     string `json:"column2,omitempty"`
	Column3     string `json:"column3,omitempty"`
	Column4     string `json:"column4,omitempty"`
	Column5     string `json:"column5,omitempty"`

	// modal
	ModalText     string `json:"modalText,omitempty"`
	ModalPosition string `json:"modalPosition,omitempty"`

	// terminal
	TerminalText string `json:"terminalText,omitempty"`
	TerminalType string `json:"terminalType,omitempty"`

	// git
	Commit        string `json:"commit,omitempty"`
	CommitMessage string `json:"commitMessage,omitempty"`
	Repo          string `json:"repo,omitempty"`

	// browser
	BrowserType        string `json:"browserType,omitempty"`
	BrowserImagePath   string `json:"browserImagePath,omitempty"`
	BrowserImageWidth  string `json:"browserImageWidth,omitempty"`
	BrowserImageHeight string `json:"browserImageHeight,omitempty"`
}

type StepEntries2 []StepEntry2

func (entries StepEntries2) calcSteps(seqNo int) (string, string, string) {
	var currentStep string
	if seqNo == 0 {
		currentStep = "_initial"
	} else {
		currentStep = entries[seqNo].Step
	}

	var prevStep string
	if seqNo == 0 {
		prevStep = ""
	} else {
		prevStep = entries[seqNo-1].Step
	}

	var nextStep string
	if seqNo == len(entries)-1 {
		nextStep = ""
	} else {
		nextStep = entries[seqNo+1].Step
	}

	return currentStep, prevStep, nextStep
}

func (e StepEntry2) columns(seqNo int) []string {
	var columns []string

	if e.Column1 != "" {
		columns = append(columns, e.Column1)
	}
	if e.Column2 != "" {
		columns = append(columns, e.Column2)
	}
	if e.Column3 != "" {
		columns = append(columns, e.Column3)
	}
	if e.Column4 != "" {
		columns = append(columns, e.Column4)
	}
	if e.Column5 != "" {
		columns = append(columns, e.Column5)
	}

	return columns
}

func (entries StepEntries2) ToGraphQLPages() ([]model.Page, error) {
	// var srcColState *state.SourceCodeColumn
	var terminalColumnState *state.TerminalColumn

	var pages []model.Page
	for seqNo, e := range entries {
		currentStep, prevStep, nextStep := entries.calcSteps(seqNo)
		columns := e.columns(seqNo)

		var colWrappers []*model.ColumnWrapper
		for _, colName := range columns {
			if colName == "terminal" {
				if terminalColumnState == nil {
					terminalColumnState = state.NewTerminalColumn()
				}

				if e.TerminalText != "" {
					terminalType, err := state.ToTerminalElementType(e.TerminalType)
					if err != nil {
						return nil, fmt.Errorf("ToGraphQLPages failed to convert terminal type, %s", err)
					}
					terminalColumnState.Transition(terminalType, e.TerminalText)
				}

				column := terminalColumnState.ToGraphQLTerminalColumn()
				colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			}
			// 			// if e.BackgroundImageColumn != nil && e.BackgroundImageColumn.Column == i {
			// 			// 	// if bgColState == nil {
			// 			// 	// 	bgColState = NewBackgroundImageColumn(..., ..., ..., ..., ...)
			// 			// 	// } else {
			// 			// 	// 	bgColState = bgColState.Transition(..., ..., ..., ..., ...)
			// 			// 	// }
			// 			// 	column := ToGraphQLBgImgCol(e.BackgroundImageColumn)
			// 			// 	colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			// 			// }

			// 			// if e.ImageDescriptionColumn != nil && e.ImageDescriptionColumn.Column == i {
			// 			// 	column := ToGraphQLImgDescCol(e.ImageDescriptionColumn)
			// 			// 	colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			// 			// }

			// 			// if e.MarkdownColumn != nil && e.MarkdownColumn.Column == i {
			// 			// 	column := ToGraphQLMarkdownColumn(e.MarkdownColumn)
			// 			// 	colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			// 			// }

			// 			// if e.GitColumn != nil && e.GitColumn.Column == i {
			// 			// 	if srcColState == nil {
			// 			// 		var err error
			// 			// 		srcColState, err = state.NewSourceCodeColumn(e.GitColumn.RepoUrl, e.GitColumn.Commit, e.Step)
			// 			// 		if err != nil {
			// 			// 			// return nil, fmt.Errorf("ToGraphQLPages failed to initialize source code, %s", err)
			// 			// 		}
			// 			// 	} else {
			// 			// 		err := srcColState.Transition(e.Step, e.GitColumn.Commit)
			// 			// 		if err != nil {
			// 			// 			// return nil, fmt.Errorf("ToGraphQLPages failed to transition source code, %s", err)
			// 			// 		}
			// 			// 	}
			// 			// }

			// 			// // once srcColState is initialized, git column persists
			// 			// if srcColState != nil {
			// 			// 	column := srcColState.ToGraphQLSourceCodeColumn()
			// 			// 	colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			// 			// }
			// 		}

			// 		// copy to avoid mutation effects afterwards
			// 		step := internal.StringRef(e.Step)
			// 		prevStep := internal.StringRef(e.PrevStep)
			// 		nextStep := internal.StringRef(e.NextStep)
		}

		page := model.Page{
			Step:     internal.StringRef(currentStep),
			PrevStep: internal.StringRef(prevStep),
			NextStep: internal.StringRef(nextStep),
			Columns:  colWrappers,
		}

		pages = append(pages, page)
	}

	return pages, nil
}

func ReadStepEntries2(filePath string) (StepEntries2, error) {
	funcName := "ReadStepEntries2"
	var entries StepEntries2

	jsonBytes, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("%s failed to read file = %s, %v", funcName, filePath, err)
	}

	json.Unmarshal(jsonBytes, &entries)
	if err != nil {
		return nil, fmt.Errorf("%s failed to unmarshal file = %s, %v", funcName, filePath, err)
	}

	return entries, err
}

func Process2(dirName string) error {
	entries, err := ReadStepEntries2(dirName + "/steps2.json")
	if err != nil {
		return fmt.Errorf("Process2 failed, %s", err)
	}

	pages, err := entries.ToGraphQLPages()
	if err != nil {
		return fmt.Errorf("Process2 failed, %s", err)
	}

	if err := ClearDirectory(dirName); err != nil {
		return fmt.Errorf("Process2 failed, %s", err)
	}

	if err := WriteResults(dirName, pages); err != nil {
		return fmt.Errorf("Process2 failed, %s", err)
	}

	return nil
}
