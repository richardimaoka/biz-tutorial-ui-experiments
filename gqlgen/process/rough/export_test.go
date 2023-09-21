package rough

import "fmt"

func InnerStateProbe(currentColumn string) *InnerState {
	return &InnerState{
		currentColumn: currentColumn,
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
