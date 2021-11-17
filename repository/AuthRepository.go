package repository

import (
	"uts-sqa/entity"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AuthRepository struct {
	Db         *mgo.Database
	Collection string
}

func (repository *AuthRepository) Register(request *entity.User) error{
	err := repository.Db.C(repository.Collection).Insert(request)
	return err
}
func (repository *AuthRepository) CheckEmail(email string) int{
	query := bson.M{"email":email}
	hasil,_ := repository.Db.C(repository.Collection).Find(query).Count()
	
	return hasil
}

func (repository *AuthRepository) Login(request *entity.Login) (*entity.User,error){
	var user entity.User
	err := repository.Db.C(repository.Collection).Find(request).One(&user)
	if err !=nil{
		return nil,err
	}
	return &user,nil
}
