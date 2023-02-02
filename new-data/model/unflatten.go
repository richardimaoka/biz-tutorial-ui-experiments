package model

import (
	"encoding/json"
	"strings"
)

func unflatten(bytes []byte) (map[string]interface{}, error) {
	var unmarshaled map[string]interface{}
	err := json.Unmarshal(bytes, &unmarshaled)
	if err != nil {
		panic(err)
	}

	return unflattenMap(unmarshaled)
}

func unflattenMap(m map[string]interface{}) (map[string]interface{}, error) {
	var nested = make(map[string]interface{})
	for k, v := range m {
		if dotIndex := strings.IndexRune(k, '.'); dotIndex == -1 {
			// no nesting
			nested[k] = v
		} else {
			// nesting needed
			parentKey := k[0:dotIndex]
			childKey := k[dotIndex+1:]

			if nested[parentKey] == nil {
				// if same parentKey not used
				children := make(map[string]interface{})
				children[childKey] = v
				nested[parentKey] = children
			} else {
				// if same paretKey already found
				children, ok := nested[parentKey].(map[string]interface{})
				if !ok {
					panic(parentKey + " conflicted!!")
				}
				children[childKey] = v
			}
		}
	}

	return nested, nil
}
