package model

import (
	"fmt"
	"testing"
)

func TestUnflatten(t *testing.T) {
	result := unflatten([]byte(`{"parent.childA": "AAA", "parent.childB": 10, "a": 250}`))
	expected := map[string]interface{}{"parent": map[string]interface{}{"childA": "AAA", "childB": 10}, "a": 250}

	for key, value := range result {
		switch t := value.(type) {
		case map[string]interface{}:
			for childKey, childValue := range t {
				fmt.Println("result  : key =", key, ", childKey =", childKey, ", childValue = ", childValue)
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
