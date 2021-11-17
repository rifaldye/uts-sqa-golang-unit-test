package main

import (
	"uts-sqa/config"
	"uts-sqa/controller"
	"uts-sqa/repository"
	"uts-sqa/service"

	"github.com/gin-gonic/gin"
)

func main(){
	db,_ := config.GetMongoDB()	
	repository := repository.AuthRepository{
		Db : db,
		Collection: "users",
	}
	service := service.AuthService{
		Repository: &repository,
		SecretKey: "ini-secret-key-sangat-secret",
	}
	controller := controller.AuthController{
		Service : &service,
	}
	server := gin.Default()
	server.POST("/user/register",controller.Register)
	server.POST("/user/login",controller.Login)
	server.GET("/user/verify/:token",controller.Validation)

	server.Run(":9000")
}