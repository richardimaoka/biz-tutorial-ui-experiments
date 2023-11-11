package state2

/**
 * Source Code Column type and methods
 */

type SourceCodeColumn struct {
	sourceCode *SourceCode
}

func (c *SourceCodeColumn) InitialCommit(commit string) error {
	return nil
}

func (c *SourceCodeColumn) ForwardCommit(nextCommit string) {
}

func (c *SourceCodeColumn) ShowFileTree() {
}

func (c *SourceCodeColumn) OpenFile(filePath string) {
}

func (c *SourceCodeColumn) Update(fields *SourceFields) {

}
