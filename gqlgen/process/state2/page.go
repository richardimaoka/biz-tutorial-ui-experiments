package state2

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Page struct {
	repo             *git.Repository
	tutorial         string
	projectDir       string
	terminalColumn   *TerminalColumn
	sourceCodeColumn *SourceColumn
}

func NewPage(repo *git.Repository, tutorial string) *Page {
	return &Page{
		repo:     repo,
		tutorial: tutorial,
	}
}

func (p *Page) Update(step *Step) error {
	switch step.FocusColumn {
	case SourceColumnType:
		if p.sourceCodeColumn == nil {
			p.sourceCodeColumn = NewSourceColumn(p.repo, p.projectDir, p.tutorial)
		}
		err := p.sourceCodeColumn.Update(&step.SourceFields)
		if err != nil {
			return fmt.Errorf("Update failed to process step = '%s', %s", step.StepId, err)
		}
		return nil

	case TerminalColumnType:
		if p.terminalColumn == nil {
			p.terminalColumn = NewTerminalColumn()
		}
		err := p.terminalColumn.Update(step.StepId, &step.TerminalFields)
		if err != nil {
			return fmt.Errorf("Update failed to process step = '%s', %s", step.StepId, err)
		}
		return nil

	default:
		return fmt.Errorf("Update failed to process step = '%s', column type = '%s' is not implemented", step.StepId, step.FocusColumn)
	}
}

func (p *Page) ToGraphQL() *model.Page2 {
	return nil
}
