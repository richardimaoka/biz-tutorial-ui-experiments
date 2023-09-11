package commits

import (
	"encoding/json"
	"fmt"
	"os"
)

type Commit struct {
	Commit  string `json:"commit"`
	Message string `json:"message"`
}

func Committtssss(tutorial string) error {
	filename := "data/" + tutorial + "/commits.json"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Committtssss failed, %s", err)
	}

	var commits []Commit
	err = json.Unmarshal(bytes, &commits)
	if err != nil {
		return fmt.Errorf("Committtssss failed, %s", err)
	}

	for _, commit := range commits {
		fmt.Printf("%s: %s\n", commit.Commit, commit.Message)
	}

	return nil
}

type RoughStep struct {
	Commit      string `json:"commit"`
	Instruction string `json:"instruction"`
}

func findMatchingRoughStep(commit *Commit, roughSteps []RoughStep) *RoughStep {
	for _, roughStep := range roughSteps {
		if commit.Message == roughStep.Instruction {
			return &roughStep
		}
	}
	return nil
}

func Reconcile(commits []Commit, roughSteps []RoughStep) ([]*RoughStep, error) {
	var updatedRoughSteps []*RoughStep

	for _, commit := range commits {
		roughStep := findMatchingRoughStep(&commit, roughSteps)
		if roughStep != nil {
			return nil, fmt.Errorf("rough step for %v not found", commit)
		}

		roughStep.Commit = commit.Commit
		updatedRoughSteps = append(updatedRoughSteps, roughStep)
	}

	return updatedRoughSteps, nil
}
