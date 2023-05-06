package processing_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

func TestTerminaEffect(t *testing.T) {
	bytes, err := os.ReadFile("testdata/test.json")
	if err != nil {
		t.Fatalf("reading file failed, %s", err)
	}

	var effects []processing.TerminalEffect
	err = json.Unmarshal(bytes, &effects)
	if err != nil {
		t.Fatalf("unmarshaling failed, %s", err)
	}

	jsonBytes, err := json.MarshalIndent(effects, "", "  ")
	if err != nil {
		t.Fatalf("marshaling failed, %s", err)
	}
	t.Errorf(string(jsonBytes))

	// cases := []struct {
	// 	ExpectedFile string
	// 	Step 1
	// }
}
