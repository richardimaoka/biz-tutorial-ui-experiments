package state2

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

/**
 * Terminal Tooltip type and methods
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

func (t *TerminalTooltip) ToGraphQL() *model.TerminalTooltip2 {
	timing := t.timing.ToGraphQL()
	return &model.TerminalTooltip2{
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

func (t *Terminal) WriteOutput(id, output string) {
	t.entries = append(t.entries, TerminalEntry{
		id:        id,
		entryType: Command,
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

func (t *Terminal) ToGraphQL() *model.Terminal2 {
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

	m := model.Terminal2{
		Name:             terminalName,
		CurrentDirectory: t.currentDirectory,
		Entries:          modelEntries,
	}
	if t.tooltip != nil {
		m.Tooltip = t.tooltip.ToGraphQL()
	}

	return &m
}

/**
 * Terminal Column type and methods
 */

type TerminalColumn struct {
	terminals []*Terminal
}

func (c *TerminalColumn) WriteCommand(
	stepId string,
	name string,
	command string,
	tooltipContents string,
) {
	var terminal *Terminal
	for _, t := range c.terminals {
		if t.terminalName == name {
			terminal = t
		}
	}
	if terminal == nil {
		terminal = NewTerminal(name)
		c.terminals = append(c.terminals, terminal)
	}

	terminal.WriteCommand(stepId, command)
	if tooltipContents == "" {
		terminal.ClearTooltip()
	} else {
		terminal.SetTooltip(tooltipContents, START)
	}
}

func (c *TerminalColumn) WriteOutput(
	stepId string,
	name string,
	output string,
	tooltipContents string,
) {
	var terminal *Terminal
	for _, t := range c.terminals {
		if t.terminalName == name {
			terminal = t
		}
	}
	if terminal == nil {
		terminal = NewTerminal(name)
		c.terminals = append(c.terminals, terminal)
	}

	terminal.WriteOutput(stepId, output)
	if tooltipContents == "" {
		terminal.ClearTooltip()
	} else {
		terminal.SetTooltip(tooltipContents, START)
	}
}

func (c *TerminalColumn) ChangeCurrentDirectory(
	name string,
	dirPath string,
) {
	var terminal *Terminal
	for _, t := range c.terminals {
		if t.terminalName == name {
			terminal = t
		}
	}
	if terminal == nil {
		terminal = NewTerminal(name)
		c.terminals = append(c.terminals, terminal)
	}

	terminal.ChangeCurrentDirectory(dirPath)
}

func (c *TerminalColumn) ToGraphQL() *model.TerminalColumn2 {
	var terminals []*model.Terminal2
	for _, t := range c.terminals {
		terminals = append(terminals, t.ToGraphQL())
	}

	return &model.TerminalColumn2{
		Terminals: terminals,
	}
}
