package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test/api/dto"
	"test/api/model"
	"test/api/service"
	"test/api/utils"
)

type BookController interface {
	GetBook(ctx *gin.Context)
	CreateBook(ctx *gin.Context)
	PutBook(ctx *gin.Context)
	PatchBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func (b bookController) GetBook(ctx *gin.Context) {
	//TODO implement me
	var books []*model.Book
	books, dataFrom := b.bookService.GetBook()
	resp := utils.BuildResponse(true, "Get Data User Berhasil dari "+dataFrom, books)
	ctx.JSON(http.StatusOK, resp)
}

func (b bookController) CreateBook(ctx *gin.Context) {
	//TODO implement me
	var DTOBook dto.CreateBook
	errCreate := ctx.ShouldBind(&DTOBook)
	if errCreate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errCreate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := b.bookService.CreateBook(DTOBook)
		if err != nil {
			response := utils.BuildErrorResponse("Tambah Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Tambah Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (b bookController) PutBook(ctx *gin.Context) {
	//TODO implement me
	var DTOPutBook dto.PutBook
	errUpdate := ctx.ShouldBind(&DTOPutBook)
	if errUpdate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errUpdate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := b.bookService.PutBook(DTOPutBook)
		if err != nil {
			response := utils.BuildErrorResponse("Update Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Update Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (b bookController) PatchBook(ctx *gin.Context) {
	//TODO implement me
	var DTOPatchBook dto.PatchBook
	errUpdate := ctx.ShouldBind(&DTOPatchBook)
	if errUpdate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errUpdate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := b.bookService.PatchBook(DTOPatchBook)
		if err != nil {
			response := utils.BuildErrorResponse("Update Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Update Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (b bookController) DeleteBook(ctx *gin.Context) {
	//TODO implement me
	id := ctx.PostForm("id")
	res, err := b.bookService.DeleteBook(id)
	if !res || err != nil {
		response := utils.BuildErrorResponse("Failed to delete", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	resp := utils.BuildResponse(true, "Delete Data Berhasil", res)
	ctx.JSON(http.StatusOK, resp)
}

func NewBookController(bookServices service.BookService) BookController {
	return &bookController{bookService: bookServices}
}
