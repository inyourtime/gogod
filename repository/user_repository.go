package repository

import (
	"context"
	"gogod/domain"
	"gogod/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	col *mongo.Collection
}

func NewUserRepository(c *mongo.Collection) domain.UserRepository {
	return &userRepository{
		col: c,
	}
}

func (r *userRepository) Create(user *model.User) (*model.User, error) {
	result, err := r.col.InsertOne(context.TODO(), user)
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
	err := r.col.FindOne(context.TODO(), filter, project).Decode(&user)
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
	err := r.col.FindOne(context.TODO(), filter, project).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
