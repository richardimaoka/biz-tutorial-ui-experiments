package process

import (
	"fmt"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/read"
)

type StepEntries []StepEntry

func (this StepEntries) ToGraphQLPages() []model.Page {
	var pages []model.Page
	for _, e := range this {
		// copy to avoid mutation effects afterwards
		step := internal.StringRef(e.Step)
		prevStep := internal.StringRef(e.PrevStep)
		nextStep := internal.StringRef(e.NextStep)
		columns := e.ToGraphQLColumns()

		page := model.Page{
			Step:     step,
			PrevStep: prevStep,
			NextStep: nextStep,
			Columns:  columns,
		}

		pages = append(pages, page)
	}

	return pages
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

func (this StepEntries) WriteResults(dirName string) error {
	for _, p := range this.ToGraphQLPages() {
		filename := fmt.Sprintf("%s/state/%s.json", dirName, *p.Step)
		err := internal.WriteJsonToFile(p, filename)
		if err != nil {
			return fmt.Errorf("WriteResults failed, %s", err)
		}
	}

	return nil
}

func ReadStepEntries(dirName string) (StepEntries, error) {
	//------------------------------------
	// 1. read from JSON files
	//------------------------------------
	steps, err := read.ReadSteps(dirName + "/steps.json")
	if err != nil {
		return nil, fmt.Errorf("InitialRead failed, %s", err)
	}

	backgroundImageColumns, err := read.ReadBackgroundImageColumns(dirName + "/bg_columns.json")
	if err != nil {
		return nil, fmt.Errorf("InitialRead failed, %s", err)
	}

	imageDescriptionColumns, err := read.ReadImageDescriptionColumns(dirName + "/img_columns.json")
	if err != nil {
		return nil, fmt.Errorf("InitialRead failed, %s", err)
	}

	markdownColumns, err := read.ReadMarkdownColumns(dirName + "/md_columns.json")
	if err != nil {
		return nil, fmt.Errorf("InitialRead failed, %s", err)
	}

	//--------------------------------------------------------
	// 2. construct data for each step
	//--------------------------------------------------------
	var entries StepEntries
	for i, step := range steps {
		bgCol := backgroundImageColumns.FindBySeqNo(step.SeqNo)
		imgCol := imageDescriptionColumns.FindBySeqNo(step.SeqNo)
		mdCol := markdownColumns.FindBySeqNo(step.SeqNo)

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
		}
		entries = append(entries, conv)
	}

	return entries, nil
}

func Process(dirName string) error {
	entries, err := ReadStepEntries(dirName)
	if err != nil {
		return fmt.Errorf("Process failed, %s", err)
	}

	if err := entries.ClearDirectory(dirName); err != nil {
		return fmt.Errorf("Process failed, %s", err)
	}

	if err := entries.WriteResults(dirName); err != nil {
		return fmt.Errorf("Process failed, %s", err)
	}
	return nil
}
