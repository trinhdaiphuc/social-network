package post

import (
	"context"
	"fmt"
	"github.com/trinhdaiphuc/social-network/internal/logger"
	"github.com/trinhdaiphuc/social-network/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	collectionName = "posts"
)

var (
	postRepo *repository
)

type PostRepository interface {
	GetList(ctx context.Context) ([]*models.Post, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*models.Post, error)
	DeleteByID(ctx context.Context, id primitive.ObjectID) (string, error)
	Create(ctx context.Context, p *models.Post) (*models.Post, error)
	CreateComment(ctx context.Context, postID primitive.ObjectID, comment *models.Comment) (*models.Post, error)
	GetCommentByID(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID) (*models.Post, error)
	DeleteCommentByID(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID) (*models.Post, error)
	FindLikeByUsername(ctx context.Context, postID primitive.ObjectID, username string) (*models.Post, error)
	CreateLike(ctx context.Context, postID primitive.ObjectID, username string) (*models.Post, error)
	DeleteLikeByID(ctx context.Context, postID primitive.ObjectID, username string) (*models.Post, error)
}

type repository struct {
	Collection *mongo.Collection
	Logger     *logger.AppLog
}

func NewPostRepository(db *mongo.Database, log *logger.AppLog) PostRepository {
	postRepo = &repository{
		Collection: db.Collection(collectionName),
		Logger:     log,
	}
	return postRepo
}

func GetPostRepository() PostRepository {
	return postRepo
}

func (r *repository) GetList(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post

	cursor, err := r.Collection.Find(ctx, bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		return nil, err
	}

	cursor.All(ctx, &posts)
	return posts, nil
}

func (r *repository) Create(ctx context.Context, p *models.Post) (*models.Post, error) {
	post := &models.Post{
		Body:      p.Body,
		CreatedAt: time.Now().Format(time.RFC3339),
		Username:  p.Username,
	}

	resullt, err := r.Collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	post.ID = resullt.InsertedID.(primitive.ObjectID)
	return post, nil
}

func (r *repository) GetByID(ctx context.Context, id primitive.ObjectID) (*models.Post, error) {
	filter := bson.M{"_id": id}

	result := r.Collection.FindOne(ctx, filter)
	post := &models.Post{}
	if err := result.Decode(post); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Not found post")
		}
		return nil, err
	}
	return post, nil
}

func (r *repository) DeleteByID(ctx context.Context, id primitive.ObjectID) (string, error) {
	filter := bson.M{"_id": id}

	result, err := r.Collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("deleted %v documents", result.DeletedCount), nil
}

func (r *repository) CreateComment(ctx context.Context, postID primitive.ObjectID, comment *models.Comment) (*models.Post, error) {
	filter := bson.M{"_id": postID}
	post := &models.Post{}
	update := bson.M{"$addToSet": bson.M{"comments": bson.M{"$each": []models.Comment{*comment}}}}
	err := r.Collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Not found post")
		}
		return nil, err
	}
	return post, err
}

func (r *repository) GetCommentByID(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID) (*models.Post, error) {
	filter := bson.M{"_id": postID, "comments._id": commentID}
	post := &models.Post{}
	err := r.Collection.FindOne(ctx, filter).Decode(post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Not found post")
		}
		return nil, err
	}
	return post, err
}

func (r *repository) DeleteCommentByID(ctx context.Context, postID primitive.ObjectID, commentID primitive.ObjectID) (*models.Post, error) {
	filter := bson.M{"_id": postID}
	post := &models.Post{}
	update := bson.M{
		"$pull": bson.M{
			"comments": bson.M{
				"_id": commentID,
			},
		},
	}
	err := r.Collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Not found post")
		}
		return nil, err
	}
	return post, err
}

func (r *repository) FindLikeByUsername(ctx context.Context, postID primitive.ObjectID, username string) (*models.Post, error) {
	filter := bson.M{"_id": postID, "likes.username": username}
	post := &models.Post{}
	if err := r.Collection.FindOne(ctx, filter).Decode(post); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Not found post")
		}
		return nil, err
	}
	return post, nil
}

func (r *repository) CreateLike(ctx context.Context, postID primitive.ObjectID, username string) (*models.Post, error) {
	filter := bson.M{"_id": postID}
	post := &models.Post{}
	update := bson.M{"$addToSet": bson.M{"likes": bson.M{"$each": []models.Like{
		{
			ID:        primitive.NewObjectID(),
			Username:  username,
			CreatedAt: time.Now().Format(time.RFC3339),
		},
	}}}}
	err := r.Collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Not found post")
		}
		return nil, err
	}
	return post, err
}

func (r *repository) DeleteLikeByID(ctx context.Context, postID primitive.ObjectID, username string) (*models.Post, error) {
	filter := bson.M{"_id": postID}
	post := &models.Post{}
	update := bson.M{"$pull": bson.M{"likes": bson.M{"username": username}}}
	err := r.Collection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("Not found post")
		}
		return nil, err
	}
	return post, err
}
