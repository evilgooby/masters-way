package routers

import (
	"mwserver/internal/auth"
	"mwserver/internal/config"
	"mwserver/internal/controllers"

	"github.com/gin-gonic/gin"
)

type mentorUserWayRouter struct {
	mentorUserWayController *controllers.MentorUserWayController
	config                  *config.Config
}

func newMentorUserWayRouter(mentorUserWayController *controllers.MentorUserWayController, config *config.Config) *mentorUserWayRouter {
	return &mentorUserWayRouter{mentorUserWayController, config}
}

func (mr *mentorUserWayRouter) setMentorUserWayRoutes(rg *gin.RouterGroup) {
	router := rg.Group("mentorUserWays")
	router.POST("", auth.AuthMiddleware(mr.config), mr.mentorUserWayController.AddMentorUserWay)
	router.DELETE("", auth.AuthMiddleware(mr.config), mr.mentorUserWayController.DeleteMentorUserWay)
}
