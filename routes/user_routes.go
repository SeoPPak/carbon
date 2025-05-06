package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yoonaji/carbon_test/controllers"
	"github.com/yoonaji/carbon/middleware"
)

type UserRouteController struct {
	UserController controllers.UserController
}

func NewRouteUserController(UserController controllers.UserController) UserRouteController {
	return UserRouteController{UserController}
}

func (pc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.Use(middleware.DeserializeUser())
	{
		users.PUT("/:userId", controllers.UpdateUser)
	}
}