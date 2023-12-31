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

func (b *Browser) SetImage(tutorial, src string) error {
	errorPrefix := "SetImage failed"

	image, err := NewImage(tutorial, src, "") // last parameter, caption = "" for browser
	if err != nil {
		return fmt.Errorf("%s, %s", errorPrefix, err)
	}

	if err := image.copyFile(); err != nil {
		return fmt.Errorf("%s, %s", errorPrefix, err)
	}

	b.image = image
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
