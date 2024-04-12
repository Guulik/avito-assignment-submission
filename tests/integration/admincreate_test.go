package integration

import (
	"Avito_trainee_assignment/tests"
	"Avito_trainee_assignment/tests/client"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

var (
	CreateURL = client.BaseURL + "/banner"
)

func TestCreate_Happy(t *testing.T) {
	_, c := client.New(t)
	body := *tests.RandomBody()
	bodyJSON, err := json.Marshal(body)
	require.NoError(t, err)

	bodyReq := bytes.NewReader(bodyJSON)
	req := client.FormRequest(http.MethodPost, CreateURL, bodyReq, client.AdminToken)
	resp, err := c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

}

func TestCreate_Duplicate(t *testing.T) {
	_, c := client.New(t)
	body := *tests.RandomBody()
	bodyJSON, err := json.Marshal(body)
	require.NoError(t, err)

	bodyReq := bytes.NewReader(bodyJSON)
	req := client.FormRequest(http.MethodPost, CreateURL, bodyReq, client.AdminToken)
	resp, err := c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	bodyJSON, err = json.Marshal(body)
	require.NoError(t, err)
	bodyReq = bytes.NewReader(bodyJSON)
	req = client.FormRequest(http.MethodPost, CreateURL, bodyReq, client.AdminToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
