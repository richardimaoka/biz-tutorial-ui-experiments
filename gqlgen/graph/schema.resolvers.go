package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal/jsonwrap"
)

// Page is the resolver for the page field.
func (r *queryResolver) Page(ctx context.Context, tutorial string, step *string) (*model.Page, error) {
	var dirName = fmt.Sprintf("data/%s/state", tutorial)
	var initialStep = "_initial"

	var filename string
	if step == nil {
		filename = fmt.Sprintf(dirName+"/%s.json", initialStep)
	} else {
		filename = fmt.Sprintf(dirName+"/%s.json", *step)
	}

	log.Printf("reading data from %s", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var page model.Page
	err = json.Unmarshal(data, &page)
	if err != nil {
		log.Printf("failed to read data from %s, %s", filename, err)
		return nil, fmt.Errorf("internal server error %s", *step)
	}

	return &page, nil
}

// Page2 is the resolver for the page2 field.
func (r *queryResolver) Page2(ctx context.Context, tutorial string, step *string) (*model.Page2, error) {
	panic(fmt.Errorf("not implemented: Page2 - page2"))
}

// Test is the resolver for the _test field.
func (r *queryResolver) Test(ctx context.Context) (*model.TestObjs, error) {
	testObj := model.TestObjs{}
	return &testObj, nil
}

// OpenFile is the resolver for the openFile field.
func (r *sourceCodeResolver) OpenFile(ctx context.Context, obj *model.SourceCode, filePath *string) (*model.OpenFile, error) {
	return nil, nil
}

// AppTestTerminalPage is the resolver for the appTestTerminalPage field.
func (r *testObjsResolver) AppTestTerminalPage(ctx context.Context, obj *model.TestObjs, step *int) (*model.TerminalColumn2, error) {
	var filename string
	if step == nil {
		filename = "data/_test/appTestTerminalPage/1.json"
	} else {
		filename = fmt.Sprintf("data/_test/appTestTerminalPage/%d.json", *step)
	}

	var m model.TerminalColumn2
	err := jsonwrap.Read(filename, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// AppTestTutorialColumnsPage is the resolver for the appTestTutorialColumnsPage field.
func (r *testObjsResolver) AppTestTutorialColumnsPage(ctx context.Context, obj *model.TestObjs) (*model.Page2, error) {
	var m model.Page2
	err := jsonwrap.Read("data/_test/appTestTutorialColumnsPage.json", &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// AppTestTutorialTutorialPage is the resolver for the appTestTutorialTutorialPage field.
func (r *testObjsResolver) AppTestTutorialTutorialPage(ctx context.Context, obj *model.TestObjs) (*model.Page2, error) {
	var m model.Page2
	err := jsonwrap.Read("data/_test/appTestTutorialTutorialPage.json", &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// AppTestSourcecodeFilecontentPage is the resolver for the appTestSourcecodeFilecontentPage field.
func (r *testObjsResolver) AppTestSourcecodeFilecontentPage(ctx context.Context, obj *model.TestObjs, step int) (*model.OpenFile, error) {
	filename := fmt.Sprintf("data/_test/appTestSourcecodeFilecontentPage/%d.json", step)

	var m model.OpenFile
	err := jsonwrap.Read(filename, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// SourceCode returns SourceCodeResolver implementation.
func (r *Resolver) SourceCode() SourceCodeResolver { return &sourceCodeResolver{r} }

// TestObjs returns TestObjsResolver implementation.
func (r *Resolver) TestObjs() TestObjsResolver { return &testObjsResolver{r} }

type queryResolver struct{ *Resolver }
type sourceCodeResolver struct{ *Resolver }
type testObjsResolver struct{ *Resolver }
