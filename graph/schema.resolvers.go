package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/yuki383/arknights-data/graph/generated"
	"github.com/yuki383/arknights-data/graph/model"
)

func (r *queryResolver) Operators(ctx context.Context) ([]*model.Operator, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Operator(ctx context.Context, input *model.GetOperatorInput) (*model.Operator, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
