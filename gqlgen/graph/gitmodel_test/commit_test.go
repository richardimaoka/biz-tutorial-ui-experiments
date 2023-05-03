package gitmodel_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/gitmodel"
)

func TestActionCommandMarshal(t *testing.T) {
	f := &gitmodel.File{}
	f.Contents()
	t.Log(f.Contents())
}
