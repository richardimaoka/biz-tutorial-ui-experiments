package main

import (
	"encoding/json"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/model"
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

	model.Unflatten(bytes)

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
