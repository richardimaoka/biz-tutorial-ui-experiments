package effect

import (
	"encoding/json"
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type MarkdownEffect struct {
	SeqNo    int    `json:"seqNo"`
	Markdown string `json:"markdown"`
}

type MarkdownEffects []MarkdownEffect

func (t MarkdownEffects) FindBySeqNo(seqNo int) *MarkdownEffect {
	for _, e := range t {
		if e.SeqNo == seqNo {
			return &e // found!
		}
	}

	return nil
}

func ReadMarkdownEffects(filePath string) (MarkdownEffects, error) {
	var effects MarkdownEffects
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := internal.JsonRead(filePath, unmarshaller)
	if err != nil {
		return nil, fmt.Errorf("ReadMarkdownEffects failed to read file, %s", err)
	}

	return effects, err
}
