package state2

import "fmt"

type Page struct {
	terminalColumn   *TerminalColumn
	sourceCodeColumn *SourceCodeColumn
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
