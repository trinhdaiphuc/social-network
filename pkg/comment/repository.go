package comment

import (
	"context"
	"github.com/trinhdaiphuc/social-network/internal/logger"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"github.com/trinhdaiphuc/social-network/pkg/post"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository interface {
	Create(ctx context.Context, postID primitive.ObjectID, comment *models.Comment) (*models.Post, error)
	DeleteByID(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID) (*models.Post, error)
}

type repository struct {
	Logger *logger.AppLog
}

func NewCommentRepository(log *logger.AppLog) CommentRepository {
	return &repository{Logger: log}
}

func (r *repository) Create(ctx context.Context, postID primitive.ObjectID, comment *models.Comment) (*models.Post, error) {
	return post.GetPostRepository().CreateComment(ctx, postID, comment)
}

func (r *repository) DeleteByID(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID) (*models.Post, error) {
	return post.GetPostRepository().DeleteCommentByID(ctx, postID, commentID)
}
