package routers

import (
	"mwserver/internal/auth"
	"mwserver/internal/config"
	"mwserver/internal/controllers"

	"github.com/gin-gonic/gin"
)

type wayTagRouter struct {
	wayTagController *controllers.WayTagController
	config           *config.Config
}

func newWayTagRouter(wayTagController *controllers.WayTagController, config *config.Config) *wayTagRouter {
	return &wayTagRouter{wayTagController, config}
}

func (wr *wayTagRouter) setWayTagRoutes(rg *gin.RouterGroup) {
	router := rg.Group("wayTags")
	router.POST("", auth.AuthMiddleware(wr.config), wr.wayTagController.AddWayTagToWay)
	router.DELETE("/:wayTagId/:wayId", auth.AuthMiddleware(wr.config), wr.wayTagController.DeleteWayTagFromWayByTagId)
}
