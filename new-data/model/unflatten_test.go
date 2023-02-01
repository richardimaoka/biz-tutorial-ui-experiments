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
		t.Errorf("%s has type = %v, not matching with %s in type = %v", v1name, v1type, v2name, v2type)
		return
	}

	if v1 != v2 {
		t.Errorf("%s = %v is not equal to %s = %v", v1name, v1, v2name, v2)
		return
	}
}

func TestUnflatten(t *testing.T) {
	result := unflatten([]byte(`{"parent.childA": "AAA", "parent.childB": 10, "a": 250, "b": "bbb"}`))
	expected := map[string]interface{}{"parent": map[string]interface{}{"childA": "AAA", "childB": 10.0}, "a": 250, "c": 2520}

	for key, value := range result {
		switch children := value.(type) {
		case map[string]interface{}:
			for childKey, childValue := range children {
				fmt.Println("result  : key =", key, ", childKey =", childKey, ", childValue = ", childValue)

				eChildren, ok := expected[key]
				if !ok {
					t.Errorf("expected[%s] does not exist, while result[%s] does", key, key)
					return
				}

				ecMap, ok := eChildren.(map[string]interface{})
				if !ok {
					t.Errorf("expected[%s] is not a map[string]interface{}, while result[%s] is", key, key)
					return
				}

				eChildValue, ok := ecMap[childKey]
				if !ok {
					t.Errorf("expected[%s][%s] does not exist, while result[%s][%s] does", key, childKey, key, childKey)
					return
				}

				compareTwoValues(t, fmt.Sprintf("result[%s][%s]", key, childKey), childValue, fmt.Sprintf("expected[%s][%s]", key, childKey), eChildValue)
			}
		default:
			eValue, ok := expected[key]
			if !ok {
				t.Errorf("expected[%s] does not exist, while result[%s] does", key, key)
				return
			}

			compareTwoValues(t, fmt.Sprintf("result[%s]", key), value, fmt.Sprintf("expected[%s]", key), eValue)
		}
	}

	for key, value := range expected {
		switch t := value.(type) {
		case map[string]interface{}:
			for childKey, childValue := range t {
				fmt.Println("expected: key =", key, ", childKey =", childKey, ", childValue = ", childValue)
			}
		default:
			fmt.Println("expected: key =", key, "value = ", value)
		}
	}
}
