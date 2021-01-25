package user

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model of user
type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
	Name     string             `bson:"name" json:"name"`
	Age      int                `bson:"age" json:"age"`
}

// type UserRequest struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

type RegisterUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required,gte=0,lte=130"`
}

// Users strcut for user list
type Users []User
