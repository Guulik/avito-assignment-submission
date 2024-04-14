package end2end

import (
	"Avito_trainee_assignment/tests"
	"Avito_trainee_assignment/tests/client"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	CreateURL                = client.BaseURL + "/banner"
	PatchURL                 = client.BaseURL + "/banner/"
	bannerGetUserURL         = client.BaseURL + "/user_banner"
	bannerDeleteFeatureOrTag = client.BaseURL + "/banner"
)

func TestChangeVisibility(t *testing.T) {
	_, c := client.New(t)

	// Create banner
	body := *tests.RandomBody(true)
	bodyJSON, err := json.Marshal(body)
	require.NoError(t, err)
	bodyReq := bytes.NewReader(bodyJSON)
	req := client.FormRequest(http.MethodPost, CreateURL, bodyReq, client.AdminToken)
	resp, err := c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	resp.Body.Close()

	// UserGet banner
	url := bannerGetUserURL + fmt.Sprintf("?feature_id=%d&tag_id=%d",
		body.Feature,
		body.Tags[0])
	req = client.FormRequest(http.MethodGet, url, nil, client.UserToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	resultBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, string(resultBytes))
	var resultContentObj map[string]interface{}
	err = json.Unmarshal(resultBytes, &resultContentObj)
	require.NoError(t, err)
	// check if result deserializable
	resp.Body.Close()

	postedId := tests.GetPostedId(&body)

	// Change visibility
	body.IsActive = false
	bodyJSON, err = json.Marshal(body)
	require.NoError(t, err)
	bodyReq = bytes.NewReader(bodyJSON)
	url = PatchURL + fmt.Sprintf("%d", postedId)
	req = client.FormRequest(http.MethodPatch, url, bodyReq, client.AdminToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()

	// Attempt to get inactive banner
	url = bannerGetUserURL + fmt.Sprintf("?feature_id=%d&tag_id=%d",
		body.Feature,
		body.Tags[0])
	req = client.FormRequest(http.MethodGet, url, nil, client.UserToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}

func TestDelete(t *testing.T) {
	_, c := client.New(t)

	// Create banner
	body := *tests.RandomBody(true)
	bodyJSON, err := json.Marshal(body)
	require.NoError(t, err)
	bodyReq := bytes.NewReader(bodyJSON)
	req := client.FormRequest(http.MethodPost, CreateURL, bodyReq, client.AdminToken)
	resp, err := c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	resp.Body.Close()

	// UserGet banner
	url := bannerGetUserURL + fmt.Sprintf("?feature_id=%d&tag_id=%d",
		body.Feature,
		body.Tags[0])
	req = client.FormRequest(http.MethodGet, url, nil, client.UserToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	resultBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, string(resultBytes))
	var resultContentObj map[string]interface{}
	err = json.Unmarshal(resultBytes, &resultContentObj)
	require.NoError(t, err)
	// check if result deserializable
	resp.Body.Close()

	// Delete banner
	url = bannerDeleteFeatureOrTag + fmt.Sprintf("?feature_id=%d&tag_id=%d",
		body.Feature,
		body.Tags[0])
	req = client.FormRequest(http.MethodDelete, url, nil, client.AdminToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
	resp.Body.Close()

	// Attempt to get deleted banner
	url = bannerGetUserURL + fmt.Sprintf("?feature_id=%d&tag_id=%d",
		body.Feature,
		body.Tags[0])
	req = client.FormRequest(http.MethodGet, url, nil, client.UserToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	resp.Body.Close()
}
