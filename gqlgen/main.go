package main

import (
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "processing" {
		if err := processing.EffectProcessing(); err != nil {
			log.Fatal(err)
		}
	} else {
		Server()
	}
}
