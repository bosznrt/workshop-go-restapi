package user

import (
	"context"
	"server/db"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository description
type UserRepository struct {
	resource   *db.Resource
	collection *mongo.Collection
}

// Repository description
type Repository interface {
	GetAllUsers() (Users, error)
	AddUser(RegisterUser) (User, error)
	GetUserByID(id string) (*User, error)
	VerifyUser(email string, password string) (*User, error)
}

// NewUserRepository create repository
func NewUserRepository(resource *db.Resource) Repository {
	collection := resource.DB.Collection("user")
	repository := &UserRepository{resource: resource, collection: collection}
	return repository
}

// GetAllUsers list all users in collection
func (userRepo UserRepository) GetAllUsers() (Users, error) {
	users := Users{}
	ctx, cancel := initContext()
	defer cancel()

	cursor, err := userRepo.collection.Find(ctx, bson.M{})
	if err != nil {
		return Users{}, err
	}

	for cursor.Next(ctx) {
		var user User
		err = cursor.Decode(&user)
		if err != nil {
			logrus.Print(err)
		}
		users = append(users, user)
	}
	return users, nil
}

// VerifyUser to get user by email
func (userRepo *UserRepository) VerifyUser(email string, password string) (*User, error) {
	user := User{}

	ctx, cancel := initContext()
	defer cancel()

	err := userRepo.collection.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userRepo *UserRepository) GetUserByID(id string) (*User, error) {
	user := User{}

	ctx, cancel := initContext()
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	err := userRepo.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// AddUser create new user into db
func (userRepo UserRepository) AddUser(userReq RegisterUser) (User, error) {
	user := User{
		Id:    primitive.NewObjectID(),
		Email: userReq.Email,
		Name:  userReq.Name,
		Age:   userReq.Age,
	}

	ctx, cancel := initContext()
	defer cancel()
	_, err := userRepo.collection.InsertOne(ctx, user)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}
