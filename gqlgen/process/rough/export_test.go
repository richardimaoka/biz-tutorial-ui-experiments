package rough

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func PredictableInnerState(currentColumn, targetFile string, repo *git.Repository) *InnerState {
	state, err := NewInnerState(targetFile, repo)
	if err != nil {
		panic(fmt.Errorf("failed to create InnerState: %s", err))
	}

	finder, err := PredictableUUIDFinder(targetFile)
	if err != nil {
		panic(fmt.Errorf("failed to create UUIDFinder: %s", err))
	}

	state.uuidFinder = finder
	state.currentColumn = currentColumn

	return state
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

func (state *InnerState) GenerateTarget(roughStepsFile string) ([]DetailedStep, error) {
	return state.generateTarget(roughStepsFile)
}
