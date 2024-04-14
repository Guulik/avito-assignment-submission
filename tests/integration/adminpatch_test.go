package integration

import (
	"Avito_trainee_assignment/tests"
	"Avito_trainee_assignment/tests/client"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var PatchURL = client.BaseURL + "/banner/"

func TestPatchContent(t *testing.T) {
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

	postedId := tests.GetPostedId(&body)
	// Change content
	newContentCreated := tests.RandomContent()
	bodyJSON, err = json.Marshal(newContentCreated)
	require.NoError(t, err)
	bodyReq = bytes.NewReader(bodyJSON)
	url := PatchURL + fmt.Sprintf("%d", postedId)
	req = client.FormRequest(http.MethodPatch, url, bodyReq, client.AdminToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	resp.Body.Close()
}
