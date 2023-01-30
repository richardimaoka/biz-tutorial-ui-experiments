package main

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/pkg2"
)

func main() {
	err := pkg2.Process("data2/action00.json", "data2/step00.json")
	if err != nil {
		panic(err)
	}
	// err := pkg.ConstructState("data/step01", "data/step02")
	// if err != nil {
	// 	panic(err)
	// }
	// filename = "step01/result.json"
	// result, err := getResult(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)
}
