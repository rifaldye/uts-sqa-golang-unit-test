package test

import (
	"fmt"
	"testing"
	"uts-sqa/entity"
	"uts-sqa/repository"
	"uts-sqa/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//inisialisasi fungsi menggunakan mock/stubs
var AuthRepository = &repository.AuthRepositoryMock{Mock: mock.Mock{}}
var AuthService = service.AuthService{
	Repository: AuthRepository,
	SecretKey: "ini-secret-key-sangat-secret",
}
func TestServiceVerify(t *testing.T){
	//index 0 data valid
	//index 1 perubahan 1 karakter pada header hash
	//index 2 perubahan 1 karakter pada payload hash
	//index 3 perubahan 1 karakter pada signature hash
	//index 4 modifikasi payload dalam jwt
	var test_input = [5]string{
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJpZmFsZHlAZ21haWwuY29tIiwibmFtZSI6InJpZmFsZHkgZWxuaW5vcnUiLCJwYXNzd29yZCI6InJpZmFsZGkxMjMifQ.jXiVeYeKjqKnbDeQcdCX0ZSJvPgH0uWNwTR87lpAIhc",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ92.eyJlbWFpbCI6InJpZmFsZHlAZ21haWwuY29tIiwibmFtZSI6InJpZmFsZHkgZWxuaW5vcnUiLCJwYXNzd29yZCI6InJpZmFsZGkxMjMifQ.jXiVeYeKjqKnbDeQcdCX0ZSJvPgH0uWNwTR87lpAIhc",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eysJlbWFpbCI6InJpZmFsZHlAZ21haWwuY29tIiwibmFtZSI6InJpZmFsZHkgZWxuaW5vcnUiLCJwYXNzd29yZCI6InJpZmFsZGkxMjMifQ.jXiVeYeKjqKnbDeQcdCX0ZSJvPgH0uWNwTR87lpAIhc",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJpZmFsZHlAZ21haWwuY29tIiwibmFtZSI6InJpZmFsZHkgZWxuaW5vcnUiLCJwYXNzd29yZCI6InJpZmFsZGkxMjMifQ.jXiVeYeKjqKnbDeQcdCX0ZSJvPgH0uWNwTR87lpAIshc",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJpZmFsZHlAZ21haWwuY29tIiwibmFtZSI6InJpZmFsZHkgZWxuaW5vcnUiLCJwYXNzd29yZCI6InJpZmFsZGkxMjMiLCJhc2FsIjoiamFrYXJ0YSJ9.3dVnToNQrDtG-Pjf9KeeKfZmZ7SYtX45n3n-dqZRdgY",
	}
	// test kalo string bukan jwt
	t.Run("is not JWT", func(t *testing.T){
		assert.NotNil(t,AuthService.VerifyToken("test"))
	})
	// test kalo string adalah jwt yang valid
	t.Run("valid JWT", func(t *testing.T){
		assert.Nil(t,AuthService.VerifyToken(test_input[0]))
	})
	// test kalo kode jwt valid diubah oleh user
	t.Run("invalid JWT Format", func(t *testing.T){
		assert.NotNil(t,AuthService.VerifyToken(test_input[1]))
		assert.NotNil(t,AuthService.VerifyToken(test_input[2]))
		assert.NotNil(t,AuthService.VerifyToken(test_input[3]))
		assert.NotNil(t,AuthService.VerifyToken(test_input[4]))
	})
}

func TestServiceRegister(t *testing.T){
	// gimana kalo ada data yang kosong
	t.Run("empty data", func(t *testing.T){
		//kalo email null
		data1 :=  entity.User{
			Name: "Rifaldy Elninoru",
			Password: "test",
		}
		assert.NotNil(t,AuthService.Register(&data1))
		//kalo nama null
		data2 :=  entity.User{
			Email: "Rifaldy@test.com",
			Password: "test",
		}
		assert.NotNil(t,AuthService.Register(&data2))
		//kalo password null
		data3 :=  entity.User{
			Email: "Rifaldy@test.com",
			Name: "Rifaldy Elninoru",
		}
		assert.NotNil(t,AuthService.Register(&data3))
		//kalau email isinya string kosong
		data4 :=  entity.User{
			Email: "",
			Name: "Rifaldy Elninoru",
			Password: "test123",
		}
		assert.NotNil(t,AuthService.Register(&data4))
		//kalau name isinya string kosong
		data5 :=  entity.User{
			Email: "test@test.com",
			Name: "",
			Password: "test123",
		}
		//kalau password isinya string kosong
		assert.NotNil(t,AuthService.Register(&data5))
		data6 :=  entity.User{
			Email: "test@test.com",
			Name: "",
			Password: "test123",
		}
		assert.NotNil(t,AuthService.Register(&data6))
	})
	// gimana kalo format email salah
	t.Run("wrong email format", func(t *testing.T){
		data :=  entity.User{
			Email: "test",
			Name: "Rifaldy Elninoru",
			Password: "test",
		}
		assert.NotNil(t,AuthService.Register(&data))
	})
	// gimana kalo email sudah ada sebelumnya
	t.Run("email already exist", func(t *testing.T){
		data :=  entity.User{
			Email: "test@test.com",
			Name: "Rifaldy Elninoru",
			Password: "test",
		}
		AuthRepository.Mock.On("CheckEmail",data.Email).Return(1)
		assert.NotNil(t,AuthService.Register(&data))
	})
	// gimana kalo data valid
	t.Run("valid", func(t *testing.T){
		data :=  entity.User{
			Email: "tester@test.com",
			Name: "Rifaldy Elninoru",
			Password: "test",
		}
		AuthRepository.Mock.On("CheckEmail",data.Email).Return(0)
		AuthRepository.Mock.On("Register",&data).Return(nil)
		assert.Nil(t,AuthService.Register(&data))
	})
}

func TestServiceLogin(t *testing.T){
	//gimana kalo ada form yang kosong
	t.Run("empty data", func(t *testing.T){
		//kalau email nul
		data1 := entity.Login{
			Password: "test",
		}
		token,err := AuthService.Login(&data1)
		assert.NotNil(t,err)
		assert.Equal(t,token,"")
		//kalau password null
		data2 := entity.Login{
			Email: "test@test.com",
		}
		token,err = AuthService.Login(&data2)
		assert.NotNil(t,err)
		assert.Equal(t,token,"")
		//kalau email isinya string kosong
		data3 := entity.Login{
			Email: "",
			Password: "test",
		}
		token,err = AuthService.Login(&data3)
		assert.NotNil(t,err)
		assert.Equal(t,token,"")
		//kalo password isinya string kosong
		data4 := entity.Login{
			Email: "",
			Password: "test",
		}
		token,err = AuthService.Login(&data4)
		assert.NotNil(t,err)
		assert.Equal(t,token,"")
		
	})
	//gimana kalo format email salah
	t.Run("wrong email format", func(t *testing.T){
		data :=  entity.User{
			Email: "test",
			Name: "Rifaldy Elninoru",
			Password: "test",
		}
		assert.NotNil(t,AuthService.Register(&data))
	})
	//gimana kalo email/password salah
	t.Run("email or password wrong", func(t *testing.T){
		data :=  entity.Login{
			Email: "test@test.com",
			Password: "test",
		}
		AuthRepository.Mock.On("Login",&data).Return(nil,fmt.Errorf("error"))
		token,err := AuthService.Login(&data)
		assert.NotNil(t,err)
		assert.Equal(t,token,"")
	})
	//gimana kalo email/password sbenar dan user melakukan verify
	t.Run("Login valid data and verify token", func(t *testing.T){
		data :=  entity.User{
			Email: "test2@test.com",
			Name: "Rifaldy Elninoru",
			Password: "test",
		}
		dataLogin := entity.Login{
			Email: data.Email,
			Password: data.Password,
		}
		AuthRepository.Mock.On("Login",&dataLogin).Return(data,nil)
		
		token,err := AuthService.Login(&dataLogin)
		assert.Nil(t,err)
		assert.NotNil(t,token)

		err = AuthService.VerifyToken(token)
		assert.Nil(t,err)
		
	})

}