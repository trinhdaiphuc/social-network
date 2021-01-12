package graph

import (
	"github.com/trinhdaiphuc/social-network/internal/logger"
	"github.com/trinhdaiphuc/social-network/pkg/comment"
	"github.com/trinhdaiphuc/social-network/pkg/like"
	"github.com/trinhdaiphuc/social-network/pkg/post"
	"github.com/trinhdaiphuc/social-network/pkg/user"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB             *mongo.Database
	Logger         *logger.AppLog
	PostService    post.PostService
	UserService    user.UserService
	CommentService comment.CommentService
	LikeService    like.LikeService
}

func NewResolver(db *mongo.Database) *Resolver {
	logger := logger.NewAppLog()
	return &Resolver{
		DB:             db,
		Logger:         logger,
		PostService:    post.NewPostService(db, logger),
		UserService:    user.NewUserService(db, logger),
		CommentService: comment.NewCommentService(logger),
		LikeService:    like.NewLikeService(logger),
	}
}
