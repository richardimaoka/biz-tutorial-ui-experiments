package effect

import (
	"encoding/json"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
)

type StepEffect struct {
	SeqNo       int    `json:"seqNo"`
	CurrentStep string `json:"currentStep"`
	NextStep    string `json:"nextStep"`
}

type GiStepEffect struct {
	SeqNo       int    `json:"seqNo"`
	CurrentStep string `json:"currentStep"`
	NextStep    string `json:"nextStep"`
	CommitHash  string `json:"commitHash"`
}

type GeneralStepEffect struct {
	SeqNo       int    `json:"seqNo"`
	CurrentStep string `json:"currentStep"`
	NextStep    string `json:"nextStep"`
	CommitHash  string `json:"commitHash"`
}

func (s GeneralStepEffect) IsGitCommitStep() bool {
	return s.CommitHash != ""
}

func ReadStepEffects(filePath string) ([]StepEffect, error) {
	var effects []StepEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := internal.JsonRead(filePath, unmarshaller)
	if err != nil {
		return nil, fmt.Errorf("ReadStepEffects failed to read file, %s", err)
	}

	for i, effect := range effects {
		if effect.SeqNo != i {
			return nil, fmt.Errorf("ReadStepEffects failed to validate, effects[%d].SeqNo must be = %d, but %d", i, i, effect.SeqNo)
		}
	}

	return effects, err
}

func ReadGeneralStepEffects(filePath string) ([]GeneralStepEffect, error) {
	var effects []GeneralStepEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := internal.JsonRead(filePath, unmarshaller)
	if err != nil {
		return nil, fmt.Errorf("ReadGeneralStepEffects failed to read file, %s", err)
	}

	for i, effect := range effects {
		if effect.SeqNo != i {
			return nil, fmt.Errorf("ReadGeneralStepEffects failed to validate, effects[%d].SeqNo must be = %d, but %d", i, i, effect.SeqNo)
		}
	}

	return effects, err
}

// TODO: retain this as an alternative function returns []GeneralStepEffect
func GitStepEffects(repoUrl string) ([]GiStepEffect, error) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
	if err != nil {
		return nil, fmt.Errorf("GitStepEffects failed to initialize source code, cannot clone repo %s, %s", repoUrl, err)
	}

	// 1. collect commits in reverse order
	head, err := repo.Head()
	if err != nil {
		return nil, fmt.Errorf("error getting HEAD reference: %w", err)
	}

	latestCommit, err := repo.CommitObject(head.Hash())
	if err != nil {
		return nil, fmt.Errorf("error getting last commit: %w", err)
	}

	commitsInverseOrder := []*object.Commit{latestCommit}
	commit := latestCommit
	for {
		if commit.NumParents() > 1 {
			return nil, fmt.Errorf("commit history branched, which is not supported")
		}
		parentCommit, err := commit.Parent(0)
		if err == object.ErrParentNotFound {
			break // reached the first commit
		} else if err != nil {
			return nil, fmt.Errorf("error getting parent of commit %s: %w", commit.Hash, err)
		}

		commit = parentCommit
		commitsInverseOrder = append(commitsInverseOrder, commit)
	}

	// 2. convert commits to effects
	var effects []GiStepEffect
	seqNo := 0
	for i := len(commitsInverseOrder) - 1; i >= 0; i-- {
		currentStep := fmt.Sprintf("%03d", seqNo)
		nextStep := fmt.Sprintf("%03d", seqNo+1)
		effect := GiStepEffect{seqNo, currentStep, nextStep, commitsInverseOrder[i].Hash.String()}
		effects = append(effects, effect)
		seqNo++
	}

	return effects, nil
}
