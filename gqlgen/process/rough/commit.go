package rough

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func (s *RoughStep) CommitConvert(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	if s.Commit == "" {
		return nil, fmt.Errorf("commit is missing for manual commit, phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
	}

	files, err := gitFilesForCommit(repo, s.Commit)
	if err != nil {
		return nil, fmt.Errorf("failed to get files for commit = %s, %s", s.Commit, err)
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("failed to get files for commit = %s, no files found", s.Commit)
	}

	if state.currentCol != "Source Code" {
		fileTreeStep := DetailedStep{
			FocusColumn:         "Source Code",
			IsFoldFileTree:      false,
			DefaultOpenFilePath: files[0],
			Commit:              s.Commit,
		}
		detailedSteps = append(detailedSteps, fileTreeStep)
	}

	// 3.2. file steps
	for i, file := range files {
		commitStep := DetailedStep{
			FocusColumn:         "Source Code",
			DefaultOpenFilePath: file,
			IsFoldFileTree:      true,
		}
		detailedSteps = append(detailedSteps, commitStep)

		if i == 5 {
			break
		}
	}

	return detailedSteps, nil
}
