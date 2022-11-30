package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/api/dto"
	"test/api/model"
	"test/api/service"
	"test/api/utils"
)

type BorrowController interface {
	GetBorrow(ctx *gin.Context)
	CreateBorrow(ctx *gin.Context)
	PutBorrow(ctx *gin.Context)
	PatchBorrow(ctx *gin.Context)
	DeleteBorrow(ctx *gin.Context)
}

type borrowController struct {
	borrowService service.BorrowService
}

func (b borrowController) GetBorrow(ctx *gin.Context) {
	//TODO implement me
	var borrowss []*model.Borrow
	borrowss, dataFrom := b.borrowService.GetBorrow()
	resp := utils.BuildResponse(true, "Get Data User Berhasil dari "+dataFrom, borrowss)
	ctx.JSON(http.StatusOK, resp)
}

func (b borrowController) CreateBorrow(ctx *gin.Context) {
	//TODO implement me
	var DTOBorrow dto.CreateBorrow
	errCreate := ctx.ShouldBind(&DTOBorrow)
	if errCreate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errCreate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := b.borrowService.CreateBorrow(DTOBorrow)
		if err != nil {
			response := utils.BuildErrorResponse("Tambah Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Tambah Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (b borrowController) PutBorrow(ctx *gin.Context) {
	//TODO implement me
	var DTOPutBorrow dto.PutBorrow
	errUpdate := ctx.ShouldBind(&DTOPutBorrow)
	if errUpdate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errUpdate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := b.borrowService.PutBorrow(DTOPutBorrow)
		if err != nil {
			response := utils.BuildErrorResponse("Update Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Update Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (b borrowController) PatchBorrow(ctx *gin.Context) {
	//TODO implement me
	var DTOPatchBorrow dto.PatchBorrow
	errUpdate := ctx.ShouldBind(&DTOPatchBorrow)
	if errUpdate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errUpdate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := b.borrowService.PatchBorrow(DTOPatchBorrow)
		if err != nil {
			response := utils.BuildErrorResponse("Update Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Update Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (b borrowController) DeleteBorrow(ctx *gin.Context) {
	//TODO implement me
	id := ctx.PostForm("id")
	res, err := b.borrowService.DeleteBorrow(id)
	if !res || err != nil {
		response := utils.BuildErrorResponse("Failed to delete", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	resp := utils.BuildResponse(true, "Delete Data Berhasil", res)
	ctx.JSON(http.StatusOK, resp)
}

func NewBorrowController(borrowServices service.BorrowService) BorrowController {
	return &borrowController{borrowService: borrowServices}
}
