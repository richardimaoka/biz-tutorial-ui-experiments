package model2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnflatten(t *testing.T) {
	var result map[string]interface{}
	Unflatten([]byte(`{"parent.childA": "AAA", "parent.childB": 10, "parent.childD": null, "a": 250}`), &result)
	expected := map[string]interface{}{"parent": map[string]interface{}{"childA": "AAA", "childB": 10.0, "childD": nil}, "a": 250.0}

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Fatalf("mismatch (-expected +result):\n%s", diff)
	}
}
