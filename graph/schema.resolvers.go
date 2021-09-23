package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/yuki383/arknights-data/graph/generated"
	"github.com/yuki383/arknights-data/graph/model"
	"github.com/yuki383/arknights-data/pkg/operators"
)

func (r *queryResolver) Operators(ctx context.Context, input *model.GetOperatorInput) ([]*model.Operator, error) {
	var resultOperators []*model.Operator
	o := operators.GetOperators()
	for _, operator := range o {
		// Check Operator Name
		if input.Name != nil && operator.Name != *input.Name {
			continue
		}

		// Check Operator Class
		if input.Class != nil && operator.Class != *input.Class {
			continue
		}

		// Check Operator Rarity
		if input.Rarity != nil && !matchRarity(operator.Rarity, input.Rarity) {
			continue
		}

		// Chack Operator Available Server
		if input.Available != nil && !matchAllAvailableServer(operator.Available, input.Available) {
			continue
		}

		op := model.Operator{
			Name:      operator.Name,
			Class:     operator.Class,
			Rarity:    operator.Rarity,
			Available: operator.Available,
		}
		resultOperators = append(resultOperators, &op)
	}

	return resultOperators, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func matchRarity(rarity int, expected []int) bool {
	for _, e := range expected {
		if rarity == e {
			return true
		}
	}

	return false
}
func matchAllAvailableServer(servers, expectedServers []string) bool {
	for _, e := range expectedServers {
		if !hasStrInSlice(servers, e) {
			return false
		}
	}

	return true
}
func hasStrInSlice(s []string, t string) bool {
	for _, e := range s {
		if e == t {
			return true
		}
	}

	return false
}
func (r *queryResolver) FindOperators(ctx context.Context, input *model.GetOperatorInput) ([]*model.Operator, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Operator(ctx context.Context, input *model.GetOperatorInput) (*model.Operator, error) {
	panic(fmt.Errorf("not implemented"))
}
