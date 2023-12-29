package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

type Browser struct {
	image *Image
}

func NewBrowser() *Browser {
	return &Browser{}
}

func (b *Browser) SetImage(tutorial, src string, width, height int) error {
	errorPrefix := "SetImage failed"

	b.image = NewImage(tutorial, src, width, height, "")

	if err := b.image.copyFile(); err != nil {
		return fmt.Errorf("%s, %s", errorPrefix, err)
	}

	return nil
}

func (b *Browser) ToGraphQL() *model.Browser {
	var imageModel *model.Image
	if b.image != nil {
		imageModel = b.image.ToGraphQL()
	}

	return &model.Browser{
		Image: imageModel,
	}
}
