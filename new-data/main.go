package main

import "github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/model"

func actionListParse() error {
	model.SplitActionListFile("data2/action_list.json", "data2", "action")
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
