package internal_test

import (
	"encoding/json"
	"os"
	"strings"
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

func TestJsonReadArrayWrite(t *testing.T) {
	//non-floating point numbers are preserved as non-floating-point numbers
	arr, err := internal.JsonReadArray("testdata/array.json")
	if err != nil {
		t.Fatalf("JsonReadArray failed, %s", err)
	}

	result, err := json.Marshal(arr)
	if err != nil {
		t.Fatalf("json.Marshal failed, %s", err)
	}

	expectedBytes, err := os.ReadFile("testdata/array.json")
	if err != nil {
		t.Fatalf("os.ReadFile failed, %s", err)
	}
	expected := string(expectedBytes)
	expected = strings.ReplaceAll(expected, "\n", "")
	expected = strings.ReplaceAll(expected, " ", "")

	if string(result) != string(expected) {
		t.Fatalf("mismatch: expected = %s, result = %s", string(expected), string(result))
	}
}

func TestMarshalThenUnmarshal(t *testing.T) {
	obj := internal.JsonObj{"a": 10, "b": 20}
	var result struct {
		A int `json:"a"`
		B int `json:"b"`
	}
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &result) }

	err := internal.MarshalThenUnmarshal(obj, unmarshaller)
	if err != nil {
		t.Fatalf("MarshalThenUnmarshal failed, %s", err)
	}

	if result.A != 10 || result.B != 20 {
		t.Fatalf("mismatch: expected = %v, result = %v", obj, result)
	}
}
