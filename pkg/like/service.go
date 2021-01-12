package like

import (
	"context"
	"github.com/trinhdaiphuc/social-network/internal/logger"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"github.com/trinhdaiphuc/social-network/pkg/post"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeService interface {
	LikePost(ctx context.Context, postID primitive.ObjectID, username string) (*models.Post, error)
}

type service struct {
	repository LikeRepository
	Logger     *logger.AppLog
}

func NewLikeService(log *logger.AppLog) LikeService {
	return &service{
		repository: NewLikeRepository(log),
		Logger:     log,
	}
}

func (s *service) LikePost(ctx context.Context, postID primitive.ObjectID, username string) (*models.Post, error) {
	postRepo := post.GetPostRepository()
	post, err := postRepo.FindLikeByUsername(ctx, postID, username)
	if err != nil {
		if err.Error() == "Not found post" {
			post, err = postRepo.CreateLike(ctx, postID, username)
			if err != nil {
				return nil, err
			}
			return post, nil
		}
		return nil, err
	}
	return postRepo.DeleteLikeByID(ctx, postID, username)
}
