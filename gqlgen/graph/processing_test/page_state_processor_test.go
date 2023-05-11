package processing

import (
	"fmt"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

func Test_PageState(t *testing.T) {
	actions, err := processing.ReadActionFromFiles("testdata/page_state/enriched", "action")
	if err != nil {
		t.Fatalf("reading actions failed, %s", err)
	}

	pageState, err := processing.InitPageStateProcessor(actions[0])
	if err != nil {
		t.Fatalf("init page state failed, %s", err)
	}
	internal.CompareAfterMarshal(t, "testdata/page_state/state/state-000.json", pageState.ToGraphQLPageState())

	for i := 1; i <= 8; i++ {
		if err := pageState.StateTransition(actions[i]); err != nil {
			t.Fatalf("state transition failed, %s", err)
		}
		fileName := fmt.Sprintf("testdata/page_state/state/state-%03d.json", i)
		internal.CompareAfterMarshal(t, fileName, pageState.ToGraphQLPageState())
	}
}
