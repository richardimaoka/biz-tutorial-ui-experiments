package processing_test

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/effect"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/preprocess/processing"
)

func Test_PageStateProcessor(t *testing.T) {
	dirName := "testdata/page_state/protoc-go-experiments"

	effects, err := effect.ConstructPageStateEffects(
		dirName+"/step-effects.json",
		dirName+"/file-effects.json",
		dirName+"/terminal-effects.json",
		dirName+"/markdown-effects.json",
	)
	if err != nil {
		t.Fatalf("ConstructPageStateEffects failed: %v", err)
	}

	stateDir := dirName + "/state"
	cases := []struct {
		ExpectedRegisteredFile   string
		ExpectedTransitionedFile string
		PageStateEffect          *effect.PageStateEffect
	}{
		{stateDir + "/state-000-register.json", stateDir + "/state-000-transition.json", effects[0]},
		{stateDir + "/state-001-register.json", stateDir + "/state-001-transition.json", effects[1]},
		{stateDir + "/state-002-register.json", stateDir + "/state-002-transition.json", effects[2]},
	}

	processor := processing.NewPageStateProcessor()
	for i, c := range cases {
		step := fmt.Sprintf("case[%d]", i)

		t.Run(step, func(t *testing.T) {
			op, err := c.PageStateEffect.ToOperation()
			if err != nil {
				t.Fatalf("ToOperation failed: %v", err)
			}

			if err := processor.RegisterNext(step, &op); err != nil {
				t.Fatalf("RegisterNext failed: %v", err)
			}
			internal.CompareWitGoldenFile(t, *updateFlag, c.ExpectedRegisteredFile, processor.ToGraphQLPageState())

			if err := processor.TransitionToNext(); err != nil {
				t.Fatalf("TransitionToNext failed: %v", err)
			}
			internal.CompareWitGoldenFile(t, *updateFlag, c.ExpectedTransitionedFile, processor.ToGraphQLPageState())
		})
	}
}
