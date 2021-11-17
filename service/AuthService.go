package service

import (
	"fmt"
	"uts-sqa/entity"
	"uts-sqa/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

type AuthService struct {
	Repository repository.AuthInterface
	SecretKey string
}

func (service *AuthService) Register(user *entity.User)error{
	validate := validator.New()
	err := validate.Struct(user)
	if err !=nil{
		return err
	}
	cek := service.Repository.CheckEmail(user.Email) 
	if cek != 0{
		return fmt.Errorf("email sudah digunakan")
	}
	err = service.Repository.Register(user)
	return err
}

func (service *AuthService) Login(request *entity.Login)(string,error){
	validate := validator.New()
	err := validate.Struct(request)
	if err !=nil{
		return "",err
	}
	user,err := service.Repository.Login(request)
	if err != nil {
		return "",err
	}

	token :=  jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["password"] = user.Password

	tokenString,err := token.SignedString([]byte(service.SecretKey))
	
	if err !=nil{
		fmt.Println("error generate jwt",err)

		return "",err
	}
	return tokenString,nil


}

func (service *AuthService) VerifyToken(tokens string)error{
	_, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
    if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Signing method invalid")
    } else if method != jwt.SigningMethodHS256 {
        return nil, fmt.Errorf("Signing method invalid")
    }

    return []byte(service.SecretKey), nil
})

if err != nil{
	return err
}
return nil
}
