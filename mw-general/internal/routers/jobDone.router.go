package routers

import (
	"mwserver/internal/auth"
	"mwserver/internal/config"
	"mwserver/internal/controllers"

	"github.com/gin-gonic/gin"
)

type jobDoneRouter struct {
	jobDoneController *controllers.JobDoneController
	config            *config.Config
}

func newJobDoneRouter(jobDoneController *controllers.JobDoneController, config *config.Config) *jobDoneRouter {
	return &jobDoneRouter{jobDoneController, config}
}

func (jr *jobDoneRouter) setJobDoneRoutes(rg *gin.RouterGroup) {
	router := rg.Group("jobDones")
	router.POST("", auth.AuthMiddleware(jr.config), jr.jobDoneController.CreateJobDone)
	router.PATCH("/:jobDoneId", auth.AuthMiddleware(jr.config), jr.jobDoneController.UpdateJobDone)
	router.DELETE("/:jobDoneId", auth.AuthMiddleware(jr.config), jr.jobDoneController.DeleteJobDoneById)
}
