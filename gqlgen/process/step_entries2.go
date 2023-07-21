package process

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

type StepEntry2 struct {
	// Uppercase fields to allow json dump for testing

	// steps
	Step     string `json:"step"`
	NColumns int    `json:"nColumns"`
	PrevStep string `json:"prevStep,omitempty"`
	NextStep string `json:"nextStep,omitempty"`

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

func (entries StepEntries2) ToGraphQLPages() []model.Page {
	// var srcColState *state.SourceCodeColumn

	var pages []model.Page
	for _, e := range entries {
		var colWrappers []*model.ColumnWrapper
		for i := 0; i < e.NColumns; i++ {
			// if e.BackgroundImageColumn != nil && e.BackgroundImageColumn.Column == i {
			// 	// if bgColState == nil {
			// 	// 	bgColState = NewBackgroundImageColumn(..., ..., ..., ..., ...)
			// 	// } else {
			// 	// 	bgColState = bgColState.Transition(..., ..., ..., ..., ...)
			// 	// }
			// 	column := ToGraphQLBgImgCol(e.BackgroundImageColumn)
			// 	colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			// }

			// if e.ImageDescriptionColumn != nil && e.ImageDescriptionColumn.Column == i {
			// 	column := ToGraphQLImgDescCol(e.ImageDescriptionColumn)
			// 	colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			// }

			// if e.MarkdownColumn != nil && e.MarkdownColumn.Column == i {
			// 	column := ToGraphQLMarkdownColumn(e.MarkdownColumn)
			// 	colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			// }

			// if e.GitColumn != nil && e.GitColumn.Column == i {
			// 	if srcColState == nil {
			// 		var err error
			// 		srcColState, err = state.NewSourceCodeColumn(e.GitColumn.RepoUrl, e.GitColumn.Commit, e.Step)
			// 		if err != nil {
			// 			// return nil, fmt.Errorf("ToGraphQLPages failed to initialize source code, %s", err)
			// 		}
			// 	} else {
			// 		err := srcColState.Transition(e.Step, e.GitColumn.Commit)
			// 		if err != nil {
			// 			// return nil, fmt.Errorf("ToGraphQLPages failed to transition source code, %s", err)
			// 		}
			// 	}
			// }

			// // once srcColState is initialized, git column persists
			// if srcColState != nil {
			// 	column := srcColState.ToGraphQLSourceCodeColumn()
			// 	colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			// }
		}

		// copy to avoid mutation effects afterwards
		step := internal.StringRef(e.Step)
		prevStep := internal.StringRef(e.PrevStep)
		nextStep := internal.StringRef(e.NextStep)

		page := model.Page{
			Step:     step,
			PrevStep: prevStep,
			NextStep: nextStep,
			Columns:  colWrappers,
		}

		pages = append(pages, page)
	}

	return pages
}

func ReadStepEntries2(dirName string) (StepEntries2, error) {
	//------------------------------------
	// 1. read from JSON files
	//------------------------------------
	steps, err := read.ReadSteps(dirName + "/steps2.json")
	if err != nil {
		return nil, fmt.Errorf("ReadStepEntries2 failed, %s", err)
	}

	//--------------------------------------------------------
	// 2. construct data for each step
	//--------------------------------------------------------
	var entries StepEntries2
	for i, step := range steps {
		var currentStep string
		if i == 0 {
			currentStep = "_initial"
		} else {
			currentStep = steps[i].Step
		}

		var prevStep string
		if i == 0 {
			prevStep = ""
		} else {
			prevStep = entries[i-1].Step
		}

		var nextStep string
		if i == len(steps)-1 {
			nextStep = ""
		} else {
			nextStep = steps[i+1].Step
		}

		entry := StepEntry2{
			Step:     currentStep,
			PrevStep: prevStep,
			NextStep: nextStep,
			NColumns: step.NColumns,
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

func Process2(dirName string) error {
	entries, err := ReadStepEntries2(dirName)
	if err != nil {
		return fmt.Errorf("Process failed, %s", err)
	}

	if err := ClearDirectory(dirName); err != nil {
		return fmt.Errorf("Process failed, %s", err)
	}

	pages := entries.ToGraphQLPages()

	if err := WriteResults(dirName, pages); err != nil {
		return fmt.Errorf("Process failed, %s", err)
	}

	return nil
}
