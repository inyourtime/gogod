package repository

import (
	"context"
	"gogod/config"
	"gogod/domain"
	"gogod/model"
	"gogod/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type blogRepository struct {
	client *mongo.Client
}

func NewBlogRepository(c *mongo.Client) domain.BlogRepository {
	return &blogRepository{
		client: c,
	}
}

func (r *blogRepository) blogCol() *mongo.Collection {
	return database.GetCollection(config.ENV, r.client, "blogs")
}

func (r *blogRepository) Create(blog *model.Blog) error {
	_, err := r.blogCol().InsertOne(context.TODO(), blog)
	if err != nil {
		return err
	}
	return nil
}
func (r *blogRepository) All() ([]model.Blog, error) {
	filter := bson.D{}
	cursor, err := r.blogCol().Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	blogs := []model.Blog{}
	if err := cursor.All(context.TODO(), &blogs); err != nil {
		return nil, err
	}
	return blogs, nil
}
func (r *blogRepository) GetByID(blogID string) (*model.Blog, error) {
	blog := model.Blog{}
	filter := bson.D{{Key: "blogId", Value: blogID}}
	err := r.blogCol().FindOne(context.TODO(), filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &blog, nil
}
func (r *blogRepository) UpdateOne(blog *model.BlogUpdateRequest) error {
	filter := bson.D{{Key: "blogId", Value: blog.BlogID}}
	_, err := r.blogCol().UpdateOne(context.TODO(), filter, bson.M{"$set": blog})
	if err != nil {
		return err
	}
	return nil
}
func (r *blogRepository) Delete(blogID string) error {
	filter := bson.D{{Key: "blogId", Value: blogID}}
	_, err := r.blogCol().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
