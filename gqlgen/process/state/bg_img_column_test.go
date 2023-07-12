package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestBackgroundImageColumnMutation1(t *testing.T) {
	s := state.BackgroundImageColumn{
		Width:  240,
		Height: 280,
		Path:   "/images/img1.png",
		URL:    "https://yourdomain.com/img1.png",
		Modal: state.Modal{
			Text:     "message in a modal",
			Position: state.ModalPositionTop,
		},
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLBgImgCol()
	goldenFile1 := "testdata/bg_img_column_golden1-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	s.Width = 1240
	s.Height = 1280
	s.Path = "/images/img2.png"
	s.URL = "https://yourdomain.com/img2.png"
	s.Modal.Text = "updated message in a modal"
	s.Modal.Position = state.ModalPositionCenter

	// ... has NO effect on the materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/bg_img_column_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLBgImgCol())
}

func TestBackgroundImageColumnMutation2(t *testing.T) {
	s := state.BackgroundImageColumn{
		Width:  240,
		Height: 280,
		Path:   "/images/img1.png",
		URL:    "https://yourdomain.com/img1.png",
		Modal: state.Modal{
			Text:     "message in a modal",
			Position: state.ModalPositionTop,
		},
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLBgImgCol()
	goldenFile1 := "testdata/bg_img_column_golden2-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	*gqlModel.Width = 1240
	*gqlModel.Height = 1280
	*gqlModel.Path = "/images/img2.png"
	*gqlModel.URL = "https://yourdomain.com/img2.png"
	*gqlModel.Modal.Text = "updated message in a modal"
	*gqlModel.Modal.Position = model.ModalPositionCenter

	// ... has NO effect on a RE-materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLBgImgCol())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/bg_img_column_golden2-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}
