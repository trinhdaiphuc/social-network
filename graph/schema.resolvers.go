package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/satori/go.uuid"
	"github.com/trinhdaiphuc/social-network/common"
	"github.com/trinhdaiphuc/social-network/graph/generated"
	"github.com/trinhdaiphuc/social-network/graph/model"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"github.com/trinhdaiphuc/social-network/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *commentResolver) ID(ctx context.Context, obj *models.Comment) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *likeResolver) ID(ctx context.Context, obj *models.Like) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, body string) (*models.Post, error) {
	user, err := tools.ForUserContext(ctx)
	if user == nil || err != nil {
		return nil, fmt.Errorf("Unauthorize")
	}
	newPost := &models.Post{
		Body:      body,
		CreatedAt: time.Now().Format(time.RFC3339),
		Username:  user.Username,
	}
	newPost, err = r.PostService.CreatePost(ctx, newPost)
	if err != nil {
		return nil, err
	}
	return newPost, nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (string, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", fmt.Errorf("Invalid id")
	}
	return r.PostService.DeletePost(ctx, oid)
}

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: password,
	}
	user, err := r.UserService.Login(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) Register(ctx context.Context, registerInput model.RegisterInput) (*models.User, error) {
	if registerInput.Password != registerInput.ConfirmPassword {
		return nil, fmt.Errorf("password and confirm password is not match")
	}
	newUser := &models.User{
		Email:     registerInput.Email,
		Username:  registerInput.Username,
		Password:  registerInput.Password,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	newUser, err := r.UserService.Register(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, postID string, body string) (*models.Post, error) {
	user, err := tools.ForUserContext(ctx)
	if user == nil || err != nil {
		return nil, fmt.Errorf("Unauthorize")
	}
	comment := &models.Comment{
		ID:        primitive.NewObjectID(),
		Username:  user.Username,
		Body:      body,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	postOID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, fmt.Errorf("Invalid post id")
	}
	return r.CommentService.CreateComment(ctx, postOID, comment)
}

func (r *mutationResolver) DeleteComment(ctx context.Context, postID string, commentID string) (*models.Post, error) {
	user, err := tools.ForUserContext(ctx)
	if user == nil || err != nil {
		return nil, fmt.Errorf("Unauthorize")
	}
	postOID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, fmt.Errorf("Invalid post id")
	}
	commentOID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, fmt.Errorf("Invalid comment id")
	}
	return r.CommentService.DeleteComment(ctx, postOID, commentOID, user.Username)
}

func (r *mutationResolver) LikePost(ctx context.Context, postID string) (*models.Post, error) {
	user, err := tools.ForUserContext(ctx)
	if user == nil || err != nil {
		return nil, fmt.Errorf("Unauthorize")
	}
	postOID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, fmt.Errorf("Invalid post id")
	}

	return r.LikeService.LikePost(ctx, postOID, user.Username)
}

func (r *postResolver) ID(ctx context.Context, obj *models.Post) (string, error) {
	return obj.ID.Hex(), nil
}

func (r *postResolver) LikeCount(ctx context.Context, obj *models.Post) (int, error) {
	return len(obj.Likes), nil
}

func (r *postResolver) CommentCount(ctx context.Context, obj *models.Post) (int, error) {
	return len(obj.Comments), nil
}

func (r *queryResolver) GetPosts(ctx context.Context) ([]*models.Post, error) {
	posts, err := r.PostService.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *queryResolver) GetPost(ctx context.Context, id string) (*models.Post, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("Invalid id")
	}
	return r.PostService.GetPost(ctx, oid)
}

func (r *queryResolver) GetUsers(ctx context.Context) ([]*models.User, error) {
	return r.UserService.GetUsers(ctx)
}

func (r *subscriptionResolver) NewPost(ctx context.Context) (<-chan *models.Post, error) {
	id := uuid.NewV4().String()
	events := make(chan *models.Post, 1)

	go func() {
		<-ctx.Done()
		delete(common.AddChannelObserver, id)
	}()

	common.AddChannelObserver[id] = events
	return events, nil
}

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.ID.Hex(), nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

// Like returns generated.LikeResolver implementation.
func (r *Resolver) Like() generated.LikeResolver { return &likeResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type commentResolver struct{ *Resolver }
type likeResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
