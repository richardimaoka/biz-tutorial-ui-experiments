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
	var filename string
	if step == nil {
		filename = fmt.Sprintf(dirName+"/state-%s.json", firstStep)
	} else {
		filename = fmt.Sprintf(dirName+"/state-%s.json", *step)
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
func (r *queryResolver) Page(ctx context.Context, tutorial string, step *string) (*model.PageState, error) {
	if tutorial == "" {
		return nil, fmt.Errorf("tutorial name must be specified")
	}

	var stepStr string
	if step == nil {
		log.Printf("step is nil, reading first step for %s", tutorial)
		var firstStep string
		firstStepFile := fmt.Sprintf("data/%s/first-step.json", tutorial)
		bytes, err := os.ReadFile(firstStepFile)
		if err != nil {
			return nil, fmt.Errorf("tutorial = %s does not exist", tutorial)
		}
		if err := json.Unmarshal(bytes, &firstStep); err != nil {
			return nil, fmt.Errorf("tutorial = %s doesn't define a valid first step", tutorial)
		}
		if firstStep == "" {
			log.Printf("first step from file = %s was empty", firstStepFile)
			return nil, fmt.Errorf("tutorial = %s doesn't define a valid first step", tutorial)
		}
		stepStr = firstStep
	} else {
		stepStr = *step
	}
	filename := fmt.Sprintf("data/%s/state/state-%s.json", tutorial, stepStr)

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("failed to read %s, %s", filename, err)
		return nil, fmt.Errorf("failed to read data for tutorial = %s, step = `%s`", tutorial, stepStr)
	}

	var pageState model.PageState
	err = json.Unmarshal(data, &pageState)
	if err != nil {
		log.Printf("failed to read data from %s, %s", filename, err)
		return nil, fmt.Errorf("internal server error %s", *step)
	}

	return &pageState, nil
}

// OpenFile is the resolver for the openFile field.
func (r *sourceCodeResolver) OpenFile(ctx context.Context, obj *model.SourceCode, filePath *string) (*model.OpenFile, error) {
	var filename string
	if obj.Step == "" {
		filename = fmt.Sprintf(dirName+"/state-%s.json", firstStep)
	} else {
		filename = fmt.Sprintf(dirName+"/state-%s.json", obj.Step)
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var dirName = "data/gqlgensandbox/state"
var firstStep = "aedf3711-8f47-4c9b-af54-12eb7c0d2d87"

type mutationResolver struct{ *Resolver }
