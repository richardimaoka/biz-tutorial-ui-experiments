package model

import (
	"encoding/json"
	"strings"
)

func unflatten(bytes []byte) map[string]interface{} {
	var unmarshaled map[string]interface{}
	err := json.Unmarshal(bytes, &unmarshaled)
	if err != nil {
		panic(err)
	}

	var nested = make(map[string]interface{})
	for k, v := range unmarshaled {
		if dotIndex := strings.IndexRune(k, '.'); dotIndex != -1 {
			parentKey := k[0:dotIndex]
			childKey := k[dotIndex+1:]

			if nested[parentKey] == nil {
				children := make(map[string]interface{})
				children[childKey] = v
				nested[parentKey] = children
			} else {
				children, ok := nested[parentKey].(map[string]interface{})
				if !ok {
					panic(parentKey + " conflicted!!")
				}
				children[childKey] = v
			}
		} else {
			nested[k] = v
		}
	}

	return nested
}
