package repository

import (
	"context"
	"gogod/config"
	"gogod/domain"
	"gogod/model"
	"gogod/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	client *mongo.Client
}

func NewUserRepository(c *mongo.Client) domain.UserRepository {
	return &userRepository{
		client: c,
	}
}

func (r *userRepository) userCol() *mongo.Collection {
	return database.GetCollection(config.ENV, r.client, "users")
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
	result, err := r.userCol().InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""
	return user, nil
}

func (r *userRepository) GetByID(_id primitive.ObjectID, withPwd bool) (*model.User, error) {
	user := model.User{}
	filter := bson.D{{Key: "_id", Value: _id}}
	project := options.FindOne()
	if !withPwd {
		project.SetProjection(bson.M{"password": 0})
	}
	err := r.userCol().FindOne(context.TODO(), filter, project).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string, withPwd bool) (*model.User, error) {
	user := model.User{}
	filter := bson.D{{Key: "email", Value: email}}
	project := options.FindOne()
	if !withPwd {
		project.SetProjection(bson.M{"password": 0})
	}
	err := r.userCol().FindOne(context.TODO(), filter, project).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) All() ([]model.User, error) {
	filter := bson.D{}
	project := options.Find().SetProjection(bson.M{"password": 0})
	cursor, err := r.userCol().Find(context.TODO(), filter, project)
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	if err := cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
