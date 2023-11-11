package state2

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type Column interface {
	ToGraphQLColumnWrapper() *model.ColumnWrapper2
}
