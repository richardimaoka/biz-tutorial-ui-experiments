package model

import (
	"fmt"
	"testing"
)

func TestUnflatten(t *testing.T) {
	result := unflatten([]byte(`{"parent.childA": "AAA", "parent.childB": 10, "a": 250}`))
	expected := map[string]interface{}{"parent.childA": "AAA", "parent.childB": 10}

	for key, value := range result {
		switch t := value.(type) {
		case map[string]interface{}:
			for childKey, childValue := range t {
				fmt.Println("result  : key =", key, ", childKey =", childKey, ", childValue = ", childValue)
			}
		default:
			fmt.Println("result  : key = ", key, "value = ", value)
		}
	}
	for key, value := range expected {
		fmt.Println("expected: key =", key, ", value =", value)
	}
	// if  == expected["parent"]["childA"] {
	// }
	// for key, v := range expected {
	// 	if v != expected[key] {
	// 		t.Errorf("expected %q but got %q", expected, result)
	// 	}
	// }
}
