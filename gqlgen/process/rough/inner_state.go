package rough

import (
	"github.com/google/uuid"
)

type InnerState struct {
	currentColumn         string
	existingCols          []string
	existingDetailedSteps []DetailedStep
	findUUID              func(rs *RoughStep, subID string) string
}

func NewInnerState(targetFile string) (*InnerState, error) {
	return &InnerState{}, nil
}

func (i *InnerState) FindUUID(rs *RoughStep, subID string) string {
	for _, ds := range i.existingDetailedSteps {
		if ds.FromRoughStep && rs.Step == ds.ParentStep && subID == ds.SubID {
			return ds.Step
		}
	}
	// if not found, then new UUID
	return uuid.NewString()
}
