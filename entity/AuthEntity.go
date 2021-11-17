package entity

type User struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Name     string `json:"name" bson:"name" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type Login struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}
