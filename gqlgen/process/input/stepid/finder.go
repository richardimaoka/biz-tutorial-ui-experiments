package stepid

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/input"
)

type Finder struct {
	// needs to hold the target steps in memory, to avoid error in the signature of all methods
	targetSteps []input.ResultStep
	idGenerator func() string
}

// Create a new UUID finder instance, to reconcile step ID against the already-generated target file
//   targetFile: File name of the target file to find UUIDs from.
//               If targetFile doesn't exist, then the finder will always generate a brand new UUID.
func NewFinder(targetFile string) (*Finder, error) {
	steps, err := readExistingSteps(targetFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read existing steps: %s", err)
	}

	return &Finder{
		targetSteps: steps,
		idGenerator: generateUUID,
	}, nil
}

func (g *Finder) StepIdFor(parentStep, subID string) string {
	for _, ds := range g.targetSteps {
		if ds.IsFromRow && parentStep == ds.ParentStep && subID == ds.SubID {
			return ds.Step
		}
	}

	// if not found, then new UUID
	return g.idGenerator()
}

// if no existing steps, then return nil = empty slice
func readExistingSteps(targetFile string) ([]input.ResultStep, error) {
	jsonBytes, err := os.ReadFile(targetFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // if targetFile not exist, then no existing DetailedStep
		} else {
			return nil, fmt.Errorf("failed to read file %s, %s", targetFile, err)
		}
	}

	var steps []input.ResultStep
	err = json.Unmarshal(jsonBytes, &steps)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s, %s", targetFile, err)
	}

	return steps, nil
}

func generateUUID() string {
	return uuid.NewString()
}
