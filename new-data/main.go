package main

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/model"
)

func main() {
	// err := model.Process()
	// if err != nil {
	// 	panic(err)
	// }

	// actionListParse()

	err := model.SplitInputListFile("data2")
	if err != nil {
		panic(err)
	}

	list, err := model.ListInputFiles("data2")
	if err != nil {
		panic(err)
	}
	fmt.Println(list)

	return
}
