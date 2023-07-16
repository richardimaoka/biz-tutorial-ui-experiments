package state

import (
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type FileNode interface {
	ToGraphQLFileNode() *model.FileNode
}

func FilePathInDir(parentDir, name string) string {
	if parentDir != "" {
		return parentDir + "/" + name
	} else {
		return name
	}
}
