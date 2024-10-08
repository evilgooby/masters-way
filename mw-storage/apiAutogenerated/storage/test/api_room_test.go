/*
Masters way storage API

Testing RoomAPIService

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

func Test_openapi_RoomAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RoomAPIService AddUserToRoom", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roomId string
		var userId string

		resp, httpRes, err := apiClient.RoomAPI.AddUserToRoom(context.Background(), roomId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService CreateRoom", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RoomAPI.CreateRoom(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService DeleteUserFromRoom", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roomId string
		var userId string

		httpRes, err := apiClient.RoomAPI.DeleteUserFromRoom(context.Background(), roomId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService GetRoomById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roomId string

		resp, httpRes, err := apiClient.RoomAPI.GetRoomById(context.Background(), roomId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RoomAPIService UpdateRoom", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var roomId string

		resp, httpRes, err := apiClient.RoomAPI.UpdateRoom(context.Background(), roomId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
