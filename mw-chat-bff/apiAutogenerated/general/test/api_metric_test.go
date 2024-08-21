/*
Masters way general API

Testing MetricAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_MetricAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MetricAPIService CreateMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MetricAPI.CreateMetric(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricAPIService DeleteMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricId string

		httpRes, err := apiClient.MetricAPI.DeleteMetric(context.Background(), metricId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricAPIService UpdateMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricId string

		resp, httpRes, err := apiClient.MetricAPI.UpdateMetric(context.Background(), metricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
