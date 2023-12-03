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
	commitHash      string
	rootDir         *Directory
	step            string
	tooltipFilePath string

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
	funcName := "commitFiles()"

	fileObjs, err := gitwrap.GetCommitFiles(repo, commitHash)
	if err != nil {
		return nil, fmt.Errorf("%s failed, %s", funcName, err)
	}

	var files Files
	for _, fileObj := range fileObjs {
		file, err := fileUnChanged(&fileObj, fileObj.Name)
		if err != nil {
			return nil, fmt.Errorf("%s failed, %s", funcName, err)
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
	funcName := "initialCommit()"
	files, err := commitFiles(s.repo, commitHash)
	if err != nil {
		return fmt.Errorf("%s failed, %s", funcName, err)
	}

	s.rootDir, err = constructDirectory(files)
	if err != nil {
		return fmt.Errorf("%s failed, %s", funcName, err)
	}

	// Increment the commit
	s.commitHash = commitHash
	return nil
}

func (s *SourceCode) nonInitialCommit(nextCommitHash string) error {
	funcName := "nonInitialCommit()"

	files, err := commitFiles(s.repo, nextCommitHash)
	if err != nil {
		return fmt.Errorf("%s failed, %s", funcName, err)
	}

	patch, err := gitwrap.GetPatch(s.repo, s.commitHash, nextCommitHash)
	if err != nil {
		return fmt.Errorf("%s failed, %s", funcName, err)
	}

	// Construct the root directory from the next commit...
	s.rootDir, err = constructDirectory(files)
	if err != nil {
		return fmt.Errorf("%s failed, %s", funcName, err)
	}

	// ...then update files with the patch between current and next
	err = s.processPatch(patch)
	if err != nil {
		return fmt.Errorf("%s failed, %s", funcName, err)
	}

	// Increment the commit
	s.commitHash = nextCommitHash
	return nil
}

// This calculates backword - (i.e.) from current commit to prev commit,
// which might feel counter intuitive, but it makes the logic so much
// simpler than forward calculation
func (s *SourceCode) processPatch(patch *object.Patch) error {
	funcName := "processPatch()"

	for _, p := range patch.FilePatches() {
		from, to := p.Files() // See Files() method's comment about when 'from' and 'to' become nil
		if from == nil {
			// Added
			file, err := s.rootDir.findFile(to.Path())
			if err != nil {
				return fmt.Errorf("%s failed, %s", funcName, err)
			}
			file.markAdded()
		} else if to == nil {
			// Deleted
			file, err := s.rootDir.findFile(from.Path())
			if err != nil {
				return fmt.Errorf("%s failed, %s", funcName, err)
			}
			file.markDeleted()
		} else if from.Path() != to.Path() {
			// Renamed
			file, err := s.rootDir.findFile(to.Path())
			if err != nil {
				return fmt.Errorf("%s failed, %s", funcName, err)
			}
			file.markRenamed(from.Path())
		} else {
			// Updated
			file, err := s.rootDir.findFile(to.Path())
			if err != nil {
				return fmt.Errorf("%s failed, %s", funcName, err)
			}

			fileBlob, err := object.GetBlob(s.repo.Storer, from.Hash())
			if err != nil {
				return fmt.Errorf("%s failed, %s", funcName, err)
			}

			fileObj := object.NewFile(from.Path(), from.Mode(), fileBlob)
			oldContents, err := fileObj.Contents()
			if err != nil {
				return fmt.Errorf("%s failed, %s", funcName, err)
			}

			filePatch := gitwrap.ToFilePatch(p)
			editOps := edits.ToOperations(filePatch.Chunks)

			file.markUpdated(oldContents, editOps)
		}
	}

	return nil
}

func (s *SourceCode) openFile(filePath string) {
	s.DefaultOpenFilePath = filePath
}

func (s *SourceCode) newTooltip(filePath, contents string, timing SourceCodeTooltipTiming, lineNumber int) error {
	funcName := "newTooltip()"

	file, err := s.rootDir.findFile(filePath)
	if err != nil {
		return fmt.Errorf("%s failed, filePath = '%s' not found, %s", funcName, filePath, err)
	}

	file.tooltip = &SourceCodeTooltip{
		markdownBody: contents,
		timing:       timing,
		lineNumber:   lineNumber,
	}

	return nil
}

func (s *SourceCode) appendTooltipContents(additionalContents string) error {
	funcName := "appendTooltipContents()"

	file, err := s.rootDir.findFile(s.tooltipFilePath)
	if err != nil {
		return fmt.Errorf("%s failed, filePath = '%s' not found, %s", funcName, s.tooltipFilePath, err)
	}

	if file.tooltip == nil {
		return fmt.Errorf("%s failed, filePath = '%s' not found", funcName, s.tooltipFilePath)
	}

	file.tooltip.markdownBody += "\n" + additionalContents
	file.tooltip.timing = SOURCE_TOOLTIP_START

	return nil
}

func (s *SourceCode) ClearTooltip() error {
	funcName := "ClearTooltip()"

	if s.tooltipFilePath != "" {
		file, err := s.rootDir.findFile(s.tooltipFilePath)
		if err != nil {
			return fmt.Errorf("%s failed, filePath = '%s' not found, %s", funcName, s.tooltipFilePath, err)
		}

		if file.tooltip == nil {
			return fmt.Errorf("%s failed, filePath = '%s' not found", funcName, s.tooltipFilePath)
		}

		file.tooltip = nil
	}

	s.tooltipFilePath = ""

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
