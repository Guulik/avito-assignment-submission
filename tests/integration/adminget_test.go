package integration

import (
	tests "Banner_Infrastructure/tests"
	"Banner_Infrastructure/tests/client"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/require"
)

var GetAllURL = client.BaseURL + "/banner"

func TestGetBanners_Happy(t *testing.T) {
	_, c := client.New(t)
	features := tests.FeaturesSet().Slice()
	tags := tests.TagsSet().Slice()

	tests := []struct {
		name string
		url  string
	}{
		{
			name: "try to get all",
			url:  GetAllURL,
		},
		{
			name: "with feature",
			url: GetAllURL + fmt.Sprintf("?feature_id=%d",
				features[rand.Intn(len(features))]),
		},
		{
			name: "with tag",
			url: GetAllURL + fmt.Sprintf("?tag_id=%d",
				tags[rand.Intn(len(tags))]),
		},
		{
			name: "with tag and feature",
			url: GetAllURL + fmt.Sprintf("?feature_id=%d&tag_id=%d",
				features[rand.Intn(len(features))],
				tags[rand.Intn(len(tags))]),
		},
		{
			name: "with random tag and feature",
			url: GetAllURL + fmt.Sprintf("?feature_id=%d&tag_id=%d",
				gofakeit.Uint32(),
				gofakeit.Uint32()),
		},
		{
			name: "with limit",
			url: GetAllURL + fmt.Sprintf("?limit=%d",
				gofakeit.Uint16()),
		},
		{
			name: "with offset",
			url: GetAllURL + fmt.Sprintf("?offset=%d",
				gofakeit.Uint16()),
		},
		{
			name: "with limit and offset",
			url: GetAllURL + fmt.Sprintf("?offset=%d&limit%d",
				gofakeit.Uint16(),
				gofakeit.Uint16()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := client.FormRequest(http.MethodGet, tt.url, nil, client.AdminToken)
			resp, err := c.Client.Do(req)

			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode)
			resBytes, _ := io.ReadAll(resp.Body)
			require.NoError(t, err, err)

			// check if result deserializable
			var resultObj []map[string]interface{}
			err = json.Unmarshal(resBytes, &resultObj)
			require.NoError(t, err)
			resp.Body.Close()
		})
	}
}

func TestAdminGet_InvalidToken(t *testing.T) {
	_, c := client.New(t)
	url := GetAllURL

	// invalid token
	req := client.FormRequest(http.MethodGet, url, nil, "dummy invalid token")
	resp, err := c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	resp.Body.Close()

	req = client.FormRequest(http.MethodGet, url, nil, client.ExpiredToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	resp.Body.Close()

	req = client.FormRequest(http.MethodGet, url, nil, client.UserToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusForbidden, resp.StatusCode)
	resp.Body.Close()
}
