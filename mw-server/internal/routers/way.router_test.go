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

func TestGetAllWays(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"

	t.Run("should get all ways successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		ways, response, err := generalApi.WayAPI.GetAllWays(ctx).Page(1).Status("all").MinDayReportsAmount(0).Execute()
		if err != nil {
			t.Fatalf("Failed to get all ways: %v", err)
		}

		copiedFromWayUuid := openapiGeneral.NullableString{}
		copiedFromWayUuid.Set(nil)

		expectedData := &openapiGeneral.SchemasGetAllWaysResponse{
			Size: 14,
			Ways: []openapiGeneral.SchemasWayPlainResponse{
				{
					ChildrenUuids:     []string{},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-08-10T00:00:00.000Z",
					DayReportsAmount:  0,
					EstimationTime:    40404040,
					FavoriteForUsers:  0,
					GoalDescription:   "dana evans goal 1",
					IsCompleted:       false,
					IsPrivate:         false,
					Mentors:           []openapiGeneral.SchemasUserPlainResponse{},
					MetricsDone:       0,
					MetricsTotal:      0,
					Name:              "dana evans way 1",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-06T05:00:00.000Z",
						Description: "A brief description about Dana.",
						Email:       "dana.evans@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
						IsMentor:    true,
						Name:        "Dana Evans",
						Uuid:        "1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1",
					},
					UpdatedAt: "2024-08-10T00:00:00.000Z",
					Uuid:      "a2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2",
					WayTags: []openapiGeneral.SchemasWayTagResponse{
						{
							Name: "some tag",
							Uuid: "46d5dd00-75fc-4563-9616-5252a6fa05d3",
						},
					},
				},
				{
					ChildrenUuids:     []string{},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-08-10T00:00:00.000Z",
					DayReportsAmount:  0,
					EstimationTime:    40404040,
					FavoriteForUsers:  0,
					GoalDescription:   "dana evans goal 2",
					IsCompleted:       false,
					IsPrivate:         false,
					Mentors:           []openapiGeneral.SchemasUserPlainResponse{},
					MetricsDone:       0,
					MetricsTotal:      0,
					Name:              "dana evans way 2",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-06T05:00:00.000Z",
						Description: "A brief description about Dana.",
						Email:       "dana.evans@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
						IsMentor:    true,
						Name:        "Dana Evans",
						Uuid:        "1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1",
					},
					UpdatedAt: "2024-08-10T00:00:00.000Z",
					Uuid:      "aa13eee9-7dca-46ed-a3f7-21d8b7ae3b72",
					WayTags:   []openapiGeneral.SchemasWayTagResponse{},
				},
				{
					ChildrenUuids:     []string{},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-08-10T00:00:00.000Z",
					DayReportsAmount:  2,
					EstimationTime:    40404040,
					FavoriteForUsers:  0,
					GoalDescription:   "dana evans goal 3",
					IsCompleted:       true,
					IsPrivate:         false,
					Mentors:           []openapiGeneral.SchemasUserPlainResponse{},
					MetricsDone:       0,
					MetricsTotal:      0,
					Name:              "dana evans way 3",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-06T05:00:00.000Z",
						Description: "A brief description about Dana.",
						Email:       "dana.evans@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
						IsMentor:    true,
						Name:        "Dana Evans",
						Uuid:        "1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1",
					},
					UpdatedAt: "2024-08-10T00:00:00.000Z",
					Uuid:      "9230479a-a481-4f83-b770-138ef4f3139c",
					WayTags:   []openapiGeneral.SchemasWayTagResponse{},
				},
				{
					ChildrenUuids:     []string{},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-07-22T00:00:00.000Z",
					DayReportsAmount:  0,
					EstimationTime:    30303030,
					FavoriteForUsers:  0,
					GoalDescription:   "alice johnson goal 1",
					IsCompleted:       false,
					IsPrivate:         false,
					Mentors:           []openapiGeneral.SchemasUserPlainResponse{},
					MetricsDone:       0,
					MetricsTotal:      0,
					Name:              "alice johnson 1",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-08T05:10:00.000Z",
						Description: "A brief description about Alice.",
						Email:       "alice.johnson@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
						IsMentor:    true,
						Name:        "Alice Johnson",
						Uuid:        "3d922e8a-5d58-4b82-9a3d-83e2e73b3f91",
					},
					UpdatedAt: "2024-07-22T00:00:00.000Z",
					Uuid:      "5cc724a0-383f-45ad-99a1-8514f51717f2",
					WayTags:   []openapiGeneral.SchemasWayTagResponse{},
				},
				{
					ChildrenUuids:     []string{},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-07-22T00:00:00.000Z",
					DayReportsAmount:  0,
					EstimationTime:    30303030,
					FavoriteForUsers:  0,
					GoalDescription:   "alice johnson goal 2",
					IsCompleted:       false,
					IsPrivate:         false,
					Mentors:           []openapiGeneral.SchemasUserPlainResponse{},
					MetricsDone:       0,
					MetricsTotal:      0,
					Name:              "alice johnson 2",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-08T05:10:00.000Z",
						Description: "A brief description about Alice.",
						Email:       "alice.johnson@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
						IsMentor:    true,
						Name:        "Alice Johnson",
						Uuid:        "3d922e8a-5d58-4b82-9a3d-83e2e73b3f91",
					},
					UpdatedAt: "2024-07-22T00:00:00.000Z",
					Uuid:      "78f86c77-2018-4511-90dc-d96df77f496a",
					WayTags:   []openapiGeneral.SchemasWayTagResponse{},
				},
				{
					ChildrenUuids: []string{
						"550e8400-e29b-41d4-a716-446655440000",
					},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-07-09T00:00:00.000Z",
					DayReportsAmount:  0,
					EstimationTime:    20202020,
					FavoriteForUsers:  1,
					GoalDescription:   "jane smith goal",
					IsCompleted:       true,
					IsPrivate:         false,
					Mentors: []openapiGeneral.SchemasUserPlainResponse{
						{
							CreatedAt:   "2024-07-08T05:10:00.000Z",
							Description: "A brief description about Alice.",
							Email:       "alice.johnson@example.com",
							ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
							IsMentor:    true,
							Name:        "Alice Johnson",
							Uuid:        "3d922e8a-5d58-4b82-9a3d-83e2e73b3f91",
						},
					},
					MetricsDone:  0,
					MetricsTotal: 0,
					Name:         "jane smith way",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-08T05:50:00.000Z",
						Description: "A brief description about Jane.",
						Email:       "jane.smith@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
						IsMentor:    true,
						Name:        "Jane Smith",
						Uuid:        "8e77b89d-57c4-4b7f-8cd4-8dfc6bcb7d1b",
					},
					UpdatedAt: "2024-07-09T00:00:00.000Z",
					Uuid:      "9e77b89d-57c4-4b7f-8cd4-8dfc6bcb7d1b",
					WayTags:   []openapiGeneral.SchemasWayTagResponse{},
				},
				{
					ChildrenUuids:     []string{},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-07-09T00:00:00.000Z",
					DayReportsAmount:  0,
					EstimationTime:    20202020,
					FavoriteForUsers:  0,
					GoalDescription:   "jane smith goal 1",
					IsCompleted:       false,
					IsPrivate:         false,
					Mentors: []openapiGeneral.SchemasUserPlainResponse{
						{
							CreatedAt:   "2024-07-08T05:50:00.000Z",
							Description: "A brief description about Jane.",
							Email:       "jane.smith@example.com",
							ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
							IsMentor:    true,
							Name:        "Jane Smith",
							Uuid:        "8e77b89d-57c4-4b7f-8cd4-8dfc6bcb7d1b",
						},
					},
					MetricsDone:  0,
					MetricsTotal: 0,
					Name:         "jane smith way 1",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-08T05:50:00.000Z",
						Description: "A brief description about Jane.",
						Email:       "jane.smith@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
						IsMentor:    true,
						Name:        "Jane Smith",
						Uuid:        "8e77b89d-57c4-4b7f-8cd4-8dfc6bcb7d1b",
					},
					UpdatedAt: "2024-07-09T00:00:00.000Z",
					Uuid:      "dce03ca6-f626-4c33-a44b-5a1b4ff62aa7",
					WayTags:   []openapiGeneral.SchemasWayTagResponse{},
				},
				{
					ChildrenUuids:     []string{},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-07-09T00:00:00.000Z",
					DayReportsAmount:  0,
					EstimationTime:    40404040,
					FavoriteForUsers:  0,
					GoalDescription:   "ronnie stanton goal",
					IsCompleted:       false,
					IsPrivate:         false,
					Mentors:           []openapiGeneral.SchemasUserPlainResponse{},
					MetricsDone:       0,
					MetricsTotal:      0,
					Name:              "ronnie stanton way",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-06T10:00:00.000Z",
						Description: "A brief description about Ronnie.",
						Email:       "ronnie.stanton@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAN",
						IsMentor:    false,
						Name:        "Ronnie Stanton",
						Uuid:        "d63d2f89-6412-4324-8587-7061bf02dca4",
					},
					UpdatedAt: "2024-07-09T00:00:00.000Z",
					Uuid:      "9972552a-c0b3-41f3-b464-284d36a36964",
					WayTags:   []openapiGeneral.SchemasWayTagResponse{},
				},
				{
					ChildrenUuids:     []string{},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-07-09T00:00:00.000Z",
					DayReportsAmount:  0,
					EstimationTime:    40404040,
					FavoriteForUsers:  0,
					GoalDescription:   "bob brown goal 2",
					IsCompleted:       false,
					IsPrivate:         false,
					Mentors:           []openapiGeneral.SchemasUserPlainResponse{},
					MetricsDone:       0,
					MetricsTotal:      0,
					Name:              "bob brown way 2",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-07T00:40:00.000Z",
						Description: "A brief description about Bob.",
						Email:       "bob.brown@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
						IsMentor:    false,
						Name:        "Bob Brown",
						Uuid:        "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2",
					},
					UpdatedAt: "2024-07-09T00:00:00.000Z",
					Uuid:      "77a9e7c4-edb4-4b61-8065-cfd0c5c2506d",
					WayTags:   []openapiGeneral.SchemasWayTagResponse{},
				},
				{
					ChildrenUuids:     []string{},
					CopiedFromWayUuid: copiedFromWayUuid,
					CreatedAt:         "2024-07-09T00:00:00.000Z",
					DayReportsAmount:  0,
					EstimationTime:    40404040,
					FavoriteForUsers:  0,
					GoalDescription:   "bob brown goal 1",
					IsCompleted:       false,
					IsPrivate:         false,
					Mentors: []openapiGeneral.SchemasUserPlainResponse{
						{
							CreatedAt:   "2024-07-08T05:50:00.000Z",
							Description: "A brief description about Jane.",
							Email:       "jane.smith@example.com",
							ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
							IsMentor:    true,
							Name:        "Jane Smith",
							Uuid:        "8e77b89d-57c4-4b7f-8cd4-8dfc6bcb7d1b",
						},
						{
							CreatedAt:   "2024-07-06T05:00:00.000Z",
							Description: "A brief description about Dana.",
							Email:       "dana.evans@example.com",
							ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
							IsMentor:    true,
							Name:        "Dana Evans",
							Uuid:        "1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1",
						},
					},
					MetricsDone:  0,
					MetricsTotal: 0,
					Name:         "bob brown way 1",
					Owner: openapiGeneral.SchemasUserPlainResponse{
						CreatedAt:   "2024-07-07T00:40:00.000Z",
						Description: "A brief description about Bob.",
						Email:       "bob.brown@example.com",
						ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
						IsMentor:    false,
						Name:        "Bob Brown",
						Uuid:        "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2",
					},
					UpdatedAt: "2024-07-09T00:00:00.000Z",
					Uuid:      "77482c3f-cae6-494d-be1d-d06c1e84450b",
					WayTags:   []openapiGeneral.SchemasWayTagResponse{},
				},
			},
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.Size, ways.Size)
		assert.ElementsMatch(t, expectedData.Ways, ways.Ways)
	})
}

