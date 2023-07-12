package state_test

import (
	"testing"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/process/state"
)

func TestModalMutation1(t *testing.T) {
	s := state.Modal{
		Text:     "message in modal",
		Position: state.ModalPositionTop,
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLModal()
	goldenFile1 := "testdata/modal_golden1-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the state ...
	s.Text = "updated message in modal"
	s.Position = state.ModalPositionCenter

	// ... has NO effect on the materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, gqlModel)

	// ... has effect on a RE-materialized GraphQL model
	goldenFile2 := "testdata/modal_golden1-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, s.ToGraphQLModal())
}

func TestModalMutation2(t *testing.T) {
	s := state.Modal{
		Text:     "message in modal",
		Position: state.ModalPositionTop,
	}

	// once GraphQL model is materialized...
	gqlModel := s.ToGraphQLModal()
	goldenFile1 := "testdata/modal_golden2-1.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile1, gqlModel)

	// ... mutation to the materialized model ...
	*gqlModel.Text = "updated message in modal"
	*gqlModel.Position = model.ModalPositionCenter

	// ... has NO effect on a RE-materialized GraphQL model
	internal.CompareAfterMarshal(t, goldenFile1, s.ToGraphQLModal())

	// ... has effect on the materialized GraphQL model
	goldenFile2 := "testdata/modal_golden2-2.json"
	internal.CompareWitGoldenFile(t, *updateFlag, goldenFile2, gqlModel)
}
