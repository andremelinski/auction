package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/andremelinski/auction/config/logger"
	"github.com/andremelinski/auction/internal/entity"
	internalerror "github.com/andremelinski/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db.Collection("users"),
	}
}

func (u *UserRepository) FindUserById(ctx context.Context, id string) (*entity.User, *internalerror.InternalError) {
	userEntityMongo := &UserEntityMongo{}

	filter := bson.M{"_id": id}

	err := u.Collection.FindOne(ctx, filter).Decode(userEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("user not found", err)
			return nil, internalerror.NewNotFoundError(
				fmt.Sprintf("user not found with this id = %s", id),
			)
		}
		return nil, internalerror.NewNotFoundError(
			"error trying to find user by id",
		)
	}

	userEntity := &entity.User{Id: userEntityMongo.Id, Name: userEntityMongo.Name}

	return userEntity, nil
}
