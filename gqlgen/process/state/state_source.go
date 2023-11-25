package state

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/gitwrap"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/edits"
)

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
		rootDir:    emptyDirectory(""),
	}
}

func commitFiles(repo *git.Repository, commitHash string) (Files, error) {
	fileObjs, err := gitwrap.GetCommitFiles(repo, commitHash)
	if err != nil {
		return nil, fmt.Errorf("commitFiles failed, %s", err)
	}

	var files Files
	for _, fileObj := range fileObjs {
		file, err := fileUnChanged(&fileObj, fileObj.Name)
		if err != nil {
			return nil, fmt.Errorf("commitFiles failed, %s", err)
		}

		files = append(files, file)
	}

	return files, nil
}

func (s *SourceCode) forwardCommit(commitHash string) error {
	if s.commitHash == commitHash {
		// if same commit, do nothing
		return nil
	} else if s.commitHash == "" {
		// if initial commit
		return s.initialCommit(commitHash)
	} else {
		// if non-initial commit
		return s.nonInitialCommit(commitHash)
	}
}

func (s *SourceCode) initialCommit(commitHash string) error {
	files, err := commitFiles(s.repo, commitHash)
	if err != nil {
		return fmt.Errorf("initialCommit failed, %s", err)
	}

	s.rootDir, err = constructDirectory(files)
	if err != nil {
		return fmt.Errorf("initialCommit failed, %s", err)
	}

	// Increment the commit
	s.commitHash = commitHash
	return nil
}

func (s *SourceCode) nonInitialCommit(nextCommitHash string) error {
	files, err := commitFiles(s.repo, nextCommitHash)
	if err != nil {
		return fmt.Errorf("nonInitialCommit failed, %s", err)
	}

	patch, err := gitwrap.GetPatch(s.repo, s.commitHash, nextCommitHash)
	if err != nil {
		return fmt.Errorf("nonInitialCommit failed, %s", err)
	}

	s.rootDir, err = constructDirectory(files)
	if err != nil {
		return fmt.Errorf("nonInitialCommit failed, %s", err)
	}

	// this calculates backword - from current to prev, but it makes the logic so much simpler than forward calculation
	for _, p := range patch.FilePatches() {
		from, to := p.Files() // See Files() method's comment about when 'from' and 'to' become nil
		if from == nil {
			//added
			file, err := s.rootDir.findFile(to.Path())
			if err != nil {
				return fmt.Errorf("nonInitialCommit failed, %s", err)
			}
			file.markAdded()
		} else if to == nil {
			// deleted
			file, err := s.rootDir.findFile(from.Path())
			if err != nil {
				return fmt.Errorf("nonInitialCommit failed, %s", err)
			}
			file.markDeleted()
		} else if from.Path() != to.Path() {
			// renamed
			file, err := s.rootDir.findFile(to.Path())
			if err != nil {
				return fmt.Errorf("nonInitialCommit failed, %s", err)
			}
			file.markRenamed(from.Path())
		} else {
			// updated
			file, err := s.rootDir.findFile(to.Path())
			if err != nil {
				return fmt.Errorf("nonInitialCommit failed, %s", err)
			}

			fileBlob, err := object.GetBlob(s.repo.Storer, from.Hash())
			if err != nil {
				return fmt.Errorf("nonInitialCommit failed, %s", err)
			}

			fileObj := object.NewFile(from.Path(), from.Mode(), fileBlob)
			oldContents, err := fileObj.Contents()
			if err != nil {
				return fmt.Errorf("nonInitialCommit failed, %s", err)
			}

			filePatch := gitwrap.ToFilePatch(p)
			editOps := edits.ToOperations(filePatch.Chunks)

			file.markUpdated(oldContents, editOps)
		}
	}

	// Increment the commit
	s.commitHash = nextCommitHash
	return nil
}

func (s *SourceCode) openFile(filePath string) {
	s.DefaultOpenFilePath = filePath
}

func (s *SourceCode) newTooltip(contents string, timing SourceCodeTooltipTiming, lineNumber int) {
	s.tooltip = &SourceCodeTooltip{
		markdownBody: contents,
		timing:       timing,
		lineNumber:   lineNumber,
	}
}

func (s *SourceCode) appendTooltipContents(additionalContents string) error {
	if s.tooltip == nil {
		return fmt.Errorf("appendTooltipContents failed, cannot append tooltip since the prev tooltip is empty")
	}

	s.tooltip.markdownBody += "\n" + additionalContents
	s.tooltip.timing = SOURCE_TOOLTIP_START

	return nil
}

func (s *SourceCode) ToGraphQL() *model.SourceCode {
	return &model.SourceCode{
		FileTree:            s.rootDir.ToGraphQLFileNodeSlice(),
		FileContents:        s.rootDir.ToGraphQLOpenFileMap(),
		Tutorial:            s.tutorial,
		Step:                s.step,
		DefaultOpenFilePath: s.DefaultOpenFilePath,
		IsFoldFileTree:      !s.showFileTree,
		ProjectDir:          s.projectDir,
	}
}
