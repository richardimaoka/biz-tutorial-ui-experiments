package rough

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func PredictableInnerState(currentColumn, targetFile string) *InnerState {
	finder, err := PredictableUUIDFinder(targetFile)
	if err != nil {
		panic(fmt.Errorf("failed to create UUIDFinder: %s", err))
	}

	return &InnerState{
		currentColumn: currentColumn,
		uuidFinder:    finder,
	}
}

func PredictableUUIDFinder(targetFile string) (*UUIDFinder, error) {
	finder, err := NewUUIDFinder(targetFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create finder, %s", err)
	}

	// replace the generator to always return empty string
	finder.uuidGenerator = func() string { return "" }

	return finder, nil
}

func (state *InnerState) GenerateTarget(roughStepsFile string, repo *git.Repository) ([]DetailedStep, error) {
	return state.generateTarget(roughStepsFile, repo)
}
