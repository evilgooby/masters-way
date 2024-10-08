package routers

import (
	"context"
	openapiGeneral "mwserver/apiAutogenerated/general"
	"mwserver/internal/auth"
	"mwserver/internal/config"
	"mwserver/internal/openapi"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlanTag(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "7cdb041b-4574-4f7b-a500-c53e74c72e94"
	wayID := "550e8400-e29b-41d4-a716-446655440000"
	jobTagID := "c73ff20b-e64e-4e5f-b270-1a40ba1bd81b"

	t.Run("should create plan job tag successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}
		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)

		// Create Day Report for Way
		requestCreateDayReport := openapiGeneral.SchemasCreateDayReportPayload{WayUuid: wayID}
		dayReport, _, err := generalApi.DayReportAPI.CreateDayReport(ctx).Request(requestCreateDayReport).Execute()
		if err != nil {
			t.Fatalf("Failed to create day report: %v", err)
		}
		dayReportID := dayReport.CompositionParticipants[0].DayReportId

		requestCreatePlan := openapiGeneral.SchemasCreatePlanPayload{
			DayReportUuid: dayReportID,
			Description:   "Description",
			IsDone:        false,
			OwnerUuid:     user,
			Time:          0,
		}
		plan, _, err := generalApi.PlanAPI.CreatePlan(ctx).Request(requestCreatePlan).Execute()
		if err != nil {
			t.Fatalf("Failed to create plan: %v", err)
		}

		requestCreateJobTag := openapiGeneral.SchemasCreatePlanJobTagPayload{
			JobTagUuid: jobTagID,
			PlanUuid:   plan.Uuid,
		}
		response, err := generalApi.PlanJobTagAPI.CreatePlanJobTag(ctx).Request(requestCreateJobTag).Execute()
		if err != nil {
			t.Fatalf("Failed to create plan job tag: %v", err)
		}

		expectedData := []openapiGeneral.SchemasJobTagResponse{
			{
				Color:       "blue",
				Description: "this is not my tag",
				Name:        "coding",
				Uuid:        jobTagID,
			},
		}

		report, _, err := generalApi.DayReportAPI.GetDayReports(ctx, wayID).Execute()

		assert.Equal(t, http.StatusNoContent, response.StatusCode)

		isExists := false
		for _, dayReport := range report.DayReports {
			for _, p := range dayReport.Plans {
				for _, tag := range p.Tags {
					if tag == expectedData[0] {
						isExists = true
						break
					}
				}
			}
		}

		if !isExists {
			t.Fatalf("Failed to create day report")
		}
	})
}

func TestDeletePlanJobTagById(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "7cdb041b-4574-4f7b-a500-c53e74c72e94"
	planID := "18cbbee6-5071-4608-b349-ffad514711cb"
	jobTagID := "5ebb8d43-b685-4090-8453-ceaa7aad2095"
	wayID := "1d922e8a-5d58-4b82-9a3d-83e2e73b3f91"

	t.Run("should create plan job tag successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}
		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)

		response, err := generalApi.PlanJobTagAPI.DeletePlanJobTag(ctx, jobTagID, planID).Execute()
		if err != nil {
			t.Fatalf("Failed to delete day report: %v", err)
		}

		assert.Equal(t, http.StatusNoContent, response.StatusCode)

		dayReport, _, err := generalApi.DayReportAPI.GetDayReports(ctx, wayID).Execute()
		if err != nil {
			t.Fatalf("Failed to get day report: %v", err)
		}

		isExists := false
		for _, dayReport := range dayReport.DayReports {
			for _, p := range dayReport.Plans {
				for _, tag := range p.Tags {
					if tag.Uuid == jobTagID {
						isExists = true
						break
					}
				}
			}
		}

		if isExists {
			t.Fatalf("Failed to delete day report")
		}
	})
}
