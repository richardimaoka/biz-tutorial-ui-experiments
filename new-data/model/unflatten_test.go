package model

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUnflatten(t *testing.T) {
	result := unflatten([]byte(`{"parent.childA": "AAA", "parent.childB": 10}`))
	expected := map[string]interface{}{"parent.childA": "AAA", "parent.childB": 10}

	fmt.Println("type of result:", reflect.TypeOf(result))
	for key, value := range result {
		fmt.Println("result  : key =", key, ", value =", value)
		children, ok := value.(map[string]interface{})
		if !ok {
		}
		for childKey, childValue := range children {
			fmt.Println("result  : key =", key, ", childKey =", childKey, ", childValue = ", childValue)
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
