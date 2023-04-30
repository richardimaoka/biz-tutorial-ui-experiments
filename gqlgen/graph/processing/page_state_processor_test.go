package processing

import "testing"

func Test_PageState(t *testing.T) {
	firstAction, err := readAction("testdata/page_state/enriched/action000.json")
	if err != nil {
		t.Fatalf("reading first action failed, %s", err)
	}

	pageState, err := InitPageStateProcessor(firstAction)
	if err != nil {
		t.Fatalf("init page state failed, %s", err)
	}

	compareAfterMarshal(t, "testdata/page_state/state/state-000.json", pageState.ToGraphQLPageState())
}
