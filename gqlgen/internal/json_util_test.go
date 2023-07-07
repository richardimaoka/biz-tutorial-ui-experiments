package internal_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

func TestJsonReadArray(t *testing.T) {
	arr, err := internal.JsonReadArray("testdata/array.json")
	if err != nil {
		t.Fatalf("JsonReadArray failed, %s", err)
	}
	expected := []internal.JsonObj{
		{"a": float64(1), "b": float64(2)},
		{"a": float64(3), "b": float64(4)},
		{"a": float64(5), "b": float64(6)},
	}
	diff := cmp.Diff(arr, expected)
	if diff != "" {
		t.Fatalf("mismatch (-expected +result):\n%s", diff)
	}
}
