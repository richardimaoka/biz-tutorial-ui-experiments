package gitmodel

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type StepFromGit struct {
	commits      []*object.Commit
	currentIndex int
}

func FirstStepFromGit(repo *git.Repository) (*StepFromGit, error) {
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

	commits := []*object.Commit{}
	for i := len(commitsInverseOrder) - 1; i >= 0; i-- {
		commits = append(commits, commitsInverseOrder[i])
	}

	if len(commits) == 0 {
		return nil, fmt.Errorf("no commits found")
	} else if len(commits) == 1 {
		return &StepFromGit{
			commits:      commits,
			currentIndex: 0,
		}, nil
	} else {
		return &StepFromGit{
			commits:      commits,
			currentIndex: 0,
		}, nil
	}
}

func (s *StepFromGit) Increment() error {
	if s.currentIndex == len(s.commits)-1 {
		return fmt.Errorf("already at the last step")
	} else {
		s.currentIndex++
		return nil
	}
}

func (s *StepFromGit) CurrenStep() string {
	return s.commits[s.currentIndex].Hash.String()
}

func (s *StepFromGit) NextStep() string {
	if s.currentIndex == len(s.commits)-1 {
		return ""
	} else {
		return s.commits[s.currentIndex+1].Hash.String()
	}
}

func (s *StepFromGit) PrevStep() string {
	if s.currentIndex == 0 {
		return ""
	} else {
		return s.commits[s.currentIndex-1].Hash.String()
	}
}
