package state

import (
	"fmt"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
)

type YouTubeEmbed struct {
	Width   int
	Height  int
	VideoId string
}

func (p *YouTubeEmbed) ToGraphQLYouTubeEmbed() *model.YouTubeEmbed {
	// copy to avoid mutation effect afterwards
	width := p.Width
	height := p.Height
	url := internal.StringRef(fmt.Sprintf("https://www.youtube.com/embed/%s?autoplay=1&mute=1", p.VideoId))

	return &model.YouTubeEmbed{
		Width:    &width,
		Height:   &height,
		EmbedURL: url,
	}
}
