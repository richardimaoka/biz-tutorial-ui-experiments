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
}

func (p *TerminalEntry) ToGraphQLTerminalEntry() *model.TerminalEntry {
	// copy to avoid mutation effect afterwards
	m := &model.TerminalEntry{
		ID:        p.id,
		EntryType: p.entryType.ToGraphQL(),
		Text:      p.text,
	}

	return m
}

/**
 * Terminal types
 */

type Terminal struct {
	step             string
	terminalName     string
	currentDirectory string
	entries          []TerminalEntry
	tooltip          *TerminalTooltip
}

func NewTerminal() *Terminal {
	return &Terminal{}
}

func (t *Terminal) AppendEntry(entryType TerminalEntryType, id, text string) {
	t.entries = append(t.entries, TerminalEntry{
		id:        id,
		entryType: entryType,
		text:      text,
	})
}

func (t *Terminal) SetTooltip(tooltipContents string, timing TerminalTooltipTiming) {
	t.tooltip = &TerminalTooltip{
		markdownBody: tooltipContents,
		timing:       timing,
	}
}

func (t *Terminal) ClearTooltip() {
	t.tooltip = nil
}

func (t *Terminal) ChangeCurrentDirectory(dirPath string) {
	t.currentDirectory = dirPath
}

func (t *Terminal) ToGraphQLTerminal() *model.Terminal2 {
	// copy to avoid mutation effect afterwards
	terminalName := stringRef(t.terminalName)

	var modelEntries []*model.TerminalEntry
	for _, e := range t.entries {
		modelEntries = append(modelEntries, &model.TerminalEntry{
			ID:        e.id,
			EntryType: e.entryType.ToGraphQL(),
			Text:      e.text,
		})
	}

	var modelTooltip *model.TerminalTooltip2
	if t.tooltip != nil {
		timing := t.tooltip.timing.ToGraphQL()
		modelTooltip = &model.TerminalTooltip2{
			MarkdownBody: t.tooltip.markdownBody,
			Timing:       &timing,
		}
	}

	return &model.Terminal2{
		Name:             terminalName,
		CurrentDirectory: t.currentDirectory,
		Entries:          modelEntries,
		Tooltip:          modelTooltip,
	}
}
