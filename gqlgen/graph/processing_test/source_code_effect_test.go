package processing_test

// import (
// 	"encoding/json"
// 	"os"
// 	"testing"

// 	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
// )

// func Test_SourceCodeGitEffect(t *testing.T) {
// 	jsonBytes, err := os.ReadFile("testdata/source_code_effect/1.json")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	var gitEffect processing.SourceCodeGitEffect
// 	err = json.Unmarshal(jsonBytes, &gitEffect)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	jsonRemarshaled, err := json.Marshal(gitEffect)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Errorf("%s", string(jsonRemarshaled))
// }
