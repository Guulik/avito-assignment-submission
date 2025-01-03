package integration

import (
	"Banner_Infrastructure/tests"
	"Banner_Infrastructure/tests/client"
	"bytes"
	"encoding/json"
	"github.com/brianvoe/gofakeit"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

var CreateURL = client.BaseURL + "/banner"

func TestCreate_Happy(t *testing.T) {
	_, c := client.New(t)
	for i := 0; i < 200; i++ {
		body := *tests.RandomBody(gofakeit.Bool())
		bodyJSON, err := json.Marshal(body)
		require.NoError(t, err)

		bodyReq := bytes.NewReader(bodyJSON)
		req := client.FormRequest(http.MethodPost, CreateURL, bodyReq, client.AdminToken)
		resp, err := c.Client.Do(req)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode)
		resp.Body.Close()
	}
}

func TestCreate_Duplicate(t *testing.T) {
	_, c := client.New(t)
	body := *tests.RandomBody(true)
	bodyJSON, err := json.Marshal(body)
	require.NoError(t, err)

	bodyReq := bytes.NewReader(bodyJSON)
	req := client.FormRequest(http.MethodPost, CreateURL, bodyReq, client.AdminToken)
	resp, err := c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)
	defer resp.Body.Close()

	bodyJSON, err = json.Marshal(body)
	require.NoError(t, err)
	bodyReq = bytes.NewReader(bodyJSON)
	req = client.FormRequest(http.MethodPost, CreateURL, bodyReq, client.AdminToken)
	resp, err = c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	defer resp.Body.Close()
}
