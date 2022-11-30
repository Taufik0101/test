package api

import (
	"github.com/gin-gonic/gin"
	"test/api/controller"
)

type Route interface {
	Routes(r *gin.Engine)
}

type route struct {
	userController   controller.UserController
	bookController   controller.BookController
	borrowController controller.BorrowController
}

func (r2 route) Routes(r *gin.Engine) {
	//TODO implement me
	user := r.Group("/user")
	{
		user.GET("/", r2.userController.GetUser)
		user.POST("/add", r2.userController.CreateUser)
		user.PUT("/put", r2.userController.PutUser)
		user.PATCH("/patch", r2.userController.PatchUser)
		user.DELETE("/delete", r2.userController.DeleteUser)
	}

	book := r.Group("/book")
	{
		book.GET("/", r2.bookController.GetBook)
		book.POST("/add", r2.bookController.CreateBook)
		book.PUT("/put", r2.bookController.PutBook)
		book.PATCH("/patch", r2.bookController.PatchBook)
		book.DELETE("/delete", r2.bookController.DeleteBook)
	}

	borrow := r.Group("/borrow")
	{
		borrow.GET("/", r2.borrowController.GetBorrow)
		borrow.POST("/add", r2.borrowController.CreateBorrow)
		borrow.PUT("/put", r2.borrowController.PutBorrow)
		borrow.PATCH("/patch", r2.borrowController.PatchBorrow)
		borrow.DELETE("/delete", r2.borrowController.DeleteBorrow)
	}
}

func NewRoute(
	userControllers controller.UserController,
	bookControllers controller.BookController,
	borrowControllers controller.BorrowController,
) Route {
	return &route{
		userController:   userControllers,
		bookController:   bookControllers,
		borrowController: borrowControllers,
	}
}
