package main

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/pkg"
)

func main() {
	filename := "data/step01/action.json"
	action, err := pkg.GetActionFromFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(action)

	a, err := pkg.ReadActionFile("data/action01.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	// filename = "step01/result.json"
	// result, err := getResult(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)
}
