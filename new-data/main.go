package main

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/model"
)

func main() {

	err := model.Processing()
	if err != nil {
		panic(err)
	}

	// actionListParse()

	// err := model.SplitActionListFile("data2")
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("main")
	return
}
