package process

import (
	"fmt"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

type StepEntry struct {
	// Uppercase fields to allow json dump for testing
	Step                   string                       `json:"step"`
	NColumns               int                          `json:"nColumns"`
	PrevStep               string                       `json:"prevStep,omitempty"`
	NextStep               string                       `json:"nextStep,omitempty"`
	BackgroundImageColumn  *read.BackgroundImageColumn  `json:"backgroundImageColumn,omitempty"`
	ImageDescriptionColumn *read.ImageDescriptionColumn `json:"imageDescriptionColumn,omitempty"`
	MarkdownColumn         *read.MarkdownColumn         `json:"markdownColumn,omitempty"`
	GitColumn              *read.GitColumn              `json:"GitColumn,omitempty"`
}

type StepEntries []StepEntry

func (entries StepEntries) ToGraphQLPages() []model.Page {
	// var bgColState *state.BackgroundImageColumn
	// var imgDescColState *state.ImageDescriptionColumn
	// var markdownColState *state.MarkdownColumn
	// var terminalColState *state.TerminalColumn
	// var srcColState *state.SourceCodeColumn

	var pages []model.Page
	for _, e := range entries {
		// copy to avoid mutation effects afterwards
		step := internal.StringRef(e.Step)
		prevStep := internal.StringRef(e.PrevStep)
		nextStep := internal.StringRef(e.NextStep)

		var colWrappers []*model.ColumnWrapper
		for i := 0; i < e.NColumns; i++ {

			if e.BackgroundImageColumn != nil && e.BackgroundImageColumn.Column == i {
				// if bgColState == nil {
				// 	bgColState = NewBackgroundImageColumn(..., ..., ..., ..., ...)
				// } else {
				// 	bgColState = bgColState.Transition(..., ..., ..., ..., ...)
				// }
				column := ToGraphQLBgImgCol(e.BackgroundImageColumn)
				colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			}

			if e.ImageDescriptionColumn != nil && e.ImageDescriptionColumn.Column == i {
				column := ToGraphQLImgDescCol(e.ImageDescriptionColumn)
				colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			}

			if e.MarkdownColumn != nil && e.MarkdownColumn.Column == i {
				column := ToGraphQLMarkdownColumn(e.MarkdownColumn)
				colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			}

			// if srcColState == nil {
			// 	srcColState = NewSourceCodeColumn(..., ..., ..., ..., ...)
			// } else {
			// 	if e.SourceCodeColumn != nil && e.SourceCodeColumn.Column == i {
			// 		// if srcColState == nil
			// 		// if srcColState
			// 		column := ToGraphQLMarkdownColumn(e.MarkdownColumn)
			// 		colWrappers = append(colWrappers, &model.ColumnWrapper{Column: column})
			// 	}
			// }
		}

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

func ReadStepEntries(dirName string) (StepEntries, error) {
	//------------------------------------
	// 1. read from JSON files
	//------------------------------------
	steps, err := read.ReadSteps(dirName + "/steps.json")
	if err != nil {
		return nil, fmt.Errorf("ReadStepEntries failed, %s", err)
	}

	backgroundImageColumns, err := read.ReadBackgroundImageColumns(dirName + "/bg_columns.json")
	if err != nil {
		return nil, fmt.Errorf("ReadStepEntries failed, %s", err)
	}

	imageDescriptionColumns, err := read.ReadImageDescriptionColumns(dirName + "/img_columns.json")
	if err != nil {
		return nil, fmt.Errorf("ReadStepEntries failed, %s", err)
	}

	markdownColumns, err := read.ReadMarkdownColumns(dirName + "/md_columns.json")
	if err != nil {
		return nil, fmt.Errorf("ReadStepEntries failed, %s", err)
	}

	gitColumns, err := read.ReadGitColumns(dirName + "/git_columns.json")
	if err != nil {
		return nil, fmt.Errorf("ReadStepEntries failed, %s", err)
	}

	//--------------------------------------------------------
	// 2. construct data for each step
	//--------------------------------------------------------
	var entries StepEntries
	for i, step := range steps {
		bgCol := backgroundImageColumns.FindBySeqNo(step.SeqNo)
		imgCol := imageDescriptionColumns.FindBySeqNo(step.SeqNo)
		mdCol := markdownColumns.FindBySeqNo(step.SeqNo)
		gitCol := gitColumns.FindBySeqNo(step.SeqNo)

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

		conv := StepEntry{
			Step:                   currentStep,
			PrevStep:               prevStep,
			NextStep:               nextStep,
			NColumns:               step.NColumns,
			BackgroundImageColumn:  bgCol,
			ImageDescriptionColumn: imgCol,
			MarkdownColumn:         mdCol,
			GitColumn:              gitCol,
		}
		entries = append(entries, conv)
	}

	return entries, nil
}

func (this StepEntries) ClearDirectory(dirName string) error {
	if err := os.RemoveAll(dirName + "/state"); err != nil {
		return fmt.Errorf("ClearDirectory failed, %s", err)
	}
	if err := os.Mkdir(dirName+"/state", 0744); err != nil {
		return fmt.Errorf("ClearDirectory failed, %s", err)
	}

	return nil
}

func (this StepEntries) WriteResults(dirName string, pages []model.Page) error {
	for _, p := range pages {
		filename := fmt.Sprintf("%s/state/%s.json", dirName, *p.Step)
		err := internal.WriteJsonToFile(p, filename)
		if err != nil {
			return fmt.Errorf("WriteResults failed, %s", err)
		}
	}

	return nil
}

func Process(dirName string) error {
	entries, err := ReadStepEntries(dirName)
	if err != nil {
		return fmt.Errorf("Process failed, %s", err)
	}

	if err := entries.ClearDirectory(dirName); err != nil {
		return fmt.Errorf("Process failed, %s", err)
	}

	pages := entries.ToGraphQLPages()

	if err := entries.WriteResults(dirName, pages); err != nil {
		return fmt.Errorf("Process failed, %s", err)
	}

	return nil
}
