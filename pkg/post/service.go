package post

import (
	"context"
	"fmt"
	"github.com/trinhdaiphuc/social-network/common"
	"github.com/trinhdaiphuc/social-network/internal/logger"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"github.com/trinhdaiphuc/social-network/pkg/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostService interface {
	GetPosts(ctx context.Context) ([]*models.Post, error)
	GetPost(ctx context.Context, id primitive.ObjectID) (*models.Post, error)
	DeletePost(ctx context.Context, id primitive.ObjectID) (string, error)
	CreatePost(ctx context.Context, p *models.Post) (*models.Post, error)
}

type service struct {
	repository PostRepository
	Logger     *logger.AppLog
	NewPostChannel chan *models.Post
}

func NewPostService(db *mongo.Database, log *logger.AppLog) PostService {
	r := NewPostRepository(db, log)
	return &service{repository: r, Logger: log}
}

func (p *service) GetPosts(ctx context.Context) ([]*models.Post, error) {
	return p.repository.GetList(ctx)
}

func (p *service) GetPost(ctx context.Context, id primitive.ObjectID) (*models.Post, error) {
	return p.repository.GetByID(ctx, id)
}

func (p *service) DeletePost(ctx context.Context, id primitive.ObjectID) (string, error) {
	return p.repository.DeleteByID(ctx, id)
}

func (p *service) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	if _, err := user.GetUserRepository().GetByUsername(ctx, post.Username); err != nil {
		return nil, fmt.Errorf("Not found username")
	}
	newPost, err := p.repository.Create(ctx, post)
	if err != nil {
		return nil, err
	}
	for _, observer := range common.AddChannelObserver {
		observer <- newPost
		//This sends new post to client via socket
	}
	return newPost, nil
}
