package rough

import "fmt"

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
