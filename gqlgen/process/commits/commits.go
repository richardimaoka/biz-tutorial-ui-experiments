package commits

import (
	"encoding/json"
	"fmt"
	"os"
)

type Commit struct {
	Commit  string `json:"commit"`
	Message string `json:"message"`
}

func Committtssss(tutorial string) error {
	filename := "data/" + tutorial + "/commits.json"

	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Errorf("Committtssss failed, %s", err)
	}

	var commits []Commit
	err = json.Unmarshal(bytes, &commits)
	if err != nil {
		fmt.Errorf("Committtssss failed, %s", err)
	}

	return nil
}
