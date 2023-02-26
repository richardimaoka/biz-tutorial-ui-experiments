package model2

type ActionTerminal struct {
	Command          string
	TerminalName     string
	Output           string //if "", no output after execution
	CurrentDirectory string //if "", current directory is not changed after execution
}
