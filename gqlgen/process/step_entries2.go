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
	Step            string `json:"step"`
	AutoNextSeconds int    `json:"autoNextSeconds,omitempty"`

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
	CurrentDir   string `json:"currentDir,omitempty"`

	// git
	Commit              string `json:"commit,omitempty"`
	CommitMessage       string `json:"commitMessage,omitempty"`
	PrevCommit          string `json:"prevCommit,omitempty"`
	RepoUrl             string `json:"repoUrl,omitempty"`
	DefaultOpenFilePath string `json:"defaultOpenFilePath,omitempty"`
	IsFoldFileTree      string `json:"isFoldFileTree,omitempty"` // string, as CSV from Google Spreadsheet has TRUE as upper-case 'TRUE'

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
	} else if seqNo == 1 {
		prevStep = "_initial"
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

func (entries StepEntries2) ToGraphQLPages(tutorial string) ([]model.Page, error) {
	var srcColmnState *state.SourceCodeColumn
	var terminalColumnState *state.TerminalColumn

	var pages []model.Page
	for seqNo, e := range entries {
		currentStep, prevStep, nextStep := entries.calcSteps(seqNo)

		columns := e.columns(seqNo)
		if len(columns) == 0 {
			return nil, fmt.Errorf("ToGraphQLPages failed at step = %s, no columns are specified", e.Step)
		}

		var colWrappers []*model.ColumnWrapper
		for _, colName := range columns {
			if colName == "Terminal" {
				if terminalColumnState == nil {
					terminalColumnState = state.NewTerminalColumn()
				}

				if e.TerminalText == "" {
					terminalColumnState.MarkAllExecuted()
				} else {
					terminalType, err := state.ToTerminalElementType(e.TerminalType)
					if err != nil {
						return nil, fmt.Errorf("ToGraphQLPages failed at step = %s to convert terminal type, %s", e.Step, err)
					}
					terminalColumnState.Transition(terminalType, e.TerminalText)
				}

				if e.CurrentDir != "" {
					terminalColumnState.Cd(e.CurrentDir)
				}

				column := terminalColumnState.ToGraphQLTerminalColumn()
				colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column, Name: internal.StringRef(colName)})
			}

			if colName == "Source Code" {
				isFoldFileTree := e.IsFoldFileTree == "TRUE"
				if srcColmnState == nil {
					var err error
					srcColmnState, err = state.NewSourceCodeColumn(e.RepoUrl, e.Commit, e.Step, e.DefaultOpenFilePath, tutorial, isFoldFileTree)
					if err != nil {
						return nil, fmt.Errorf("ToGraphQLPages failed at step = %s to initialize source code, %s", e.Step, err)
					}
				} else if e.Commit != "" {
					err := srcColmnState.Transition(e.Step, e.Commit, e.DefaultOpenFilePath, isFoldFileTree)
					if err != nil {
						return nil, fmt.Errorf("ToGraphQLPages failed at step %s to transition source code, %s", e.Step, err)
					}
				}

				column := srcColmnState.ToGraphQLSourceCodeColumn()
				colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column, Name: internal.StringRef(colName)})
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

		modalText := e.ModalText
		modalPosition, _ := state.ToModalPosition(e.ModalPosition) // err is ignored, and modal position will be null
		modalState := state.Modal{Text: modalText, Position: modalPosition}

		autoNextSeconds := e.AutoNextSeconds

		page := model.Page{
			Step:            internal.StringRef(currentStep),
			PrevStep:        internal.StringRef(prevStep),
			NextStep:        internal.StringRef(nextStep),
			AutoNextSeconds: &autoNextSeconds,
			Columns:         colWrappers,
			FocusColumn:     internal.StringRef(e.FocusColumn),
			Modal:           modalState.ToGraphQLModal(),
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

func Process2(tutorial string) error {
	dirName := "data/" + tutorial

	entries, err := ReadStepEntries2(dirName + "/steps2.json")
	if err != nil {
		return fmt.Errorf("Process2 failed, %s", err)
	}

	pages, err := entries.ToGraphQLPages(tutorial)
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
