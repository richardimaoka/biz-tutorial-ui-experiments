package rough

import (
	"fmt"
	"strings"
)

func (s *RoughStep) TerminalConvert(state *InnerState) ([]DetailedStep, error) {
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
		TerminalText: s.Instruction,
		TerminalType: "command",
		CurrentDir:   currentDir,    // Go zero value is ""
		Commit:       s.Commit,      // Go zero value is ""
		TerminalName: s.Instruction, // Go zero value is ""
	}
	detailedSteps = append(detailedSteps, cmdStep)

	// 2. output step
	if s.Instruction2 != "" {
		outputStep := DetailedStep{
			TerminalText: s.Instruction2,
			TerminalType: "output",
		}
		detailedSteps = append(detailedSteps, outputStep)
	}

	return detailedSteps, nil
}
