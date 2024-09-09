package controllers

import (
	"net/http"

	"mwserver/internal/schemas"
	"mwserver/internal/services"
	"mwserver/pkg/util"

	"github.com/gin-gonic/gin"
)

type MetricController struct {
	metricService *services.MetricService
	wayService    *services.WayService
}

func NewMetricController(metricService *services.MetricService, wayService *services.WayService) *MetricController {
	return &MetricController{metricService, wayService}
}

// Create Metric  handler
// @Summary Create a new metric
// @Description
// @Tags metric
// @ID create-metric
// @Accept  json
// @Produce  json
// @Param request body schemas.CreateMetricPayload true "query params"
// @Success 200 {object} schemas.MetricResponse
// @Router /metrics [post]
func (mc *MetricController) CreateMetric(ctx *gin.Context) {
	var payload *schemas.CreateMetricPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	metric, err := mc.metricService.CreateMetric(ctx, payload)
	util.HandleErrorGin(ctx, err)

	err = mc.wayService.UpdateWayIsCompletedStatus(ctx, metric.WayID)
	util.HandleErrorGin(ctx, err)

	ctx.JSON(http.StatusOK, metric.MetricResponse)
}

// Update Metric handler
// @Summary Update metric by UUID
// @Description
// @Tags metric
// @ID update-metric
// @Accept  json
// @Produce  json
// @Param request body schemas.UpdateMetricPayload true "query params"
// @Param metricId path string true "metric UUID"
// @Success 200 {object} schemas.MetricResponse
// @Router /metrics/{metricId} [patch]
func (mc *MetricController) UpdateMetric(ctx *gin.Context) {
	var payload *schemas.UpdateMetricPayload
	metricID := ctx.Param("metricId")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Failed payload", "error": err.Error()})
		return
	}

	metric, err := mc.metricService.UpdateMetric(ctx, &services.UpdateMetricParams{
		MetricID:         metricID,
		Description:      payload.Description,
		IsDone:           payload.IsDone,
		MetricEstimation: payload.MetricEstimation,
	})
	util.HandleErrorGin(ctx, err)

	err = mc.wayService.UpdateWayIsCompletedStatus(ctx, metric.WayID)
	util.HandleErrorGin(ctx, err)

	ctx.JSON(http.StatusOK, metric.MetricResponse)
}

// Deleting Metric handlers
// @Summary Delete metric by UUID
// @Description
// @Tags metric
// @ID delete-metric
// @Accept  json
// @Produce  json
// @Param metricId path string true "metric ID"
// @Success 204
// @Router /metrics/{metricId} [delete]
func (mc *MetricController) DeleteMetricById(ctx *gin.Context) {
	metricID := ctx.Param("metricId")

	wayID, err := mc.metricService.DeleteMetricById(ctx, metricID)
	util.HandleErrorGin(ctx, err)

	err = mc.wayService.UpdateWayIsCompletedStatus(ctx, wayID)
	util.HandleErrorGin(ctx, err)

	ctx.Status(http.StatusNoContent)
}
