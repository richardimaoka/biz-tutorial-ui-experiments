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
	state.currentCol = "Terminal"
	commitSteps, err := s.CommitConvert(state, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to convert commit steps, %s", err)
	}
	detailedSteps = append(detailedSteps, commitSteps...)

	return detailedSteps, nil
}
