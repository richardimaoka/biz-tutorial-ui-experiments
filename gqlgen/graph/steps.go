package graph

import (
	"fmt"
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

type initStep struct {
	InitialStep string `json:"initialStep"`
}

func tutorialExists(tutorial string) bool {
	dirName := "data/" + tutorial
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return false
	}
	return true
}
func stepFile(tutorial string, step *string) (string, error) {
	dirName := fmt.Sprintf("data/%s/state", tutorial)

	if step == nil {
		log.Printf("no step is given, so reading initial step file")
		initStepFile := fmt.Sprintf("%s/_initial_step.json", dirName)

		var init initStep
		err := jsonwrap.Read(initStepFile, &init)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("%s/%s.json", dirName, init.InitialStep), nil
	}

	return fmt.Sprintf("%s/%s.json", dirName, *step), nil
}
