package services

import (
	"context"
	db "mwserver/internal/db/sqlc"
	"mwserver/internal/schemas"
	"mwserver/pkg/util"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type IMetricRepository interface {
	CreateMetric(ctx context.Context, arg db.CreateMetricParams) (db.Metric, error)
	UpdateMetric(ctx context.Context, arg db.UpdateMetricParams) (db.UpdateMetricRow, error)
	DeleteMetric(ctx context.Context, metricsUuid pgtype.UUID) (db.Metric, error)
}

type MetricService struct {
	metricRepository IMetricRepository
}

func NewMetricService(metricRepository IMetricRepository) *MetricService {
	return &MetricService{metricRepository}
}

type MetricResult struct {
	MetricResponse *schemas.MetricResponse
	WayID          string
}

func (ms *MetricService) CreateMetric(ctx context.Context, payload *schemas.CreateMetricPayload) (*MetricResult, error) {
	now := time.Now()
	parsedTime, err := time.Parse(util.DEFAULT_STRING_LAYOUT, payload.DoneDate)
	args := db.CreateMetricParams{
		Description:      payload.Description,
		IsDone:           payload.IsDone,
		DoneDate:         pgtype.Timestamp{Time: parsedTime, Valid: err != nil},
		MetricEstimation: int32(payload.MetricEstimation),
		WayUuid:          pgtype.UUID{Bytes: uuid.MustParse(payload.WayUuid), Valid: true},
		UpdatedAt:        pgtype.Timestamp{Time: now, Valid: true},
	}

	metric, err := ms.metricRepository.CreateMetric(ctx, args)
	if err != nil {
		return nil, err
	}

	return &MetricResult{
		MetricResponse: &schemas.MetricResponse{
			Uuid:             util.ConvertPgUUIDToUUID(metric.Uuid).String(),
			Description:      metric.Description,
			IsDone:           metric.IsDone,
			DoneDate:         util.MarshalPgTimestamp(metric.DoneDate),
			MetricEstimation: metric.MetricEstimation,
		},
		WayID: util.ConvertPgUUIDToUUID(metric.WayUuid).String(),
	}, nil
}

type UpdateMetricParams struct {
	MetricID         string
	Description      *string
	IsDone           *bool
	MetricEstimation *int32
}

func (ms *MetricService) UpdateMetric(ctx context.Context, params *UpdateMetricParams) (*MetricResult, error) {
	now := time.Now()

	var isDonePg pgtype.Bool
	var doneDatePg pgtype.Timestamp
	if params.IsDone != nil {
		isDonePg = pgtype.Bool{Bool: *params.IsDone, Valid: true}
		doneDatePg = pgtype.Timestamp{Time: now, Valid: *params.IsDone}
	}

	var descriptionPg pgtype.Text
	if params.Description != nil {
		descriptionPg = pgtype.Text{String: *params.Description, Valid: true}
	}

	var metricEstimationPg pgtype.Int4
	if params.MetricEstimation != nil {
		metricEstimationPg = pgtype.Int4{Int32: *params.MetricEstimation, Valid: true}
	}

	args := db.UpdateMetricParams{
		Uuid:             pgtype.UUID{Bytes: uuid.MustParse(params.MetricID), Valid: true},
		UpdatedAt:        pgtype.Timestamp{Time: now, Valid: true},
		Description:      descriptionPg,
		IsDone:           isDonePg,
		DoneDate:         doneDatePg,
		MetricEstimation: metricEstimationPg,
	}

	metric, err := ms.metricRepository.UpdateMetric(ctx, args)
	if err != nil {
		return nil, err
	}

	return &MetricResult{
		MetricResponse: &schemas.MetricResponse{
			Uuid:             util.ConvertPgUUIDToUUID(metric.Uuid).String(),
			Description:      metric.Description,
			IsDone:           metric.IsDone,
			DoneDate:         util.MarshalPgTimestamp(metric.DoneDate),
			MetricEstimation: metric.MetricEstimation,
		},
		WayID: util.ConvertPgUUIDToUUID(metric.WayUuid).String(),
	}, nil
}

func (ms *MetricService) DeleteMetricById(ctx context.Context, metricID string) (string, error) {
	removedMetric, err := ms.metricRepository.DeleteMetric(ctx, pgtype.UUID{Bytes: uuid.MustParse(metricID), Valid: true})
	if err != nil {
		return "", err
	}

	return util.ConvertPgUUIDToUUID(removedMetric.WayUuid).String(), nil
}
