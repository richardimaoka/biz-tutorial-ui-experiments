package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type Slide interface {
	ToGraphQLSlideWrapper() *model.SlideWrapper
}
