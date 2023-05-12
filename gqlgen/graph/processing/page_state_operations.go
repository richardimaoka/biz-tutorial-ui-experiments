package processing

type PageStateOperation struct {
	FileOps           []FileSystemOperation
	TerminalOperation TerminalOperation
}
