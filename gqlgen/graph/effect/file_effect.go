package effect

import (
	"encoding/json"
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/processing"
)

type FileEffect struct {
	SeqNo         int    `json:"seqNo"`
	OperationType string `json:"operationType"`
	FilePath      string `json:"filePath"`
	Content       string `json:"content"`
}

type FileEffects []FileEffect

func ReadFileEffects(filePath string) (FileEffects, error) {
	var effects []FileEffect
	unmarshaller := func(jsonBytes []byte) error { return json.Unmarshal(jsonBytes, &effects) }
	err := jsonRead(filePath, unmarshaller)
	if err != nil {
		return nil, fmt.Errorf("ReadFileEffects failed to read file, %s", err)
	}

	return effects, err
}

func (f FileEffects) FilterBySeqNo(seqNo int) FileEffects {
	var effectsBySeqNo []FileEffect
	for _, e := range f {
		if e.SeqNo == seqNo {
			effectsBySeqNo = append(effectsBySeqNo, e)
		}
	}
	return effectsBySeqNo
}

func (f FileEffects) ToOperation() ([]processing.FileOperation, error) {
	ops := []processing.FileOperation{}
	for i, e := range f {
		op, err := e.ToOperation()
		if err != nil {
			// this should never happen
			return nil, fmt.Errorf("Failed in ToOperation() in FileEffects[%d]: %v", i, err)
		}

		ops = append(ops, op)
	}

	return ops, nil
}

func (f FileEffect) ToOperation() (processing.FileOperation, error) {
	switch f.OperationType {
	case "FileAdd":
		return processing.FileAdd{FilePath: f.FilePath, Content: f.Content, IsFullContent: true}, nil
	case "FileUpdate":
		return processing.FileUpdate{FilePath: f.FilePath, Content: f.Content}, nil
	case "FileDelete":
		return processing.FileDelete{FilePath: f.FilePath}, nil
	case "DirectoryAdd":
		return processing.DirectoryAdd{FilePath: f.FilePath}, nil
	case "DirectoryDelete":
		return processing.DirectoryDelete{FilePath: f.FilePath}, nil
	case "FileUpsert":
		return processing.FileUpsert{FilePath: f.FilePath, Content: f.Content, IsFullContent: true}, nil
	default:
		// this should never happen
		return nil, fmt.Errorf("FileEffect.ToOperation failed, wrong operation type = %s", f.OperationType)
	}
}
