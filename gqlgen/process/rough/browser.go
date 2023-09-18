package rough

import (
	"strings"

	"github.com/go-git/go-git/v5"
)

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
