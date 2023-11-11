package state

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Page struct {
	repo       *git.Repository
	tutorial   string
	projectDir string

	terminalColumn   *TerminalColumn
	sourceCodeColumn *SourceColumn
	browserColumn    *BrowserColumn

	columns []Column
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
			p.columns = append(p.columns, p.sourceCodeColumn)
		}
		err := p.sourceCodeColumn.Update(&step.SourceFields)
		if err != nil {
			return fmt.Errorf("Update failed to process step = '%s', %s", step.StepId, err)
		}
		return nil

	case TerminalColumnType:
		if p.terminalColumn == nil {
			p.terminalColumn = NewTerminalColumn()
			p.columns = append(p.columns, p.terminalColumn)
		}
		err := p.terminalColumn.Update(step.StepId, &step.TerminalFields)
		if err != nil {
			return fmt.Errorf("Update failed to process step = '%s', %s", step.StepId, err)
		}
		return nil

	case BrowserColumnType:
		if p.browserColumn == nil {
			p.browserColumn = NewBrowserColumn()
			p.columns = append(p.columns, p.browserColumn)
		}
		err := p.browserColumn.Update(&step.BrowserFields)
		if err != nil {
			return fmt.Errorf("Update failed to process step = '%s', %s", step.StepId, err)
		}
		return nil

	default:
		return fmt.Errorf("Update failed to process step = '%s', column type = '%s' is not implemented", step.StepId, step.FocusColumn)
	}
}

func (p *Page) ToGraphQL() *model.Page2 {
	var modelColumns []*model.ColumnWrapper2
	for _, c := range p.columns {
		modelColumns = append(modelColumns, c.ToGraphQLColumnWrapper())
	}

	return &model.Page2{
		Columns: modelColumns,
	}
}
