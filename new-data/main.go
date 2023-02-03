package main

import "github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/model"

func actionListParse() error {
	err := model.SplitActionListFile("data2/action_list.json", "data2", "action")
	if err != nil {
		panic(err)
	}
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