func TestGetWayById(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"
	wayID := "9972552a-c0b3-41f3-b464-284d36a36964"

	t.Run("should get all ways successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		way, response, err := generalApi.WayAPI.GetWayByUuid(ctx, wayID).Execute()
		if err != nil {
			t.Fatalf("Failed to get way: %v", err)
		}

		expectedData := &openapiGeneral.SchemasWayPlainResponse{
			ChildrenUuids:     []string{},
			CopiedFromWayUuid: openapiGeneral.NullableString{},
			CreatedAt:         "2024-07-09 00:00:00",
			DayReportsAmount:  0,
			EstimationTime:    40404040,
			FavoriteForUsers:  0,
			GoalDescription:   "ronnie stanton goal",
			IsCompleted:       false,
			IsPrivate:         false,
			Mentors: []openapiGeneral.SchemasUserPlainResponse{
				{
					CreatedAt:   "2024-07-09 00:00:00",
					Description: "A brief description about Ronnie.",
					Email:       "ronnie.stanton@example.com",
					ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F",
					IsMentor:    false,
					Name:        "Ronnie Stanton",
					Uuid:        "d63d2f89-6412-4324-8587-7061bf02dca4",
				},
			},
			MetricsDone:  0,
			MetricsTotal: 0,
			Name:         "ronnie stanton way",
			Owner: openapiGeneral.SchemasUserPlainResponse{
				CreatedAt:   "2024-07-09 00:00:00",
				Description: "A brief description about Ronnie.",
				Email:       "ronnie.stanton@example.com",
				ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F",
				IsMentor:    false,
				Name:        "Ronnie Stanton",
				Uuid:        "d63d2f89-6412-4324-8587-7061bf02dca4",
			},
			UpdatedAt: "2024-07-09 00:00:00",
			Uuid:      "9972552a-c0b3-41f3-b464-284d36a36964",
			WayTags:   []openapiGeneral.SchemasWayTagResponse{},
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.Uuid, way.Uuid)
		assert.Equal(t, expectedData.Name, way.Name)
		assert.Equal(t, expectedData.GoalDescription, way.GoalDescription)
	})
}

