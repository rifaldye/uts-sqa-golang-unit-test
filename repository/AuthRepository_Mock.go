package repository

import (
	"fmt"
	"uts-sqa/entity"

	"github.com/stretchr/testify/mock"
)

type AuthRepositoryMock struct {
	Mock mock.Mock
}

func (repository *AuthRepositoryMock) Register(request *entity.User) error {
	//disini cek kalo data yang dibalikin dari test adalah null maka mengembalikan null, dan sebaliknya
	arguments := repository.Mock.Called(request)
	if arguments.Get(0) == nil {
		return nil
	}else{
		err := arguments.Get(0).(error)	
		return err
	}
}
func (repository *AuthRepositoryMock) CheckEmail(email string) int {
	//disini dicek, kalo gk ada yangdibalikin dari test maka return 1 (email sudah digunakan)
	arguments := repository.Mock.Called(email)
	if arguments.Get(0) == nil {
		return 1
	}else{
		data := arguments.Get(0).(int)	
		return data
	}
}

func (repository *AuthRepositoryMock) Login(request *entity.Login) (*entity.User, error) {
	arguments := repository.Mock.Called(request)
	//disini dicek kalo gk ada data user yang dibalikin dari test, maka gagal login dan mengembalikan error
	if arguments.Get(0) == nil {
		return nil,fmt.Errorf("error")
	}else{
		data := arguments.Get(0).(entity.User)	
		return &data,nil
	}
}
