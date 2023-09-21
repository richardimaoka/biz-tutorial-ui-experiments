package rough

import "fmt"

type InnerState struct {
	currentColumn string
	existingCols  []string
	uuidFinder    *UUIDFinder
}

func NewInnerState(targetFile string) (*InnerState, error) {
	finder, err := NewUUIDFinder(targetFile)
	if err != nil {
		return nil, fmt.Errorf("failed to create UUIDFinder: %s", err)
	}

	return &InnerState{uuidFinder: finder}, nil
}
