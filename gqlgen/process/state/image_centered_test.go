package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestImageCenteredMutation1(t *testing.T) {
	s := state.ImageCentered{
		Width:          240,
		Height:         280,
		OriginalWidth:  480,
		OriginalHeight: 560,
		Path:           "/images/img1.png",
		URL:            "https://yourdomain.com/img1.png",
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLImageCentered()
	goldenFile1 := "testdata/image_centered_golden1-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	s.Width = 1240
	s.Height = 1280
	s.OriginalWidth = 2480
	s.OriginalHeight = 2560
	s.Path = "/images/img2.png"
	s.URL = "https://yourdomain.com/img2.png"

	// ... has NO effect on the materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/image_centered_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLImageCentered())
}

func TestImageCenteredMutation2(t *testing.T) {
	s := state.ImageCentered{
		Width:          240,
		Height:         280,
		OriginalWidth:  480,
		OriginalHeight: 560,
		Path:           "/images/img1.png",
		URL:            "https://yourdomain.com/img1.png",
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLImageCentered()
	goldenFile1 := "testdata/image_centered_golden2-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	*gqlModel.Width = 1240
	*gqlModel.Height = 1280
	// *gqlModel.OriginalWidth = 2480
	// *gqlModel.OriginalHeight = 2560
	*gqlModel.Path = "/images/img2.png"
	*gqlModel.URL = "https://yourdomain.com/img2.png"

	// ... has NO effect on a RE-materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLImageCentered())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/image_centered_golden2-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}
