package model

import (
	"fmt"
	"reflect"
	"testing"
)

func compareTwoValues(t *testing.T, v1name string, v1 interface{}, v2name string, v2 interface{}) {
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

	if v1 != v2 {
		t.Errorf("%s = %v is not equal to %s = %v", v1name, v1, v2name, v2)
		return
	}
}

func compareTwoMaps(t *testing.T, m1name string, m1 interface{}, m2name string, m2 interface{}) {
	m1Map, ok := m1.(map[string]interface{})
	if !ok {
		t.Errorf("%s has type = %v, not map[string]interface{}", reflect.TypeOf(m1), m1name)
	}

	m2Map, ok := m2.(map[string]interface{})
	if !ok {
		t.Errorf("%s has type = %v, not map[string]interface{}", reflect.TypeOf(m2), m2name)
	}

	for k, v1 := range m1Map {
		v2, ok := m2Map[k]
		if !ok {
			t.Errorf("%s[%s] does not exist, while %s[%s] does", m2name, k, m1name, k)
			return
		}

		compareTwoValues(t, fmt.Sprintf("%s[%s]", m1name, k), v1, fmt.Sprintf("%s[%s]", m2name, k), v2)
	}

	for k, v2 := range m2Map {
		v1, ok := m1Map[k]
		if !ok {
			t.Errorf("%s[%s] does not exist, while %s[%s] does", m1name, k, m2name, k)
			return
		}

		compareTwoValues(t, fmt.Sprintf("%s[%s]", m1name, k), v1, fmt.Sprintf("%s[%s]", m2name, k), v2)
	}
}

func TestUnflatten(t *testing.T) {
	result := unflatten([]byte(`{"parent.childA": "AAA", "parent.childB": 10, "parent.childD": null, "a": 250, "b": "bbb", "d": 1234}`))
	expected := map[string]interface{}{"parent": map[string]interface{}{"childA": "AAASA", "childB": 10.0, "childC": nil}, "a": 250, "c": 2520}

	for key, value := range result {
		switch children := value.(type) {
		case map[string]interface{}:
			t.Run("a", func(t *testing.T) {
				compareTwoMaps(t, fmt.Sprintf("result[%s]", key), children, fmt.Sprintf("expected[%s]", key), expected[key])
			})
		default:
			eValue, ok := expected[key]
			if !ok {
				t.Errorf("expected[%s] does not exist, while result[%s] does", key, key)
				continue
			}
			t.Run("b", func(t *testing.T) {
				compareTwoValues(t, fmt.Sprintf("result[%s]", key), value, fmt.Sprintf("expected[%s]", key), eValue)
			})
		}
	}
}
