package rough_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/rough"
)

func TestDetailedStepStruct(t *testing.T) {
	// production DetailedStep struct
	ds := rough.DetailedStep{}
	fields1 := reflect.VisibleFields(reflect.TypeOf(ds))
	m1 := make(map[string]string)
	for _, field := range fields1 {
		m1[field.Name] = field.Type.String()
	}

	// test DetaieldStep struct
	dsTest := rough.DetailedStepTest{}
	fields2 := reflect.VisibleFields(reflect.TypeOf(dsTest))
	m2 := make(map[string]string)
	for _, field := range fields2 {
		m2[field.Name] = field.Type.String()
	}

	if diff := cmp.Diff(m1, m2); diff != "" {
		t.Fatalf("mismatch (-expected +result):\n%s", diff)
	}
}

func TestDetailedStepReadStruct(t *testing.T) {
	// production DetailedStep struct
	ds := rough.DetailedStep{}
	fields1 := reflect.VisibleFields(reflect.TypeOf(ds))
	m1 := make(map[string]string)
	for _, field := range fields1 {
		if field.Type.String() == "bool" {
			m1[field.Name] = "string"
		} else {
			m1[field.Name] = field.Type.String()
		}
	}

	// read DetaieldStep struct
	dsRead := rough.DetailedStepRead{}
	fields2 := reflect.VisibleFields(reflect.TypeOf(dsRead))
	m2 := make(map[string]string)
	for _, field := range fields2 {

		m2[field.Name] = field.Type.String()
	}

	if diff := cmp.Diff(m1, m2); diff != "" {
		t.Fatalf("mismatch (-expected +result):\n%s", diff)
	}
}
