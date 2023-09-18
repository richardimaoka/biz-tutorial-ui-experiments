package rough

import (
	"github.com/go-git/go-git/v5"
)

func (s *RoughStep) SourceErrorConvert(state *InnerState, repo *git.Repository) ([]DetailedStep, error) {
	var detailedSteps []DetailedStep

	sourceErrorStep := DetailedStep{
		FocusColumn:         "Source Code",
		DefaultOpenFilePath: s.Instruction, // Go zero value is ""
	}

	detailedSteps = append(detailedSteps, sourceErrorStep)

	return detailedSteps, nil
}
