package model

import "encoding/json"

func unflatten(bytes []byte) map[string]interface{} {
	var unmarshaled map[string]interface{}
	err := json.Unmarshal(bytes, &unmarshaled)
	if err != nil {
		panic(err)
	}

	var nested = make(map[string]interface{})
	nested["parent"] = make(map[string]interface{})
	children, ok := nested["parent"].(map[string]interface{})
	if !ok {
		panic("chldren is not map[string]interface{P}")
	}

	children["childA"] = "AAA"

	children["childB"] = 10

	return nested
}
