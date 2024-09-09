package routers

import (
	"mwserver/internal/auth"
	"mwserver/internal/config"
	"mwserver/internal/controllers"

	"github.com/gin-gonic/gin"
)

type userRouter struct {
	userController *controllers.UserController
	config         *config.Config
}

func newUserRouter(userController *controllers.UserController, config *config.Config) *userRouter {
	return &userRouter{userController, config}
}

func (ur *userRouter) setUserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("users")
	router.GET("", ur.userController.GetAllUsers)
	router.GET("/:userId", ur.userController.GetUserById)
	router.PATCH("/:userId", auth.AuthMiddleware(ur.config), ur.userController.UpdateUser)

	router.GET("/list-by-ids", auth.AuthMiddleware(ur.config), ur.userController.GetUsersByIDs)
}
