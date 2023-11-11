package result

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state2"
)

func toTerminalColumn(c *state2.TerminalColumn, stepId string, fields *TerminalFields) error {
	switch fields.TerminalStepType {
	case TerminalCommand:
		c.WriteOutput(stepId, fields.TerminalName, fields.TerminalText, fields.TerminalTooltipContents)
	case TerminalOutput:
		c.WriteOutput(stepId, fields.TerminalName, fields.TerminalText, fields.TerminalTooltipContents)
	case TerminalCd:
		c.ChangeCurrentDirectory(fields.TerminalName, fields.CurrentDir)
	case TerminalMove:
		// no update is needed
	case TerminalOpen:
		// no update is needed
	default:
		return fmt.Errorf("toTerminalColumn failed, type = '%s' is not implemented", fields.TerminalStepType)
	}

	return nil
}
