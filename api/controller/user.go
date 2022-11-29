package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/api/model"
	"test/api/service"
	"test/api/utils"
)

type UserController interface {
	GetUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func (u userController) GetUser(ctx *gin.Context) {
	//TODO implement me
	var users []model.User
	users, dataFrom := u.userService.GetUser()
	resp := utils.BuildResponse(true, "Get Data User Berhasil dari "+dataFrom, users)
	ctx.JSON(http.StatusOK, resp)
}

func NewUserController(userServices service.UserService) UserController {
	return &userController{userService: userServices}
}
