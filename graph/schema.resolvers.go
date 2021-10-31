package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lythesuknguyen/rikkei/graph/generated"
	"github.com/lythesuknguyen/rikkei/graph/model"
)

func (r *queryResolver) Members(ctx context.Context, input *model.Params) ([]*model.Member, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Skills(ctx context.Context, input *model.Params) ([]*model.Skill, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