func TestGetWayStatisticsById(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"
	wayID := "550e8400-e29b-41d4-a716-446655440000"

	t.Run("should get all ways successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)
		wayStatistics, response, err := generalApi.WayAPI.GetWayStatisticsByUuid(ctx, wayID).Execute()
		if err != nil {
			t.Fatalf("Failed to get way statistics: %v", err)
		}

		expectedData := &openapiGeneral.SchemasWayStatisticsTriplePeriod{
			TotalTime: openapiGeneral.SchemasWayStatistics{
				TimeSpentByDayChart: []openapiGeneral.SchemasTimeSpentByDayPoint{
					{Date: "2024-07-09T00:00:00.000Z", Value: 0},
					{Date: "2024-07-10T00:00:00.000Z", Value: 0},
					{Date: "2024-07-11T00:00:00.000Z", Value: 0},
					{Date: "2024-07-12T00:00:00.000Z", Value: 0},
					{Date: "2024-07-13T00:00:00.000Z", Value: 0},
					{Date: "2024-07-14T00:00:00.000Z", Value: 0},
					{Date: "2024-07-15T00:00:00.000Z", Value: 0},
					{Date: "2024-07-16T00:00:00.000Z", Value: 0},
					{Date: "2024-07-17T00:00:00.000Z", Value: 0},
					{Date: "2024-07-18T00:00:00.000Z", Value: 0},
					{Date: "2024-07-19T00:00:00.000Z", Value: 0},
					{Date: "2024-07-20T00:00:00.000Z", Value: 0},
					{Date: "2024-07-21T00:00:00.000Z", Value: 0},
					{Date: "2024-07-22T00:00:00.000Z", Value: 0},
					{Date: "2024-07-23T00:00:00.000Z", Value: 0},
					{Date: "2024-07-24T00:00:00.000Z", Value: 0},
					{Date: "2024-07-25T00:00:00.000Z", Value: 0},
					{Date: "2024-07-26T00:00:00.000Z", Value: 0},
					{Date: "2024-07-27T00:00:00.000Z", Value: 0},
					{Date: "2024-07-28T00:00:00.000Z", Value: 0},
					{Date: "2024-07-29T00:00:00.000Z", Value: 0},
					{Date: "2024-07-30T00:00:00.000Z", Value: 0},
					{Date: "2024-07-31T00:00:00.000Z", Value: 0},
					{Date: "2024-08-01T00:00:00.000Z", Value: 10},
					{Date: "2024-08-02T00:00:00.000Z", Value: 0},
					{Date: "2024-08-03T00:00:00.000Z", Value: 120},
					{Date: "2024-08-04T00:00:00.000Z", Value: 0},
					{Date: "2024-08-05T00:00:00.000Z", Value: 60},
					{Date: "2024-08-06T00:00:00.000Z", Value: 0},
					{Date: "2024-08-07T00:00:00.000Z", Value: 60},
					{Date: "2024-08-08T00:00:00.000Z", Value: 0},
					{Date: "2024-08-09T00:00:00.000Z", Value: 60},
					{Date: "2024-08-10T00:00:00.000Z", Value: 0},
					{Date: "2024-08-11T00:00:00.000Z", Value: 60},
					{Date: "2024-08-12T00:00:00.000Z", Value: 0},
					{Date: "2024-08-13T00:00:00.000Z", Value: 60},
					{Date: "2024-08-14T00:00:00.000Z", Value: 0},
					{Date: "2024-08-15T00:00:00.000Z", Value: 0},
				},
				OverallInformation: openapiGeneral.SchemasOverallInformation{
					TotalTime:                 430,
					TotalReports:              8,
					FinishedJobs:              8,
					AverageTimePerCalendarDay: 11,
					AverageTimePerWorkingDay:  53,
					AverageJobTime:            54,
				},
				LabelStatistics: openapiGeneral.SchemasLabelStatistics{
					Labels: []openapiGeneral.SchemasLabelInfo{
						{
							JobsAmount:           1,
							JobsAmountPercentage: 12,
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "2461357d-f2f0-43a7-9f1d-79fd1eaa64f5",
								Name:        "general meeting",
								Color:       "yellow",
								Description: "this is not my tag",
							},
							Time:           60,
							TimePercentage: 13,
						},
						{
							JobsAmount:           3,
							JobsAmountPercentage: 37,
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "5ebb8d43-b685-4090-8453-ceaa7aad2095",
								Name:        "database",
								Color:       "green",
								Description: "this is not my tag",
							},
							Time:           130,
							TimePercentage: 30,
						},
						{
							JobsAmount:           2,
							JobsAmountPercentage: 25,
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "c73ff20b-e64e-4e5f-b270-1a40ba1bd81b",
								Name:        "coding",
								Color:       "blue",
								Description: "this is not my tag",
							},
							Time:           120,
							TimePercentage: 27,
						},
						{
							JobsAmount:           2,
							JobsAmountPercentage: 25,
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "60e7860d-58d9-4035-93d0-9bb825fe734c",
								Name:        "meeting 1:1",
								Color:       "red",
								Description: "this is not my tag",
							},
							Time:           120,
							TimePercentage: 27,
						},
					},
				},
			},
			LastMonth: openapiGeneral.SchemasWayStatistics{
				TimeSpentByDayChart: []openapiGeneral.SchemasTimeSpentByDayPoint{
					{Date: "2024-07-16T00:00:00.000Z", Value: 0},
					{Date: "2024-07-17T00:00:00.000Z", Value: 0},
					{Date: "2024-07-18T00:00:00.000Z", Value: 0},
					{Date: "2024-07-19T00:00:00.000Z", Value: 0},
					{Date: "2024-07-20T00:00:00.000Z", Value: 0},
					{Date: "2024-07-21T00:00:00.000Z", Value: 0},
					{Date: "2024-07-22T00:00:00.000Z", Value: 0},
					{Date: "2024-07-23T00:00:00.000Z", Value: 0},
					{Date: "2024-07-24T00:00:00.000Z", Value: 0},
					{Date: "2024-07-25T00:00:00.000Z", Value: 0},
					{Date: "2024-07-26T00:00:00.000Z", Value: 0},
					{Date: "2024-07-27T00:00:00.000Z", Value: 0},
					{Date: "2024-07-28T00:00:00.000Z", Value: 0},
					{Date: "2024-07-29T00:00:00.000Z", Value: 0},
					{Date: "2024-07-30T00:00:00.000Z", Value: 0},
					{Date: "2024-07-31T00:00:00.000Z", Value: 0},
					{Date: "2024-08-01T00:00:00.000Z", Value: 10},
					{Date: "2024-08-02T00:00:00.000Z", Value: 0},
					{Date: "2024-08-03T00:00:00.000Z", Value: 120},
					{Date: "2024-08-04T00:00:00.000Z", Value: 0},
					{Date: "2024-08-05T00:00:00.000Z", Value: 60},
					{Date: "2024-08-06T00:00:00.000Z", Value: 0},
					{Date: "2024-08-07T00:00:00.000Z", Value: 60},
					{Date: "2024-08-08T00:00:00.000Z", Value: 0},
					{Date: "2024-08-09T00:00:00.000Z", Value: 60},
					{Date: "2024-08-10T00:00:00.000Z", Value: 0},
					{Date: "2024-08-11T00:00:00.000Z", Value: 60},
					{Date: "2024-08-12T00:00:00.000Z", Value: 0},
					{Date: "2024-08-13T00:00:00.000Z", Value: 60},
					{Date: "2024-08-14T00:00:00.000Z", Value: 0},
					{Date: "2024-08-15T00:00:00.000Z", Value: 0},
				},
				OverallInformation: openapiGeneral.SchemasOverallInformation{
					AverageJobTime:            54,
					AverageTimePerCalendarDay: 14,
					AverageTimePerWorkingDay:  53,
					FinishedJobs:              8,
					TotalReports:              8,
					TotalTime:                 430,
				},
				LabelStatistics: openapiGeneral.SchemasLabelStatistics{
					Labels: []openapiGeneral.SchemasLabelInfo{
						{
							JobsAmount:           1,
							JobsAmountPercentage: 12,
							Label: openapiGeneral.SchemasLabel{
								Color:       "yellow",
								Description: "this is not my tag",
								Name:        "general meeting",
								Uuid:        "2461357d-f2f0-43a7-9f1d-79fd1eaa64f5",
							},
							Time:           60,
							TimePercentage: 13,
						},
						{
							JobsAmount:           3,
							JobsAmountPercentage: 37,
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "5ebb8d43-b685-4090-8453-ceaa7aad2095",
								Name:        "database",
								Color:       "green",
								Description: "this is not my tag",
							},
							Time:           130,
							TimePercentage: 30,
						},
						{
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "c73ff20b-e64e-4e5f-b270-1a40ba1bd81b",
								Name:        "coding",
								Color:       "blue",
								Description: "this is not my tag",
							},
							JobsAmount:           2,
							JobsAmountPercentage: 25,
							Time:                 120,
							TimePercentage:       27,
						},
						{
							JobsAmount:           2,
							JobsAmountPercentage: 25,
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "60e7860d-58d9-4035-93d0-9bb825fe734c",
								Name:        "meeting 1:1",
								Color:       "red",
								Description: "this is not my tag",
							},
							Time:           120,
							TimePercentage: 27,
						},
					},
				},
			},
			LastWeek: openapiGeneral.SchemasWayStatistics{
				TimeSpentByDayChart: []openapiGeneral.SchemasTimeSpentByDayPoint{
					{Date: "2024-08-09T00:00:00.000Z", Value: 60},
					{Date: "2024-08-10T00:00:00.000Z", Value: 0},
					{Date: "2024-08-11T00:00:00.000Z", Value: 60},
					{Date: "2024-08-12T00:00:00.000Z", Value: 0},
					{Date: "2024-08-13T00:00:00.000Z", Value: 60},
					{Date: "2024-08-14T00:00:00.000Z", Value: 0},
					{Date: "2024-08-15T00:00:00.000Z", Value: 0},
				},
				OverallInformation: openapiGeneral.SchemasOverallInformation{
					TotalTime:                 180,
					TotalReports:              4,
					FinishedJobs:              3,
					AverageTimePerCalendarDay: 26,
					AverageTimePerWorkingDay:  45,
					AverageJobTime:            60,
				},
				LabelStatistics: openapiGeneral.SchemasLabelStatistics{
					Labels: []openapiGeneral.SchemasLabelInfo{
						{
							JobsAmount:           1,
							JobsAmountPercentage: 33,
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "5ebb8d43-b685-4090-8453-ceaa7aad2095",
								Name:        "database",
								Color:       "green",
								Description: "this is not my tag",
							},
							Time:           60,
							TimePercentage: 33,
						},
						{
							JobsAmount:           1,
							JobsAmountPercentage: 33,
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "c73ff20b-e64e-4e5f-b270-1a40ba1bd81b",
								Name:        "coding",
								Color:       "blue",
								Description: "this is not my tag",
							},
							Time:           60,
							TimePercentage: 33,
						},
						{
							JobsAmount:           1,
							JobsAmountPercentage: 33,
							Label: openapiGeneral.SchemasLabel{
								Uuid:        "60e7860d-58d9-4035-93d0-9bb825fe734c",
								Name:        "meeting 1:1",
								Color:       "red",
								Description: "this is not my tag",
							},
							Time:           60,
							TimePercentage: 33,
						},
					},
				},
			},
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData, wayStatistics)
	})
}

