package integration

import (
	"Banner_Infrastructure/tests"
	"Banner_Infrastructure/tests/client"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	bannerDeleteID           = client.BaseURL + "/banner/"
	bannerDeleteFeatureOrTag = client.BaseURL + "/banner"
)

func TestDelete_Happy(t *testing.T) {
	_, c := client.New(t)
	// existing active banner
	lastTest := tests.GetRandomActive()
	testsTable := []struct {
		name string
		url  string
	}{
		{
			name: "by Id",
			url: bannerDeleteID + fmt.Sprintf("%d",
				tests.GetRandomActive().ID),
		},
		{
			name: "with feature",
			url: bannerDeleteFeatureOrTag + fmt.Sprintf("?feature_id=%d",
				tests.GetRandomActive().FeatureId),
		},
		{
			name: "by tag",
			url: bannerDeleteFeatureOrTag + fmt.Sprintf("?tag_id=%d",
				tests.GetRandomActive().TagIds[0]),
		},
		{
			name: "by tag and feature",
			url: bannerDeleteFeatureOrTag + fmt.Sprintf("?feature_id=%d&tag_id=%d",
				lastTest.FeatureId,
				lastTest.TagIds[0]),
		},
	}

	for _, tt := range testsTable {
		t.Run(tt.name, func(t *testing.T) {
			req := client.FormRequest(http.MethodDelete, tt.url, nil, client.AdminToken)
			resp, err := c.Client.Do(req)
			require.NoError(t, err)
			require.Equal(t, http.StatusNoContent, resp.StatusCode)
			resp.Body.Close()
		})
	}
}

func TestDelete_BadRequest(t *testing.T) {
	_, c := client.New(t)
	// existing active banner
	activeBanner := tests.GetRandomActive()

	teststable := []struct {
		name string
		url  string
	}{
		{
			name: "by Id and feature",
			url: bannerDeleteID + fmt.Sprintf("%d?feature_id=%d",
				activeBanner.ID,
				activeBanner.FeatureId),
		},
		{
			name: "by Id and tag",
			url: bannerDeleteID + fmt.Sprintf("%d?tag_id=%d",
				activeBanner.ID,
				activeBanner.TagIds[0]),
		},
		{
			name: "by Id, feature, tag",
			url: bannerDeleteID + fmt.Sprintf("%d?feature_id=%d&tag_id=%d",
				activeBanner.ID,
				activeBanner.FeatureId,
				activeBanner.TagIds[0]),
		},
		{
			name: "empty",
			url:  bannerDeleteFeatureOrTag,
		},
	}
	for _, tt := range teststable {
		t.Run(tt.name, func(t *testing.T) {
			req := client.FormRequest(http.MethodDelete, tt.url, nil, client.AdminToken)
			resp, err := c.Client.Do(req)
			require.NoError(t, err)
			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			resp.Body.Close()
		})
	}
}
