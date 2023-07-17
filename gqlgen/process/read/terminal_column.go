package read

import (
	"encoding/json"
	"fmt"
	"os"
)

type TerminalColumn struct {
	SeqNo  int    `json:"seqNo"`
	Column int    `json:"column"`
	Type   string `json:"type"`
	Text   string `json:"text"`
}

type TerminalColumns []TerminalColumn

func ReadTerminalColumns(filePath string) (TerminalColumns, error) {
	funcName := "ReadTerminalColumns"
	var entries TerminalColumns

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

func (t TerminalColumns) FindBySeqNo(seqNo int) *TerminalColumn {
	for _, e := range t {
		if e.SeqNo == seqNo {
			return &e // found!
		}
	}

	return nil
}
