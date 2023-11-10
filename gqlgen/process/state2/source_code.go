package state2

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
)

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

func (s *SourceCode) initialCommit(commitHash string) error {
	files, err := gitwrap.GetCommitFiles(s.repo, commitHash)
	if err != nil {
		return fmt.Errorf("initialCommit failed, %s", err)
	}

	s.rootDir, err = constructDirectory(files)
	if err != nil {
		return fmt.Errorf("initialCommit failed, %s", err)
	}

	return nil
}

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
