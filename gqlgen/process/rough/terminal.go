package rough

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
)

func (s *RoughStep) TerminalConvert(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	// 0. check if it's a valid terminal step
	if s.Instruction == "" && s.Instruction2 == "" {
		return nil, fmt.Errorf("step is missing both 'instruction' and 'instruction2', phase = '%s', type = '%s', comment = '%s'", s.Phase, s.Type, s.Comment)
	}

	// 1.   command step
	// 1.1. check if it's a 'cd' command
	var currentDir string
	if strings.HasPrefix(s.Instruction, "cd ") {
		currentDir = strings.TrimPrefix(s.Instruction, "cd ")
	}
	// 1.2. create command step
	cmdStep := DetailedStep{
		FocusColumn:  "Terminal",
		TerminalType: "command",
		TerminalText: s.Instruction,
		TerminalName: s.Instruction3, // Go zero value is ""
		CurrentDir:   currentDir,     // Go zero value is ""
		Commit:       s.Commit,       // Go zero value is ""
	}
	detailedSteps = append(detailedSteps, cmdStep)

	// 2. output step
	if s.Instruction2 != "" {
		outputStep := DetailedStep{
			FocusColumn:  "Terminal",
			TerminalType: "output",
			TerminalText: s.Instruction2,
		}
		detailedSteps = append(detailedSteps, outputStep)
	}

	// 3. source code steps
	if s.Commit != "" {
		files, err := gitFilesForCommit(repo, s.Commit)
		if err != nil {
			return nil, fmt.Errorf("failed to get files for commit = %s, %s", s.Commit, err)
		}
		if len(files) == 0 {
			return nil, fmt.Errorf("failed to get files for commit = %s, no files found", s.Commit)
		}

		// 3.1. file tree step
		sourceCodeStep := DetailedStep{
			FocusColumn:         "Source Code",
			IsFoldFileTree:      false,
			DefaultOpenFilePath: files[0],
			Commit:              s.Commit,
		}
		detailedSteps = append(detailedSteps, sourceCodeStep)

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
	}

	return detailedSteps, nil
}
