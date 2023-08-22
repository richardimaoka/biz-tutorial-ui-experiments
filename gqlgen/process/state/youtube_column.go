package state

import "github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"

type YouTubeColumn struct {
	YouTube YouTubeEmbed
}

func NewYouTubeColumn() *YouTubeColumn {
	return &YouTubeColumn{}
}

func (p *YouTubeColumn) Process(videoId string, width, height int) error {

	p.YouTube = YouTubeEmbed{
		VideoId: videoId,
		Width:   width,
		Height:  height,
	}

	return nil
}

func (p *YouTubeColumn) ToGraphQLYouTubeColumn() *model.YouTubeColumn {
	// copy to avoid mutation effect afterwards
	youtube := p.YouTube.ToGraphQLYouTubeEmbed()

	return &model.YouTubeColumn{
		Youtube: youtube,
	}
}
