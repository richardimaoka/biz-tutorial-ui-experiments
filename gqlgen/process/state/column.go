package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type Column interface {
	ToGraphQLColumnWrapper() *model.ColumnWrapper2
}
