package state2

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

/**
 * Terminal Tooltip types
 */
type TerminalTooltipTiming string

const (
	START TerminalTooltipTiming = "START"
	END   TerminalTooltipTiming = "END"
)

func (t TerminalTooltipTiming) ToGraphQL() model.TerminalTooltipTiming2 {
	switch t {
	case START:
		return model.TerminalTooltipTiming2Start
	case END:
		return model.TerminalTooltipTiming2End
	default:
		panic(fmt.Sprintf("TerminalTooltipTiming = '%s' is invalid", t))
	}
}

type TerminalTooltip struct {
	markdownBody string
	timing       TerminalTooltipTiming
}

/**
 * Terminal Entry types
 */

type TerminalEntryType string

const (
	Command TerminalEntryType = "COMMAND"
	Output  TerminalEntryType = "OUTPUT"
)

func (t TerminalEntryType) ToGraphQL() model.TerminalEntryType {
	switch t {
	case Command:
		return model.TerminalEntryTypeCommand
	case Output:
		return model.TerminalEntryTypeOutput
	default:
		panic(fmt.Sprintf("TerminalEntryType = '%s' is invalid", t))
	}
}

type TerminalEntry struct {
	id        string
	entryType TerminalEntryType
	text      string
	tooltip   *TerminalTooltip
}

func (p *TerminalEntry) ToGraphQLTerminalEntry() *model.TerminalEntry {
	// copy to avoid mutation effect afterwards
	m := &model.TerminalEntry{
		ID:        p.id,
		EntryType: p.entryType.ToGraphQL(),
		Text:      p.text,
	}
	if p.tooltip != nil {
		timing := p.tooltip.timing.ToGraphQL()
		m.Tooltip = &model.TerminalTooltip2{
			MarkdownBody: p.tooltip.markdownBody,
			Timing:       &timing,
		}
	}

	return m
}
