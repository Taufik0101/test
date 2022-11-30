package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/api/dto"
	"test/api/model"
	"test/api/service"
	"test/api/utils"
)

type UserController interface {
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	PutUser(ctx *gin.Context)
	PatchUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func (u userController) PutUser(ctx *gin.Context) {
	//TODO implement me
	var DTOPutUser dto.PutUser
	errUpdate := ctx.ShouldBind(&DTOPutUser)
	if errUpdate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errUpdate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := u.userService.PutUser(DTOPutUser)
		if err != nil {
			response := utils.BuildErrorResponse("Update Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Update Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (u userController) PatchUser(ctx *gin.Context) {
	//TODO implement me
	var DTOPatchUser dto.PatchUser
	errUpdate := ctx.ShouldBind(&DTOPatchUser)
	if errUpdate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errUpdate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := u.userService.PatchUser(DTOPatchUser)
		if err != nil {
			response := utils.BuildErrorResponse("Update Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Update Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (u userController) DeleteUser(ctx *gin.Context) {
	//TODO implement me
	id := ctx.PostForm("id")
	res, err := u.userService.DeleteUser(id)
	if !res || err != nil {
		response := utils.BuildErrorResponse("Failed to delete", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	resp := utils.BuildResponse(true, "Delete Data Berhasil", res)
	ctx.JSON(http.StatusOK, resp)

}

func (u userController) CreateUser(ctx *gin.Context) {
	//TODO implement me
	var DTOUser dto.CreateUser
	errCreate := ctx.ShouldBind(&DTOUser)
	if errCreate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errCreate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := u.userService.CreateUser(DTOUser)
		if err != nil {
			response := utils.BuildErrorResponse("Tambah Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Tambah Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (u userController) GetUser(ctx *gin.Context) {
	//TODO implement me
	var users []*model.User
	users, dataFrom := u.userService.GetUser()
	resp := utils.BuildResponse(true, "Get Data User Berhasil dari "+dataFrom, users)
	ctx.JSON(http.StatusOK, resp)
}

func NewUserController(userServices service.UserService) UserController {
	return &userController{userService: userServices}
}