func TestCreateWay(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"

	t.Run("should get all ways successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)

		request := openapiGeneral.SchemasCreateWayPayload{
			CopiedFromWayUuid: openapiGeneral.NullableString{},
			EstimationTime:    0,
			GoalDescription:   "Random Description",
			IsCompleted:       false,
			IsPrivate:         false,
			Name:              "Random name",
			OwnerUuid:         "1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1",
		}

		way, response, err := generalApi.WayAPI.CreateWay(ctx).Request(request).Execute()
		if err != nil {
			t.Fatalf("Failed to create way: %v", err)
		}

		expectedData := &openapiGeneral.SchemasWayPlainResponse{
			ChildrenUuids:    []string{},
			DayReportsAmount: 0,
			FavoriteForUsers: 0,
			GoalDescription:  "Random Description",
			IsCompleted:      false,
			IsPrivate:        false,
			Mentors:          []openapiGeneral.SchemasUserPlainResponse{},
			MetricsDone:      0,
			MetricsTotal:     0,
			Name:             "Random name",
			Owner: openapiGeneral.SchemasUserPlainResponse{
				CreatedAt:   "2024-07-06T05:00:00.000Z",
				Description: "A brief description about Dana.",
				Email:       "dana.evans@example.com",
				ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F&psig=AOvVaw2zWpFWOHXwuTI0x6EM4vXB&ust=1719409370844000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCID3x67x9oYDFQAAAAAdAAAAABAT",
				IsMentor:    true,
				Name:        "Dana Evans",
				Uuid:        "1b3d5e7f-5a1e-4d3a-b1a5-d1a1d5b7a7e1",
			},
			WayTags: []openapiGeneral.SchemasWayTagResponse{},
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.Name, way.Name)
		assert.Equal(t, expectedData.ChildrenUuids, way.ChildrenUuids)
		assert.Equal(t, expectedData.DayReportsAmount, way.DayReportsAmount)
		assert.Equal(t, expectedData.FavoriteForUsers, way.FavoriteForUsers)
		assert.Equal(t, expectedData.GoalDescription, way.GoalDescription)
		assert.Equal(t, expectedData.IsCompleted, way.IsCompleted)
		assert.Equal(t, expectedData.IsPrivate, way.IsPrivate)
		assert.ElementsMatch(t, expectedData.Mentors, way.Mentors)
		assert.Equal(t, expectedData.MetricsDone, way.MetricsDone)
		assert.Equal(t, expectedData.MetricsTotal, way.MetricsTotal)
		assert.Equal(t, expectedData.Owner, way.Owner)
		assert.ElementsMatch(t, expectedData.WayTags, way.WayTags)
	})
}

