package routers

import (
	"mwserver/internal/auth"
	"mwserver/internal/config"
	"mwserver/internal/controllers"

	"github.com/gin-gonic/gin"
)

type fromUserMentoringRequestRouter struct {
	fromUserMentoringRequestController *controllers.FromUserMentoringRequestController
	config                             *config.Config
}

func newFromUserMentoringRequestRouter(fromUserMentoringRequestController *controllers.FromUserMentoringRequestController, config *config.Config) *fromUserMentoringRequestRouter {
	return &fromUserMentoringRequestRouter{fromUserMentoringRequestController, config}
}

func (cr *fromUserMentoringRequestRouter) setFromUserMentoringRequestRoutes(rg *gin.RouterGroup) {
	router := rg.Group("fromUserMentoringRequests")
	router.POST("", auth.AuthMiddleware(cr.config), cr.fromUserMentoringRequestController.CreateFromUserMentoringRequest)
	router.DELETE("/:userUuid/:wayUuid", auth.AuthMiddleware(cr.config), cr.fromUserMentoringRequestController.DeleteFromUserMentoringRequestById)
}
