package effect

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

type GitEffect struct {
	SeqNo          int    `json:"seqNo"`
	CommitHash     string `json:"commitHash"`
	PrevCommitHash string `json:"prevCommitHash"`
}

type GitEffects []GitEffect

func ReadGitEffects(filePath string) (GitEffects, error) {
	var effects GitEffects
	err := jsonwrap.Read(filePath, &effects)
	if err != nil {
		return nil, fmt.Errorf("ReadGitEffects failed to read file, %s", err)
	}

	return effects, err
}

func (f GitEffects) FilterBySeqNo(seqNo int) GitEffects {
	var effectsBySeqNo GitEffects
	for _, e := range f {
		if e.SeqNo == seqNo {
			effectsBySeqNo = append(effectsBySeqNo, e)
		}
	}
	return effectsBySeqNo
}

func (f GitEffects) FileEffects() FileEffects {
	return nil
}
