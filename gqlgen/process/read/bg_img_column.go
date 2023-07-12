package read

import (
	"encoding/json"
	"fmt"
	"os"
)

type BackgroundImageColumn struct {
	SeqNo          int    `json:"seqNo"`
	Column         int    `json:"column"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	OriginalWidth  int    `json:"originalWidth"`
	OriginalHeight int    `json:"originalHeight"`
	Path           string `json:"path"`
	ModalText      string `json:"modal.text"`
	ModalPosition  string `json:"modal.position"`
}

type BackgroundImageColumns []BackgroundImageColumn

func ReadBackgroundImageColumns(filePath string) (BackgroundImageColumns, error) {
	funcName := "ReadBackgroundImageColumns"
	var entries BackgroundImageColumns

	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%s failed to read file = %s, %v", funcName, filePath, err)
	}

	json.Unmarshal(jsonBytes, &entries)
	if err != nil {
		return nil, fmt.Errorf("%s failed to unmarshal file = %s, %v", funcName, filePath, err)
	}

	return entries, err
}

func (t BackgroundImageColumns) FindBySeqNo(seqNo int) *BackgroundImageColumn {
	for _, e := range t {
		if e.SeqNo == seqNo {
			return &e // found!
		}
	}

	return nil
}
