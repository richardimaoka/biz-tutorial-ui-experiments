package pkg

type Terminal struct {
	Elements []TerminalElement
}

type TerminalElement interface {
	String() string
}

type TerminalCommand struct {
	TypeName string `json:"__typename"`
	Command  string
}

func (elem TerminalCommand) String() string {
	return elem.Command
}

type TerminalCommandOutput struct {
	TypeName string `json:"__typename"`
	Output   string
}

func InitialTerminal() Terminal {
	return Terminal{}
}

func CopyTerminal(src *Terminal, dst *Terminal) {
	copy(src.Elements, dst.Elements)
}

func ConvToTerminalElement(command Command) TerminalElement {
	return TerminalCommand{TypeName: command.TypeName, Command: command.Command}
}

func (t *Terminal) AppendElement(elem TerminalElement) {
	t.Elements = append(t.Elements)
}
