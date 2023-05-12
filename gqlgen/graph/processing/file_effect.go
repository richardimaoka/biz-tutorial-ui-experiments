package processing

import (
	"encoding/json"
	"fmt"
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
	err := jsonRead("ReadFileEffects", filePath, unmarshaller)
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

//TODO: remove this function
func fileEffectsBySeqNo(seqNo int, effects []FileEffect) []FileEffect {
	var effectsBySeqNo []FileEffect
	for _, e := range effects {
		if e.SeqNo == seqNo {
			effectsBySeqNo = append(effectsBySeqNo, e)
		}
	}
	return effectsBySeqNo
}

func (f FileEffect) ToOperation() FileSystemOperation {
	switch f.OperationType {
	case "FileAdd":
		return FileAdd{FilePath: f.FilePath, Content: f.Content, IsFullContent: true}
	case "FileUpdate":
		return FileUpdate{FilePath: f.FilePath, Content: f.Content}
	case "FileDelete":
		return FileDelete{FilePath: f.FilePath}
	case "DirectoryAdd":
		return DirectoryAdd{FilePath: f.FilePath}
	case "DirectoryDelete":
		return DirectoryDelete{FilePath: f.FilePath}
	default:
		// this hsould never happen
		return nil
	}
}

//TODO: remove this function
func (f FileEffect) ToOperation2() (FileSystemOperation, error) {
	switch f.OperationType {
	case "FileAdd":
		return FileAdd{FilePath: f.FilePath, Content: f.Content, IsFullContent: true}, nil
	case "FileUpdate":
		return FileUpdate{FilePath: f.FilePath, Content: f.Content}, nil
	case "FileDelete":
		return FileDelete{FilePath: f.FilePath}, nil
	case "DirectoryAdd":
		return DirectoryAdd{FilePath: f.FilePath}, nil
	case "DirectoryDelete":
		return DirectoryDelete{FilePath: f.FilePath}, nil
	default:
		return nil, fmt.Errorf("FileEffect.ToOperation() found invalid OperationType = %s", f.OperationType)
	}
}
