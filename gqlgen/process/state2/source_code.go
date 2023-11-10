package state2

import "github.com/go-git/go-git/v5"

type SourceCodeTooltipTiming string

const (
	SOURCE_TOOLTIP_START SourceCodeTooltipTiming = "START"
	SOURCE_TOOLTIP_END   SourceCodeTooltipTiming = "END"
)

type SourceCodeTooltip struct {
	markdownBody string
	timing       SourceCodeTooltipTiming
	lineNumber   int
}

type SourceCode struct {
	// metadata, can be set only at initialization
	tutorial   string
	projectDir string
	repo       *git.Repository

	// inner state updated at each step
	commitHash string
	rootDir    *Directory
	step       string
	tooltip    *SourceCodeTooltip

	// metadata, can be set from caller anytime
	DefaultOpenFilePath string
	showFileTree        bool
}

func NewSourceCode(repo *git.Repository, projectDir, tutorial string) *SourceCode {
	return &SourceCode{
		repo:       repo,
		projectDir: projectDir,
		tutorial:   tutorial,
	}
}

type SourceCodeColumn struct {
	sourceCode *SourceCode
}

func (c *SourceCodeColumn) ForwardCommit(nextCommit string) {
}

func (c *SourceCodeColumn) ShowFileTree() {
}

func (c *SourceCodeColumn) OpenFile(filePath string) {
}
