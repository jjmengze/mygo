package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/jjmengze/mygo/internal/delivery/graph/generated"
	"github.com/jjmengze/mygo/internal/delivery/graph/model"
)

func (r *mutationResolver) AddPost(ctx context.Context, input model.AddPostInput) (*model.Post, error) {
	fmt.Println("tittle", input.Title)
	fmt.Println("content", input.Content)
	return &model.Post{
		ID: "100",
		//Author:     nil,
		Title:   &input.Title,
		Content: input.Content,
		//LikeGivers: nil,
	}, nil
}

func (r *mutationResolver) LikePost(ctx context.Context, postID string) (*model.Post, error) {
	tittle := "test"
	content := "test"
	return &model.Post{
		ID: postID,
		//Author:     nil,
		Title:   &tittle,
		Content: &content,
		//LikeGivers: nil,
	}, nil
}

func (r *postResolver) Author(ctx context.Context, obj *model.Post) (*model.User, error) {
	//id, _ := strconv.Atoi(obj.ID)
	name := "test"
	return &model.User{
		ID:   obj.ID,
		Name: &name,
	}, nil
}

func (r *postResolver) LikeGivers(ctx context.Context, obj *model.Post) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
