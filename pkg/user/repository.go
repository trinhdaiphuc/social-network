package user

import (
	"context"
	"github.com/trinhdaiphuc/social-network/internal/logger"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "users"
)

var (
	userRepo *repository
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetList(ctx context.Context) ([]*models.User, error)
	GetByUsername(ctx context.Context, username string) (*models.User, error)
}

type repository struct {
	Collection *mongo.Collection
	Logger     *logger.AppLog
}

func NewUserRepository(db *mongo.Database, log *logger.AppLog) UserRepository {
	mod := []mongo.IndexModel{
		{
			Keys: bson.M{
				"email": -1, // index in ascending order
			},
			// create UniqueIndex option
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.M{
				"username": -1, // index in ascending order
			},
			// create UniqueIndex option
			Options: options.Index().SetUnique(true),
		},
	}

	ctx := context.Background()
	userCollection := db.Collection(collectionName)
	userCollection.Indexes().CreateMany(ctx, mod)
	userRepo = &repository{
		Collection: userCollection,
		Logger:     log,
	}
	return userRepo
}

func GetUserRepository() UserRepository {
	return userRepo
}

func (r *repository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	result, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r *repository) GetList(ctx context.Context) ([]*models.User, error) {
	var users []*models.User

	cursor, err := r.Collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	cursor.All(ctx, &users)
	return users, nil
}

func (r *repository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	result := r.Collection.FindOne(ctx, bson.M{"username": username})
	user := &models.User{}
	if err := result.Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}
