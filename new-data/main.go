package main

import (
	"encoding/json"
	"os"
)

func actionListParse() error {
	bytes, err := os.ReadFile("data2/action_list.json")
	if err != nil {
		return err
	}

	var unmarshaled []map[string]interface{}
	err = json.Unmarshal(bytes, &unmarshaled)
	if err != nil {
		return err
	}

	// m := make(map[string]interface{})
	// err = model.Unflatten(bytes, m)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func main() {
	// err := model.Process()
	// if err != nil {
	// 	panic(err)
	// }

	actionListParse()

	return
}
