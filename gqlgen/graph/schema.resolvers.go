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
	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/internal"
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

// Test is the resolver for the _test field.
func (r *queryResolver) Test(ctx context.Context) (*model.TestObjs, error) {
	testObj := model.TestObjs{}
	return &testObj, nil
}

// OpenFile is the resolver for the openFile field.
func (r *sourceCodeResolver) OpenFile(ctx context.Context, obj *model.SourceCode, filePath *string) (*model.OpenFile, error) {
	var dirName = fmt.Sprintf("data/%s/state", obj.Tutorial)
	var initialStep = "_initial"

	var filename string
	if obj.Step == "" {
		filename = fmt.Sprintf(dirName+"/%s.json", initialStep)
	} else {
		filename = fmt.Sprintf(dirName+"/%s.json", obj.Step)
	}

	log.Printf("OpenFile() reading data from %s", filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var page model.Page
	err = json.Unmarshal(data, &page)
	if err != nil {
		return nil, fmt.Errorf("internal server error - failed to unmarshal page from %s", filename)
	}

	var openFilePath string
	if filePath != nil {
		openFilePath = *filePath
	} else if obj.DefaultOpenFilePath != "" {
		log.Printf("filePath argument is null, so using default file path %s", obj.DefaultOpenFilePath)
		openFilePath = obj.DefaultOpenFilePath
	} else {
		log.Printf("returning empty openFile as filePath argument is null and defaultOpenFilePath is empty")
		// return nil openFile, instead of error, so that the entire page can still render
		// TODO: enable default open file returning, once SourceCode has defaultOpenFilePath set
		return nil, nil
	}

	var sourceCode *model.SourceCode
	for _, col := range page.Columns {
		if col.Name != nil && *col.Name == "Source Code" {
			scCol, ok := col.Column.(*model.SourceCodeColumn)
			if !ok {
				log.Printf("OpenFile() failed to cast column to SourceCodeColumn")
				return nil, fmt.Errorf("internal server error")
			}
			sourceCode = scCol.SourceCode
		}
	}

	if sourceCode == nil {
		log.Printf("source code is nil")
		return nil, fmt.Errorf("internal server error")
	}

	openFile, ok := sourceCode.FileContents[openFilePath]
	if !ok {
		log.Printf("OpenFile() file not found: %s", openFilePath)
		// return nil openFile, instead of error, so that the entire page can still render
		return nil, nil
	}

	log.Printf("OpenFile() returning file for: %s", openFilePath)
	return &openFile, nil
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
	err := internal.JsonRead2(filename, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

// AppTestTutorialColumnsPage is the resolver for the appTestTutorialColumnsPage field.
func (r *testObjsResolver) AppTestTutorialColumnsPage(ctx context.Context, obj *model.TestObjs) (*model.Page2, error) {
	var m model.Page2
	err := internal.JsonRead2("data/_test/appTestTutorialColumnsPage.json", &m)
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) PageState(ctx context.Context, step *string) (*model.PageState, error) {
	var dirName = "data/apollo-client-getting-started/state"
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

	var pageState model.PageState
	err = json.Unmarshal(data, &pageState)
	if err != nil {
		log.Printf("failed to read data from %s, %s", filename, err)
		return nil, fmt.Errorf("internal server error %s", *step)
	}

	return &pageState, nil
}
func (r *queryResolver) SourceCode(ctx context.Context) (*model.SourceCode, error) {
	panic(fmt.Errorf("not implemented: SourceCode - sourceCode"))
}
