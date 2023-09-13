package rough_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestRough(t *testing.T) {
	filename := "testdata/rough-step.json"
	goldenFile := "testdata/detailed-steps-golden.json"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	var roughStep rough.RoughStep
	err = json.Unmarshal(bytes, &roughStep)
	if err != nil {
		t.Fatalf("failed to unmarshal json: %v", err)
	}

	result := roughStep.Convert("c8238063-dd5a-4cd4-9d62-5c9c8ebd35ec", []string{})
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile, result)
}

// func TestRough1(t *testing.T) {
// 	cases := []struct {
// 		rough    rough.RoughStep
// 		expected []rough.DetailedStep
// 	}{
// 		{
// 			rough: rough.RoughStep{
// 				Instruction: "mkdir gqlgen-todos",
// 			},
// 			expected: []rough.DetailedStep{{
// 				FocusColumn:  "Terminal",
// 				TerminalText: "mkdir gqlgen-todos",
// 				TerminalType: "command",
// 			}},
// 		},
// 	}

// 	for _, c := range cases {
// 		steps := c.rough.Convert()
// 		if steps[0] != c.expected[0] {
// 			fmt.Printf("expected: %+v\n", c.expected[0])
// 			fmt.Printf("actual: %+v\n", steps[0])
// 			t.Fail()
// 		}
// 	}
// }
