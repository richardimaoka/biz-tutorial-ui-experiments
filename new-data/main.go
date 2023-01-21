package main

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/pkg"
)

func main() {
	err := pkg.SplitInputFile("data/input01.json", "data/step01")
	if err != nil {
		panic(err)
	}
	// filename = "step01/result.json"
	// result, err := getResult(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)
}
