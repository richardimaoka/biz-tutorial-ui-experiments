package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type FileNode interface {
	ToGraphQLFileNode() *model.FileNode
	FilePath() string
}
