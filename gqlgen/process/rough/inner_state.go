package rough

import (
	"fmt"

	"github.com/google/uuid"
)

type InnerState struct {
	currentColumn         string
	existingCols          []string
	existingDetailedSteps []DetailedStep
}

func NewInnerState(targetFile string) (*InnerState, error) {
	existing, err := readExistingDetailedSteps(targetFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read existing detailed steps, %s", err)
	}

	return &InnerState{
		existingDetailedSteps: existing,
	}, nil
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
