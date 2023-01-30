package main

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/new-data/pkg2"
)

func main() {
	err := pkg2.Process()
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
