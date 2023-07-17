package read

import (
	"encoding/json"
	"fmt"
	"os"
)

type MarkdownColumn struct {
	SeqNo                int    `json:"seqNo"`
	Column               int    `json:"column"`
	DescriptionContents  string `json:"description.contents"`
	DescriptionAlignment string `json:"description.alignment"`
}

type MarkdownColumns []MarkdownColumn

func ReadMarkdownColumns(filePath string) (MarkdownColumns, error) {
	funcName := "ReadMarkdownColumns"
	var entries MarkdownColumns

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

func (t MarkdownColumns) FindBySeqNo(seqNo int) *MarkdownColumn {
	for _, e := range t {
		if e.SeqNo == seqNo {
			return &e // found!
		}
	}

	return nil
}
