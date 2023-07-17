package read

import (
	"encoding/json"
	"fmt"
	"os"
)

type GitColumn struct {
	SeqNo   int    `json:"seqNo"`
	Column  int    `json:"column"`
	Commit  string `json:"commit"`
	RepoUrl string `json:"repoUrl"`
}

type GitColumns []GitColumn

func ReadGitColumns(filePath string) (GitColumns, error) {
	funcName := "ReadGitColumns"
	var entries GitColumns

	jsonBytes, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("%s failed to read file = %s, %v", funcName, filePath, err)
	}

	json.Unmarshal(jsonBytes, &entries)
	if err != nil {
		return nil, fmt.Errorf("%s failed to unmarshal file = %s, %v", funcName, filePath, err)
	}

	return entries, err
}

func (t GitColumns) FindBySeqNo(seqNo int) *GitColumn {
	for _, e := range t {
		if e.SeqNo == seqNo {
			return &e // found!
		}
	}

	return nil
}
