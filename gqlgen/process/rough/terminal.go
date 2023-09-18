package rough

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
)

func (s *RoughStep) Conversion(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	switch s.Type {
	case "terminal":
		return s.TerminalConvert(state, repo)
	case "commit":
		return s.CommitConvert(state, repo)
	case "source error":
		return s.SourceErrorConvert(state, repo)
	case "browser":
		return s.BrowserConvert(state, repo)
	default:
		return nil, fmt.Errorf("unknown type = '%s', phase = '%s', comment = '%s'", s.Type, s.Phase, s.Comment)
	}
}

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
		state.currentCol = "Terminal"
		commitSteps, err := s.CommitConvert(state, repo)
		if err != nil {
			return nil, fmt.Errorf("failed to convert commit steps, %s", err)
		}
		detailedSteps = append(detailedSteps, commitSteps...)
	}

	return detailedSteps, nil
}

func (s *RoughStep) SourceErrorConvert(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	sourceErrorStep := DetailedStep{
		FocusColumn:         "Source Code",
		DefaultOpenFilePath: s.Instruction, // Go zero value is ""
	}

	detailedSteps = append(detailedSteps, sourceErrorStep)

	return detailedSteps, nil
}

func (s *RoughStep) BrowserConvert(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	if s.Instruction == "" {
		browserStep := DetailedStep{
			FocusColumn: "Browser",
		}
		detailedSteps = append(detailedSteps, browserStep)
	} else {
		split := strings.Split(s.Instruction, ",")
		for _, s := range split {
			browserImageName := strings.ReplaceAll(s, " ", "")
			browserStep := DetailedStep{
				FocusColumn:      "Browser",
				BrowserImageName: browserImageName,
			}
			detailedSteps = append(detailedSteps, browserStep)
		}
	}

	return detailedSteps, nil
}
