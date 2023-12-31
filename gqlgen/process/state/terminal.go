package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

/**
 * Terminal Tooltip type and methods
 */
type TerminalTooltipTiming string

const (
	TERMINAL_TOOLTIP_START TerminalTooltipTiming = "START"
	TERMINAL_TOOLTIP_END   TerminalTooltipTiming = "END"
)

func (t TerminalTooltipTiming) ToGraphQL() model.TerminalTooltipTiming {
	switch t {
	case TERMINAL_TOOLTIP_START:
		return model.TerminalTooltipTimingStart
	case TERMINAL_TOOLTIP_END:
		return model.TerminalTooltipTimingEnd
	default:
		panic(fmt.Sprintf("TerminalTooltipTiming = '%s' is invalid", t))
	}
}

type TerminalTooltip struct {
	markdownBody string
	timing       TerminalTooltipTiming
}

func (t TerminalTooltipTiming) toGraphQL() model.TerminalTooltipTiming {
	switch t {
	case TERMINAL_TOOLTIP_START:
		return model.TerminalTooltipTimingStart
	case TERMINAL_TOOLTIP_END:
		return model.TerminalTooltipTimingEnd
	default:
		panic(fmt.Sprintf("TerminalTooltipTiming has an invalid value = '%s'", t))
	}
}

func (t *TerminalTooltip) ToGraphQL() *model.TerminalTooltip {
	timing := t.timing.ToGraphQL()
	return &model.TerminalTooltip{
		Timing:       &timing,
		MarkdownBody: t.markdownBody,
	}
}

/**
 * Terminal Entry type and methods
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
	id                string
	entryType         TerminalEntryType
	text              string
	isCommandExecuted bool
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
 * Terminal type and methods
 */

type Terminal struct {
	terminalName     string
	currentDirectory string
	entries          []TerminalEntry
	tooltip          *TerminalTooltip
}

func NewTerminal(name string) *Terminal {
	return &Terminal{terminalName: name}
}

func (t *Terminal) WriteCommand(id, command string) {
	t.entries = append(t.entries, TerminalEntry{
		id:        id,
		entryType: Command,
		text:      command,
	})
}

func (t *Terminal) ExecuteLastCommand() error {
	errorPrefix := "Terminal ExecuteLastCommand() failed"

	if len(t.entries) == 0 {
		return fmt.Errorf("%s, terminal = '%s' is empty", errorPrefix, t.terminalName)
	}

	// Go pointer technique to change `t.entries[]` fields
	lastEntry := &t.entries[len(t.entries)-1]
	//   if we do `lastEntry := t.entries[...]`, then `lastEntry.isCommandExecuted = false`
	//   will not change the original `t.entries[]`

	if lastEntry.entryType != Command {
		return fmt.Errorf(
			"%s, last element (id = '%s') of terminal '%s' is expected entryType = Command but '%s'",
			errorPrefix,
			lastEntry.id,
			t.terminalName,
			lastEntry.entryType,
		)
	}

	lastEntry.isCommandExecuted = true
	return nil
}

func (t *Terminal) WriteOutput(id, output string) {
	t.entries = append(t.entries, TerminalEntry{
		id:        id,
		entryType: Output,
		text:      output,
	})
}

func (t *Terminal) ChangeCurrentDirectory(dirPath string) {
	t.currentDirectory = dirPath
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

func (t *Terminal) ToGraphQL() *model.Terminal {
	// copy to avoid mutation effect afterwards
	terminalName := stringRef(t.terminalName)

	var modelEntries []*model.TerminalEntry
	for _, e := range t.entries {
		executed := e.isCommandExecuted
		modelEntries = append(modelEntries, &model.TerminalEntry{
			ID:                e.id,
			EntryType:         e.entryType.ToGraphQL(),
			Text:              e.text,
			IsCommandExecuted: &executed,
		})
	}

	m := model.Terminal{
		Name:             terminalName,
		CurrentDirectory: t.currentDirectory,
		Entries:          modelEntries,
	}

	if t.tooltip != nil {
		m.Tooltip = t.tooltip.ToGraphQL()
	}

	return &m
}
