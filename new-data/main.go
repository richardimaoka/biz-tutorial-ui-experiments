package main

import (
	"fmt"
	"log"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/model"
)

func main() {
	fmt.Println("running main")
	err := model.Processing()
	if err != nil {
		log.Fatal(err)
	}

	// actionListParse()

	// err := model.SplitActionListFile("data2")
	// if err != nil {
	// 	panic(err)
	// }

	return
}
