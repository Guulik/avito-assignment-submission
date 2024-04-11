package tests

import (
	"Avito_trainee_assignment/tests/client"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

const (
	bannerGetUser = client.BaseURL + "/user_banner"
)

func TestUserGet_Happy(t *testing.T) {
	_, c := client.New(t)

	feature_id, tag_id := 10, 8
	url := bannerGetUser + fmt.Sprintf("?feature_id=%v&tag_id=%v", feature_id, tag_id)

	req := client.FormRequest(http.MethodGet, url, nil, false)
	resp, err := c.Client.Do(req)
	if err != nil {
		t.Error(err)
	}
	require.NoError(t, err)
	resultBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, string(resultBytes))

	var resultContentObj map[string]interface{}
	err = json.Unmarshal(resultBytes, &resultContentObj)

	require.NoError(t, err)
}
