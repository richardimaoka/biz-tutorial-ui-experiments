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

// Step is the resolver for the step field.
func (r *queryResolver) Step(ctx context.Context, stepNum *int) (*model.Step, error) {
	filename := fmt.Sprintf("data/tutorial1/step%2d.json", stepNum)
	log.Printf("reading data from %s", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("stepNum = %d not found", *stepNum))
	}

	var step model.Step

	//TODO: Instead of loading everything in one shot, load non-Union and Union separately, and combine them
	err = json.Unmarshal(data, &step)
	if err != nil {
		log.Printf("ERROR: %s", err)
		return nil, errors.New("internal server error")
	}

	return &step, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) TerminalElements(ctx context.Context, step *int) ([]model.TerminalElement, error) {
	panic(fmt.Errorf("not implemented: TerminalElements - terminalElements"))
}

type mutationResolver struct{ *Resolver }
