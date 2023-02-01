package model

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnflatten(t *testing.T) {
	result := unflatten([]byte(`{"parent.childA": "AAA", "parent.childB": 10, "a": 250}`))
	expected := map[string]interface{}{"parent": map[string]interface{}{"childA": "AAA", "childB": 10.0}, "a": 250}

	for key, value := range result {
		switch children := value.(type) {
		case map[string]interface{}:
			for childKey, childValue := range children {
				fmt.Println("result  : key =", key, ", childKey =", childKey, ", childValue = ", childValue)

				if reflect.ValueOf(childValue).Kind() == reflect.Ptr {
					panic("child value pointer!")
				}

				e, ok := expected[key]
				if !ok {
					panic("not okkk!!")
				}

				ec, ok := e.(map[string]interface{})
				if !ok {
					panic("not okkk2!!")
				}

				if childValue != ec[childKey] {
					fmt.Println(reflect.TypeOf(childValue))
					fmt.Println(childValue)
					fmt.Println(reflect.TypeOf(ec[childKey]))
					fmt.Println(ec[childKey])
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
