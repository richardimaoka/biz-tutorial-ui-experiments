package state2

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Page struct {
	repo             *git.Repository
	tutorial         string
	terminalColumn   *TerminalColumn
	sourceCodeColumn *SourceCodeColumn
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
		p.sourceCodeColumn.Update(&step.SourceFields)
		return nil
	case TerminalColumnType:
		err := p.terminalColumn.Update(step.StepId, &step.TerminalFields)
		if err != nil {
			return fmt.Errorf("Update faield, %s", err)
		}
		return nil
	default:
		return fmt.Errorf("Update faield, column type = '%s' is not implemented", step.FocusColumn)
	}
}

func (p *Page) ToGraphQL() *model.Page2 {
	return nil
}
