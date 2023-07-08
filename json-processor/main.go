package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/json-processor/effect"
)

const dir_to_scan string = "images"

func imageProcessing() {
	files, _ := ioutil.ReadDir(dir_to_scan)
	for _, imgFile := range files {

		if reader, err := os.Open(filepath.Join(dir_to_scan, imgFile.Name())); err == nil {
			defer reader.Close()
			im, format, err := image.Decode(reader)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s, %v\n", imgFile.Name(), format, err)
				continue
			}
			fmt.Printf("%s %d %d\n", imgFile.Name(), im.Bounds().Dx(), im.Bounds().Dy())
		} else {
			fmt.Println("Impossible to open the file:", err)
		}
	}
}

func main() {
	file := "data/img_columns.json"
	objs, err := effect.ReadImageColumnEffects(file)
	if err != nil {
		panic(err)
	}
	for _, obj := range objs {
		fmt.Println(obj)
	}
}
