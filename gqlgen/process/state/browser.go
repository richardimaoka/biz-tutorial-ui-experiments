package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type Browser struct {
	image *Image
}

func NewBrowser(tutorial, src string, width, height int) *Browser {
	return &Browser{
		image: NewImage(tutorial, src, width, height, ""),
	}
}

func (b *Browser) ToGraphQL() *model.Browser {
	return &model.Browser{
		Image: b.image.ToGraphQL(),
	}
}
