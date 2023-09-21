package rough

func InnerStateProbe(currentColumn string) *InnerState {
	return &InnerState{
		currentColumn: currentColumn,
	}
}
