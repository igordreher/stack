/*
Auth API

Testing ClientsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package authclient

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "github.com/formancehq/auth/authclient"
)

func Test_authclient_ClientsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test ClientsApiService AddScopeToClient", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clientId interface{}
        var scopeId interface{}

        resp, httpRes, err := apiClient.ClientsApi.AddScopeToClient(context.Background(), clientId, scopeId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClientsApiService CreateClient", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ClientsApi.CreateClient(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClientsApiService CreateSecret", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clientId interface{}

        resp, httpRes, err := apiClient.ClientsApi.CreateSecret(context.Background(), clientId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClientsApiService DeleteClient", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clientId interface{}

        resp, httpRes, err := apiClient.ClientsApi.DeleteClient(context.Background(), clientId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClientsApiService DeleteScopeFromClient", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clientId interface{}
        var scopeId interface{}

        resp, httpRes, err := apiClient.ClientsApi.DeleteScopeFromClient(context.Background(), clientId, scopeId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClientsApiService DeleteSecret", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clientId interface{}
        var secretId interface{}

        resp, httpRes, err := apiClient.ClientsApi.DeleteSecret(context.Background(), clientId, secretId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClientsApiService ListClients", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.ClientsApi.ListClients(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClientsApiService ReadClient", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clientId interface{}

        resp, httpRes, err := apiClient.ClientsApi.ReadClient(context.Background(), clientId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test ClientsApiService UpdateClient", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var clientId interface{}

        resp, httpRes, err := apiClient.ClientsApi.UpdateClient(context.Background(), clientId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
