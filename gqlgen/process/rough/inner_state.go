package rough

type InnerState struct {
	currentColumn string
	existingCols  []string
	uuidFinder    *UUIDFinder
}

func NewInnerState(targetFile string) (*InnerState, error) {
	return &InnerState{}, nil
}
