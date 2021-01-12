package user

import (
	"context"
	"fmt"
	"github.com/trinhdaiphuc/social-network/config"
	"github.com/trinhdaiphuc/social-network/internal"
	"github.com/trinhdaiphuc/social-network/internal/logger"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	Login(ctx context.Context, user *models.User) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
}

type service struct {
	repository UserRepository
	Logger     *logger.AppLog
}

func NewUserService(db *mongo.Database, log *logger.AppLog) UserService {
	r := NewUserRepository(db, log)
	return &service{repository: r, Logger: log}
}

func (s *service) Register(ctx context.Context, user *models.User) (*models.User, error) {
	// Hash password
	user.Password = internal.HashPassword(user.Password)

	// Create user
	user, err := s.repository.Create(ctx, user)
	if err != nil {
		s.Logger.Errorf("Register error %#v", err)
		if writeErr, ok := err.(mongo.WriteException); ok {
			if writeErr.WriteErrors[0].Code == 11000 {
				return nil, fmt.Errorf("Your account email or username is already taken.")
			}
		}
		return nil, err
	}

	// Create token
	user.Token, err = internal.CreateTokenWithUser(config.GetConfig().JwtKey, user, 5)
	return user, nil
}

func (s *service) Login(ctx context.Context, user *models.User) (*models.User, error) {
	// Get user by username
	getUser, err := s.repository.GetByUsername(ctx, user.Username)
	if err != nil {
		s.Logger.Errorf("Register error %#v", err)
		return nil, fmt.Errorf("Not found user")
	}

	// Check user password
	if ok := internal.CheckPasswordHash(user.Password, getUser.Password); !ok {
		return nil, fmt.Errorf("Invalid password")
	}

	// Create token
	getUser.Token, err = internal.CreateTokenWithUser(config.GetConfig().JwtKey, getUser, 12*60)
	return getUser, nil
}

func (s *service) GetUsers(ctx context.Context) ([]*models.User, error) {
	return s.repository.GetList(ctx)
}