func TestUpdateWay(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"
	wayID := "9972552a-c0b3-41f3-b464-284d36a36964"

	t.Run("should get all ways successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)

		newName := "New newName"
		newStateIsCompleted := true

		request := openapiGeneral.SchemasUpdateWayPayload{
			IsCompleted: &newStateIsCompleted,
			Name:        &newName,
		}

		way, response, err := generalApi.WayAPI.UpdateWay(ctx, wayID).Request(request).Execute()

		expectedData := &openapiGeneral.SchemasWayPlainResponse{
			ChildrenUuids:     []string{},
			CopiedFromWayUuid: openapiGeneral.NullableString{},
			CreatedAt:         "2024-07-09 00:00:00",
			DayReportsAmount:  0,
			EstimationTime:    40404040,
			FavoriteForUsers:  0,
			GoalDescription:   "ronnie stanton goal",
			IsCompleted:       true,
			IsPrivate:         false,
			Mentors: []openapiGeneral.SchemasUserPlainResponse{
				{
					CreatedAt:   "2024-07-09 00:00:00",
					Description: "A brief description about Ronnie.",
					Email:       "ronnie.stanton@example.com",
					ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F",
					IsMentor:    false,
					Name:        "Ronnie Stanton",
					Uuid:        "d63d2f89-6412-4324-8587-7061bf02dca4",
				},
			},
			MetricsDone:  0,
			MetricsTotal: 0,
			Name:         "New newName",
			Owner: openapiGeneral.SchemasUserPlainResponse{
				CreatedAt:   "2024-07-09 00:00:00",
				Description: "A brief description about Ronnie.",
				Email:       "ronnie.stanton@example.com",
				ImageUrl:    "https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.gettyimages.com%2F",
				IsMentor:    false,
				Name:        "Ronnie Stanton",
				Uuid:        "d63d2f89-6412-4324-8587-7061bf02dca4",
			},
			UpdatedAt: "2024-07-09 00:00:00",
			Uuid:      "9972552a-c0b3-41f3-b464-284d36a36964",
			WayTags:   []openapiGeneral.SchemasWayTagResponse{},
		}

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, expectedData.Name, way.Name)
		assert.Equal(t, expectedData.IsCompleted, way.IsCompleted)
	})
}

