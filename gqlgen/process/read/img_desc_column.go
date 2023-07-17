package read

import (
	"encoding/json"
	"fmt"
	"os"
)

type ImageDescriptionColumn struct {
	SeqNo                int    `json:"seqNo"`
	Column               int    `json:"column"`
	Width                int    `json:"width"`
	Height               int    `json:"height"`
	OriginalWidth        int    `json:"originalWidth"`
	OriginalHeight       int    `json:"originalHeight"`
	Path                 string `json:"path"`
	DescriptionContents  string `json:"description.contents"`
	DescriptionAlignment string `json:"description.alignment"`
}

type ImageDescriptionColumns []ImageDescriptionColumn

func ReadImageDescriptionColumns(filePath string) (ImageDescriptionColumns, error) {
	funcName := "ReadImageDescriptionColumns"
	var entries ImageDescriptionColumns

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

func (t ImageDescriptionColumns) FindBySeqNo(seqNo int) *ImageDescriptionColumn {
	for _, e := range t {
		if e.SeqNo == seqNo {
			return &e // found!
		}
	}

	return nil
}
