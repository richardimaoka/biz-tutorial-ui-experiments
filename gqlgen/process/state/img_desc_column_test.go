package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestImageDescriptionColumnMutation1(t *testing.T) {
	s := state.ImageDescriptionColumn{
		Image: state.ImageCentered{
			Width:          240,
			Height:         280,
			OriginalWidth:  480,
			OriginalHeight: 560,
			Path:           "/images/img1.png",
			URL:            "https://yourdomain.com/img1.png",
		},
		Description: state.Markdown{
			Contents:  "markdown default contents",
			Alignment: state.MarkdownAlignmentLeft,
		},
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLImgDescCol()
	goldenFile1 := "testdata/img_desc_column_golden1-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	s.Image.Width = 1240
	s.Image.Height = 1280
	s.Image.Path = "/images/img2.png"
	s.Image.URL = "https://yourdomain.com/img2.png"
	s.Description.Contents = "updated contents in markdown"
	s.Description.Alignment = state.MarkdownAlignmentCenter

	// ... has NO effect on the materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/img_desc_column_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLImgDescCol())
}

func TestImageDescriptionColumnMutation2(t *testing.T) {
	s := state.ImageDescriptionColumn{
		Image: state.ImageCentered{
			Width:          240,
			Height:         280,
			OriginalWidth:  480,
			OriginalHeight: 560,
			Path:           "/images/img1.png",
			URL:            "https://yourdomain.com/img1.png",
		},
		Description: state.Markdown{
			Contents:  "markdown default contents",
			Alignment: state.MarkdownAlignmentLeft,
		},
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLImgDescCol()
	goldenFile1 := "testdata/img_desc_column_golden2-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	*gqlModel.Image.Width = 1240
	*gqlModel.Image.Height = 1280
	*gqlModel.Image.Path = "/images/img2.png"
	*gqlModel.Image.URL = "https://yourdomain.com/img2.png"
	*gqlModel.Description.Contents = "updated contents in markdown"
	*gqlModel.Description.Alignment = model.MarkdownAlignmentCenter

	// ... has NO effect on a RE-materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLImgDescCol())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/img_desc_column_golden2-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}
