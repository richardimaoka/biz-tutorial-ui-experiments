package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/richardimaoka/biz-tutorial-ui-experiments/gqlgen/graph/model"
)

// PageState is the resolver for the pageState field.
func (r *queryResolver) PageState(ctx context.Context, step *string) (*model.PageState, error) {
	var filename string
	if step == nil {
		filename = "data/tutorial3/state-000.json"
	} else {
		filename = fmt.Sprintf("data/state/state-%s.json", *step)
	}

	log.Printf("reading data from %s", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var pageState model.PageState
	err = json.Unmarshal(data, &pageState)
	if err != nil {
		return nil, fmt.Errorf("internal server error - failed to unmarshal PageState from %s", filename)
	}

	return &pageState, nil
}

// Terminal is the resolver for the terminal field.
func (r *queryResolver) Terminal(ctx context.Context, step int) (*model.Terminal, error) {
	filename := fmt.Sprintf("data/tutorial2/terminal%03d.json", step)
	log.Printf("reading data from %s", filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s not found", filename))
	}

	var terminal model.Terminal

	err = json.Unmarshal(data, &terminal)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return nil, errors.New("internal server error")
	}

	return &terminal, nil
}

// OpenFile is the resolver for the openFile field.
func (r *sourceCodeResolver) OpenFile(ctx context.Context, obj *model.SourceCode, filePath string) (*model.OpenFile, error) {
	filename := fmt.Sprintf("data/state/state-%s.json", obj.Step)
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

	openFile, ok := pageState.SourceCode.FileContents[filePath]
	if !ok {
		return nil, fmt.Errorf("internal server error - cannot load openFile %s", filePath)
	}
	fmt.Print("oooopenfile", openFile)
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
type mutationResolver struct{ *Resolver }
