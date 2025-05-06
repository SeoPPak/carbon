package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yoonaji/carbon_test/controllers"
	//"github.com/yoonaji/carbon/middleware"
)

type AuthRouteController struct {
	AuthController controllers.AuthController
}

func NewRouteAuthController(AuthController controllers.AuthController) AuthRouteController {
	return AuthRouteController{AuthController}
}

func (pc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", controllers.Logout)
		auth.POST("/refresh", controllers.Refresh)
	}
}