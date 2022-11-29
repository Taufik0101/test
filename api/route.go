package api

import (
	"github.com/gin-gonic/gin"
	"test/api/controller"
)

type Route interface {
	Routes(r *gin.Engine)
}

type route struct {
	userController controller.UserController
}

func (r2 route) Routes(r *gin.Engine) {
	//TODO implement me
	user := r.Group("/user")
	{
		user.GET("/", r2.userController.GetUser)
	}
}

func NewRoute(userControllers controller.UserController) Route {
	return &route{userController: userControllers}
}
