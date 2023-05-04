package gitmodel_test

import (
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/gitmodel"
)

func TestStepFromGit(t *testing.T) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "https://github.com/richardimaoka/gqlgensandbox",
	})
	if err != nil {
		t.Fatalf("error cloning repo: %v", err)
	}

	type Expected struct {
		PrevStep    string
		CurrentStep string
		NextStep    string
	}

	expectedValues := []Expected{
		{
			CurrentStep: "91a99d0c0558d2fc03c930d19afa97fc141f0c2e",
			PrevStep:    "",
			NextStep:    "83f8ad84dea56e8e5549832fb98eb8b5b9db4912",
		},
		{
			CurrentStep: "83f8ad84dea56e8e5549832fb98eb8b5b9db4912",
			PrevStep:    "91a99d0c0558d2fc03c930d19afa97fc141f0c2e",
			NextStep:    "86a03f4f18b081b07e058f0e9f96503772a50cf0",
		},
		{
			CurrentStep: "86a03f4f18b081b07e058f0e9f96503772a50cf0",
			PrevStep:    "83f8ad84dea56e8e5549832fb98eb8b5b9db4912",
			NextStep:    "490808086bded6b27f3651b095aefb7bb6708da2",
		},
		{
			CurrentStep: "490808086bded6b27f3651b095aefb7bb6708da2",
			PrevStep:    "86a03f4f18b081b07e058f0e9f96503772a50cf0",
			NextStep:    "9f835b8aaafdfc55933f52aae5a7c6e9864432aa",
		},
		{
			CurrentStep: "9f835b8aaafdfc55933f52aae5a7c6e9864432aa",
			PrevStep:    "490808086bded6b27f3651b095aefb7bb6708da2",
			NextStep:    "99a2c7f129cbebab3b17034fa93ad579d0fe29f6",
		},
		{
			CurrentStep: "99a2c7f129cbebab3b17034fa93ad579d0fe29f6",
			PrevStep:    "9f835b8aaafdfc55933f52aae5a7c6e9864432aa",
			NextStep:    "20c5ef14fc6a0deae8a528beee3ed0f984da9ae1",
		},
		{
			CurrentStep: "20c5ef14fc6a0deae8a528beee3ed0f984da9ae1",
			PrevStep:    "99a2c7f129cbebab3b17034fa93ad579d0fe29f6",
			NextStep:    "4bc48072066d6e9fe339fae1341c196d4be03286",
		},
		{
			CurrentStep: "4bc48072066d6e9fe339fae1341c196d4be03286",
			PrevStep:    "20c5ef14fc6a0deae8a528beee3ed0f984da9ae1",
			NextStep:    "8d08178cb98df959288f2c4f8d0aff1bb20d6fc9",
		},
		{
			CurrentStep: "8d08178cb98df959288f2c4f8d0aff1bb20d6fc9",
			PrevStep:    "4bc48072066d6e9fe339fae1341c196d4be03286",
			NextStep:    "813c7822a3232c43edd9cc02286f063851ff2b54",
		},
		{
			CurrentStep: "813c7822a3232c43edd9cc02286f063851ff2b54",
			PrevStep:    "8d08178cb98df959288f2c4f8d0aff1bb20d6fc9",
			NextStep:    "a234864d58a12d50458ee563ba59c628c6861286",
		},
		{
			CurrentStep: "a234864d58a12d50458ee563ba59c628c6861286",
			PrevStep:    "813c7822a3232c43edd9cc02286f063851ff2b54",
			NextStep:    "18c23ac5d49428845afe91f2d189968265afc19f",
		},
		{
			CurrentStep: "18c23ac5d49428845afe91f2d189968265afc19f",
			PrevStep:    "a234864d58a12d50458ee563ba59c628c6861286",
			NextStep:    "e02dc3bbdf21a533f1812507134cf1484a971f5b",
		},
		{
			CurrentStep: "e02dc3bbdf21a533f1812507134cf1484a971f5b",
			PrevStep:    "18c23ac5d49428845afe91f2d189968265afc19f",
			NextStep:    "929e04606a6eb7619f0e0949c2bdc2a1688a2d25",
		},
		{
			CurrentStep: "929e04606a6eb7619f0e0949c2bdc2a1688a2d25",
			PrevStep:    "e02dc3bbdf21a533f1812507134cf1484a971f5b",
			NextStep:    "b08a390257a68951b2cf05a463655c852de7a4de",
		},
		{
			CurrentStep: "b08a390257a68951b2cf05a463655c852de7a4de",
			PrevStep:    "929e04606a6eb7619f0e0949c2bdc2a1688a2d25",
			NextStep:    "f745b8e233b2adfd11a63e7855f18a1986c7c084",
		},
		{
			CurrentStep: "f745b8e233b2adfd11a63e7855f18a1986c7c084",
			PrevStep:    "b08a390257a68951b2cf05a463655c852de7a4de",
			NextStep:    "700a1d749f1d3e86330ebe163d64a9fe58e25fd2",
		},
		{
			CurrentStep: "700a1d749f1d3e86330ebe163d64a9fe58e25fd2",
			PrevStep:    "f745b8e233b2adfd11a63e7855f18a1986c7c084",
			NextStep:    "8c62836cfbbf9a9d0ce957d0edc4084e4bc88e60",
		},
		{
			CurrentStep: "8c62836cfbbf9a9d0ce957d0edc4084e4bc88e60",
			PrevStep:    "700a1d749f1d3e86330ebe163d64a9fe58e25fd2",
			NextStep:    "4dd8f51d6acbee9d61b24dc26715ecc48a5d2456",
		},
		{
			CurrentStep: "4dd8f51d6acbee9d61b24dc26715ecc48a5d2456",
			PrevStep:    "8c62836cfbbf9a9d0ce957d0edc4084e4bc88e60",
			NextStep:    "",
		},
	}

	step, err := gitmodel.FirstStepFromGit(repo)
	if err != nil {
		t.Fatalf("unexpected FirstStepFromGit failure: %v", err)
	}

	for _, e := range expectedValues {
		t.Run(e.CurrentStep, func(t *testing.T) {
			if step.CurrenStep() != e.CurrentStep {
				t.Errorf("calculated CurrentStep() = %v, want %v", step.CurrenStep(), e.CurrentStep)
			}
			if step.PrevStep() != e.PrevStep {
				t.Errorf("calculated PrevStep() = %v, want %v", step.PrevStep(), e.PrevStep)
			}
			if step.NextStep() != e.NextStep {
				t.Errorf("calculated NextStep() = %v, want %v", step.NextStep(), e.NextStep)
			}

			if err := step.Increment(); err != nil && e.NextStep != "" {
				t.Fatalf("unexpected step.Increment failure at %+v: %v", step, err)
			}
		})
	}
}
