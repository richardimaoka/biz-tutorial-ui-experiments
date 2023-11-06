package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

type StepEffect struct {
	SeqNo      int    `json:"seqNo"`
	Step       string `json:"step"`
	CommitHash string `json:"commitHash"`
}

func (s StepEffect) IsGitCommitStep() bool {
	return s.CommitHash != ""
}

func ReadStepEffects(filePath string) ([]StepEffect, error) {
	funcName := "ReadStepEffects"
	var effects []StepEffect
	err := jsonwrap.JsonRead(filePath, &effects)
	if err != nil {
		return nil, fmt.Errorf("%s failed to read file, %s", funcName, err)
	}

	for i, effect := range effects {
		if effect.SeqNo != i {
			return nil, fmt.Errorf("%s failed to validate, effects[%d].SeqNo must be = %d, but %d", funcName, i, i, effect.SeqNo)
		}
	}

	return effects, err
}

// func GitStepEffects(repoUrl string) ([]GiStepEffect, error) {
// 	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: repoUrl})
// 	if err != nil {
// 		return nil, fmt.Errorf("GitStepEffects failed to initialize source code, cannot clone repo %s, %s", repoUrl, err)
// 	}

// 	// 1. collect commits in reverse order
// 	head, err := repo.Head()
// 	if err != nil {
// 		return nil, fmt.Errorf("error getting HEAD reference: %w", err)
// 	}

// 	latestCommit, err := repo.CommitObject(head.Hash())
// 	if err != nil {
// 		return nil, fmt.Errorf("error getting last commit: %w", err)
// 	}

// 	commitsInverseOrder := []*object.Commit{latestCommit}
// 	commit := latestCommit
// 	for {
// 		if commit.NumParents() > 1 {
// 			return nil, fmt.Errorf("commit history branched, which is not supported")
// 		}
// 		parentCommit, err := commit.Parent(0)
// 		if err == object.ErrParentNotFound {
// 			break // reached the first commit
// 		} else if err != nil {
// 			return nil, fmt.Errorf("error getting parent of commit %s: %w", commit.Hash, err)
// 		}

// 		commit = parentCommit
// 		commitsInverseOrder = append(commitsInverseOrder, commit)
// 	}

// 	// 2. convert commits to effects
// 	var effects []GiStepEffect
// 	seqNo := 0
// 	for i := len(commitsInverseOrder) - 1; i >= 0; i-- {
// 		currentStep := fmt.Sprintf("%03d", seqNo)
// 		nextStep := fmt.Sprintf("%03d", seqNo+1)
// 		effect := GiStepEffect{seqNo, currentStep, nextStep, commitsInverseOrder[i].Hash.String()}
// 		effects = append(effects, effect)
// 		seqNo++
// 	}

// 	return effects, nil
// }
