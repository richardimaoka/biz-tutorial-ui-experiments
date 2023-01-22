package pkg

type TerminalElement interface {
	String() string
}

type TerminalCommand struct {
	TypeName string `json:"__typename"`
	Command  string
}

type TerminalCommandOutput struct {
	TypeName string `json:"__typename"`
	Output   string
}

type Terminal struct {
	Elements []TerminalElement
}

func (src *Terminal) Copy(dst *Terminal) {
	copy(src.Elements, dst.Elements)
}

func (t *Terminal) AppendElement(elem TerminalElement) {
	t.Elements = append(t.Elements)
}
