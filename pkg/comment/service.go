package comment

import (
	"context"
	"fmt"
	"github.com/trinhdaiphuc/social-network/internal/logger"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"github.com/trinhdaiphuc/social-network/pkg/post"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentService interface {
	CreateComment(ctx context.Context, postID primitive.ObjectID, comment *models.Comment) (*models.Post, error)
	DeleteComment(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID, username string) (*models.Post, error)
}

type service struct {
	repository CommentRepository
	Logger     *logger.AppLog
}

func NewCommentService(log *logger.AppLog) CommentService {
	return &service{
		repository: NewCommentRepository(log),
		Logger:     log,
	}
}

func (s *service) CreateComment(ctx context.Context, postID primitive.ObjectID, comment *models.Comment) (*models.Post, error) {
	return s.repository.Create(ctx, postID, comment)
}

func (s *service) DeleteComment(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID, username string) (*models.Post, error) {
	p, err := post.GetPostRepository().GetCommentByID(ctx, postID, commentID)
	if err != nil {
		return nil, err
	}
	if p.Username == username {
		return s.repository.DeleteByID(ctx, postID, commentID)
	}
	for _, v := range p.Comments {
		if v.ID == commentID && v.Username == username {
			return s.repository.DeleteByID(ctx, postID, commentID)
		}
	}
	return nil, fmt.Errorf("Action not allowed")
}
