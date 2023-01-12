package main

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/pkg"
)

func main() {
	a, err := pkg.ReadActionFile("data/action01.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)

	err = pkg.SplitInputFile("data/step01/input.json")
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