func TestDeleteWay(t *testing.T) {
	newConfig, err := config.LoadConfig("../../")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	generalApi := openapi.MakeGeneralAPIClient(&newConfig)
	_, err = generalApi.DevAPI.DevResetDbGet(context.Background()).Execute()
	if err != nil {
		t.Fatalf("Failed to reset db: %v", err)
	}

	user := "d2cb5e1b-44df-48d3-b7a1-34f3d7a5b7e2"
	deletedWayID := "9972552a-c0b3-41f3-b464-284d36a36964"

	t.Run("should get all ways successfully", func(t *testing.T) {
		token, err := auth.GenerateJWT(user, newConfig.SecretSessionKey)
		if err != nil {
			t.Fatalf("Failed to generate JWT: %v", err)
		}

		ctx := context.WithValue(context.Background(), auth.ContextKeyAuthorization, "Bearer "+token)

		response, err := generalApi.WayAPI.DeleteWay(ctx, deletedWayID).Execute()
		if err != nil {
			t.Fatalf("Failed to delete Way: %v", err)
		}

		assert.Equal(t, http.StatusNoContent, response.StatusCode)

		ways, response, err := generalApi.WayAPI.GetAllWays(ctx).Status("all").Execute()
		if err != nil {
			t.Fatalf("Failed to get all Ways: %v", err)
		}

		isDeleted := true
		for _, way := range ways.GetWays() {
			if way.Uuid == deletedWayID {
				isDeleted = false
				break
			}
		}

		if !isDeleted {
			t.Fatalf("Failed to delete Way: %v", deletedWayID)
		}
	})
}
