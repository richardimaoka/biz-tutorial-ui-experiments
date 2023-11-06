package rough

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type UUIDFinder struct {
	targetSteps   []DetailedStep
	uuidGenerator func() string
}

func NewUUIDFinder(targetFile string) (*UUIDFinder, error) {
	detailedSteps, err := readExistingSteps(targetFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read existing steps: %s", err)
	}

	return &UUIDFinder{
		targetSteps:   detailedSteps,
		uuidGenerator: generateUUID,
	}, nil
}

func (g *UUIDFinder) FindOrGenerateUUID(rs *ResultStep, subID string) string {
	for _, ds := range g.targetSteps {
		if ds.FromRoughStep && rs.Step == ds.ParentStep && subID == ds.SubID {
			return ds.Step
		}
	}

	// if not found, then new UUID
	return g.uuidGenerator()
}

func readExistingSteps(targetFile string) ([]DetailedStep, error) {
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

	return detailedSteps, nil
}

func generateUUID() string {
	return uuid.NewString()
}
