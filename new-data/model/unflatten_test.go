package model

import (
	"fmt"
	"reflect"
	"testing"
)

func compareTwo(t *testing.T, v1name string, v1 interface{}, v2name string, v2 interface{}) {
	if reflect.ValueOf(v1).Kind() == reflect.Ptr {
		t.Errorf("%s is a pointer of type %v", v1name, reflect.TypeOf(v1))
		return
	}

	if reflect.ValueOf(v2).Kind() == reflect.Ptr {
		t.Errorf("%s is a pointer of type %v", v2name, reflect.TypeOf(v2))
		return
	}

	v1type := reflect.TypeOf(v1)
	v2type := reflect.TypeOf(v2)
	if v1type != v2type {
		t.Errorf("%s = %v has type = %v, not matching with %s = %v having type = %v", v1name, v1, v1type, v2name, v2, v2type)
		return
	}

	m1, isMap := v1.(map[string]interface{})
	if isMap {
		m2, ok := v2.(map[string]interface{})
		if !ok {
			t.Errorf("%s has type = %v, but must be map[string]interface{}", v2name, reflect.TypeOf(m2))
			return
		}
		compareTwoMaps(t, v1name, m1, v2name, m2)
	} else {
		if v1 != v2 {
			t.Errorf("%s = %v not equal to %s = %v", v1name, v1, v2name, v2)
		}
	}
}

func compareTwoMaps(t *testing.T, m1name string, m1 map[string]interface{}, m2name string, m2 map[string]interface{}) {
	kCompared := []string{}
	for k, v1 := range m1 {
		kCompared = append(kCompared, k)
		v2, ok := m2[k]
		if !ok {
			t.Errorf("%s[%s] exists but %s[%s] does not exist, ", m1name, k, m2name, k)
			continue
		}

		compareTwo(t, fmt.Sprintf("%s[%s]", m1name, k), v1, fmt.Sprintf("%s[%s]", m2name, k), v2)
	}

	for k, v2 := range m2 {
		// if already compared, skip k
		matched := -1
		for i, kc := range kCompared {
			if k == kc {
				matched = i
			}
		}
		if matched != -1 {
			continue
		}

		// not compared yet, then do compare
		v1, ok := m1[k]
		if !ok {
			t.Errorf("%s[%s] does not exist, while %s[%s] does", m1name, k, m2name, k)
			continue
		}

		compareTwo(t, fmt.Sprintf("%s[%s]", m1name, k), v1, fmt.Sprintf("%s[%s]", m2name, k), v2)
	}
}

func TestUnflatten(t *testing.T) {
	result := unflatten([]byte(`{"parent.childA": "AAA", "parent.childB": 10, "parent.childD": null, "a": 250, "b": "bbb", "d": 1234}`))
	expected := map[string]interface{}{"parent": map[string]interface{}{"childA": "AAASA", "childB": 10.0, "childC": nil}, "a": 250, "c": 2520}

	compareTwo(t, "expected", expected, "result", result)
}
