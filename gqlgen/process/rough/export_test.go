package rough

import (
	"encoding/json"
	"fmt"

	"github.com/go-git/go-git/v5"
)

var CommitConvertInternal = commitConvertInternal
var TerminalConvertInternal = terminalConvertInternal

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

func StaticUUIDFinder(staticUUID string) *UUIDFinder {
	return &UUIDFinder{
		uuidGenerator: func() string { return "" },
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

func (state *InnerState) GenerateTarget(roughStepsFile string) ([]DetailedStep, error) {
	return state.generateTarget(roughStepsFile)
}

func ToOmitEmptyStructs(dsSlice []DetailedStep) []DetailedStepTest {
	var dstSlice []DetailedStepTest

	for _, ds := range dsSlice {
		jsonBytes, err := json.Marshal(ds)
		if err != nil {
			panic(err)
		}

		var dst DetailedStepTest
		err = json.Unmarshal(jsonBytes, &dst)
		if err != nil {
			panic(err)
		}

		dstSlice = append(dstSlice, dst)
	}

	return dstSlice
}
