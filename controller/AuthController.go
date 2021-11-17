package controller

import (
	"net/http"
	"uts-sqa/entity"
	"uts-sqa/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service *service.AuthService
}

func (controller *AuthController) Register(c *gin.Context){
	var user entity.User
	err := c.BindJSON(&user)
	if err !=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message" :"gagal bind data"})
		return
	}
	err = controller.Service.Register(&user)
	if err!=nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message" :"email sudah ada/data tidak sesuai"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message": "berhasil daftar"})
}

func (controller *AuthController) Login(c *gin.Context){
	var login entity.Login
	err := c.BindJSON(&login)
	if err !=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message" :"gagal bind data"})
		return
	}
	token,err := controller.Service.Login(&login)
	if err!=nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message" :"username/password salah atau data tidak sesuai"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message": "berhasil login","token":token})
}

func (controller *AuthController) Validation(c *gin.Context){
	token := c.Param("token")

	err := controller.Service.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message" :"token tidak valid"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message": "token valid"})
}