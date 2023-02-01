package model

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnflatten(t *testing.T) {
	result := unflatten([]byte(`{"parent.childA": "AAA", "parent.childB": 10, "a": 250}`))
	expected := map[string]interface{}{"parent": map[string]interface{}{"childA": "AAA", "chilsdB": 10.0}, "a": 250}

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

				if reflect.ValueOf(childValue).Kind() == reflect.Ptr {
					t.Errorf("result[%s][%s] is a pointer %s of type %v", key, childKey, childValue, reflect.TypeOf(childValue))
					return
				}

				if reflect.ValueOf(eChildValue).Kind() == reflect.Ptr {
					t.Errorf("result[%s][%s] is a pointer %s of type %v", key, childKey, eChildValue, reflect.TypeOf(eChildValue))
					return
				}

				if childValue != eChildValue {
					t.Errorf("result[%s][%s] = %v is not equal to expected[%s][%s] = %v", key, childKey, childValue, key, childKey, eChildValue)
					return
				}
			}
		default:
			fmt.Println("result  : key =", key, "value = ", value)
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
