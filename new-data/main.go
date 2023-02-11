package main

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/model"
)

func main() {
	// err := model.Process()
	// if err != nil {
	// 	panic(err)
	// }

	// actionListParse()

	err := model.SplitActionListFile("data2")
	if err != nil {
		panic(err)
	}

	return
}
