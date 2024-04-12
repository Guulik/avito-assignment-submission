package integration

import (
	"Avito_trainee_assignment/tests"
	"Avito_trainee_assignment/tests/client"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

var (
	bannerGetUserURL = client.BaseURL + "/user_banner"
)

func TestUserGet_Happy(t *testing.T) {
	_, c := client.New(t)
	//existing active banner
	activeBanner := tests.GetRandomActive()
	url := bannerGetUserURL + fmt.Sprintf("?feature_id=%d&tag_id=%d",
		activeBanner.FeatureId,
		activeBanner.TagIds[0])
	req := client.FormRequest(http.MethodGet, url, nil, client.UserToken)
	resp, err := c.Client.Do(req)

	require.NoError(t, err)
	resultBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, string(resultBytes))

	//check if result deserializable
	var resultContentObj map[string]interface{}
	err = json.Unmarshal(resultBytes, &resultContentObj)

	require.NoError(t, err)
}

func TestUserGet_BadRequest(t *testing.T) {
	_, c := client.New(t)

	tests := []struct {
		name        string
		url         string
		expectedErr string
	}{
		{
			name: "with negative feature",
			url: bannerGetUserURL + fmt.Sprintf("?feature_id=%v&tag_id=%v",
				-1*int(gofakeit.Uint32()),
				gofakeit.Uint32()),
			expectedErr: "incorrect featureId",
		},
		{
			name: "with negative tag",
			url: bannerGetUserURL + fmt.Sprintf("?feature_id=%v&tag_id=%v",
				1*int(gofakeit.Uint32()),
				-1*int(gofakeit.Uint32())),
			expectedErr: "incorrect tagId",
		},
		{
			name: "with negative both feature and tag",
			url: bannerGetUserURL + fmt.Sprintf("?feature_id=%v&tag_id=%v",
				-1*int(gofakeit.Uint32()),
				-1*int(gofakeit.Uint32())),
			expectedErr: "incorrect featureId",
		},
		{
			name:        "without tagId",
			url:         bannerGetUserURL + fmt.Sprintf("?feature_id=%v", int(gofakeit.Uint32())),
			expectedErr: "incorrect tagId",
		},
		{
			name:        "without featureId",
			url:         bannerGetUserURL + fmt.Sprintf("?tag_id=%v", int(gofakeit.Uint32())),
			expectedErr: "incorrect featureId",
		},
		{
			name:        "without both featureId and tagId",
			url:         bannerGetUserURL,
			expectedErr: "incorrect featureId",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := client.FormRequest(http.MethodGet, tt.url, nil, client.UserToken)
			resp, err := c.Client.Do(req)

			require.NoError(t, err)
			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			resBytes, _ := io.ReadAll(resp.Body)
			require.Contains(t, string(resBytes), tt.expectedErr)
		})
	}
}

func TestUserGet_NotFound(t *testing.T) {
	_, c := client.New(t)
	inactiveBanner := tests.GetFirstInactive()
	tests := []struct {
		name        string
		url         string
		expectedErr string
	}{
		{
			//most likely the banner will not exist by random combination.
			//though the opposite is possible
			name: "random non-existent banner",
			url: bannerGetUserURL + fmt.Sprintf("?feature_id=%v&tag_id=%v",
				gofakeit.Uint32(),
				gofakeit.Uint32()),
			expectedErr: "Баннер для пользователя не найден",
		},
		{
			name: "inactive banner",
			url: bannerGetUserURL + fmt.Sprintf("?feature_id=%v&tag_id=%v",
				inactiveBanner.FeatureId,
				inactiveBanner.TagIds[0]),
			expectedErr: "Баннер для пользователя не найден",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := client.FormRequest(http.MethodGet, tt.url, nil, client.UserToken)
			resp, err := c.Client.Do(req)

			require.NoError(t, err)
			require.Equal(t, http.StatusNotFound, resp.StatusCode)
		})
	}
}

func TestUserGet_InvalidToken(t *testing.T) {
	_, c := client.New(t)
	activeBanner := tests.GetRandomActive()
	url := bannerGetUserURL + fmt.Sprintf("?feature_id=%d&tag_id=%d",
		activeBanner.FeatureId,
		activeBanner.TagIds[0])

	req := client.FormRequest(http.MethodGet, url, nil, "dummy invalid token")
	resp, err := c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	req = client.FormRequest(http.MethodGet, url, nil, client.ExpiredToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
