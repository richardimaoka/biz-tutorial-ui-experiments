package processing

type PageStateOperation struct {
	SourceCodeOperation SourceCodeOperation
	TerminalOperation   TerminalOperation
	MarkdownOperation   *MarkdownOperation //currently a concrete struct, so making it a pointer to allow nil
}
