package state

type TerminalColumn struct {
	terminal Terminal
}

func NewTerminalColumn() *Terminal {
	return &Terminal{}
}
