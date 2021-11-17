package repository

import "uts-sqa/entity"

type AuthInterface interface {
	Register(request *entity.User) error
	CheckEmail(email string) int
	Login(*entity.Login) (*entity.User,error)
}