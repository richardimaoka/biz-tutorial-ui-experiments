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
)

// PageState is the resolver for the pageState field.
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

// OpenFile is the resolver for the openFile field.
func (r *sourceCodeResolver) OpenFile(ctx context.Context, obj *model.SourceCode, filePath *string) (*model.OpenFile, error) {
	var dirName = "data/apollo-client-getting-started/state"
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

	var pageState model.PageState
	err = json.Unmarshal(data, &pageState)
	if err != nil {
		return nil, fmt.Errorf("internal server error - failed to unmarshal PageState from %s", filename)
	}

	if filePath == nil {
		log.Printf("returning nil as filePath empty")
		// return nil openFile, instead of error, so that the entire page can still render
		// TODO: enable default open file returning, once SourceCode has defaultOpenFilePath set
		return nil, nil
	}

	openFile, ok := pageState.SourceCode.FileContents[*filePath]
	if !ok {
		log.Printf("OpenFile() file not found: %s", *filePath)
		// return nil openFile, instead of error, so that the entire page can still render
		return nil, nil
	}

	log.Printf("OpenFile() returning file for: %s", *filePath)
	return &openFile, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// SourceCode returns SourceCodeResolver implementation.
func (r *Resolver) SourceCode() SourceCodeResolver { return &sourceCodeResolver{r} }

type queryResolver struct{ *Resolver }
type sourceCodeResolver struct{ *Resolver }
