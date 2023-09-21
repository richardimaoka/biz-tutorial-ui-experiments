package rough

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type UUIDGenerator interface {
	FindOrGenerateUUID(rs *RoughStep, subID string) string
}

type UUIDGenFromTarget struct {
	existingSteps []DetailedStep
}

func NewUUIDGenerator(targetFile string) (UUIDGenerator, error) {
	jsonBytes, err := os.ReadFile(targetFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // if targetFile not exist, then no existing DetailedStep
		} else {
			return nil, fmt.Errorf("failed to read file %s, %s", targetFile, err)
		}
	}

	var detailedSteps []DetailedStep
	err = json.Unmarshal(jsonBytes, &detailedSteps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s, %s", targetFile, err)
	}

	return &UUIDGenFromTarget{existingSteps: detailedSteps}, nil
}

func (g *UUIDGenFromTarget) FindOrGenerateUUID(rs *RoughStep, subID string) string {
	for _, ds := range g.existingSteps {
		if ds.FromRoughStep && rs.Step == ds.ParentStep && subID == ds.SubID {
			return ds.Step
		}
	}

	// if not found, then new UUID
	return uuid.NewString()
}
