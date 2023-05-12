package processing

import "encoding/json"

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

func fileEffectsBySeqNo(seqNo int, effects []FileEffect) []FileEffect {
	var effectsBySeqNo []FileEffect
	for _, e := range effects {
		if e.SeqNo == seqNo {
			effectsBySeqNo = append(effectsBySeqNo, e)
		}
	}
	return effectsBySeqNo
}
