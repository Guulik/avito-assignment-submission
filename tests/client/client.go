package client

import (
	"Avito_trainee_assignment/internal/config"
	"context"
	"io"
	"net/http"
	"os"
	"testing"
)

type Suite struct {
	*testing.T
	Cfg    *config.Config
	Client *http.Client
}

const (
	BaseURL   = "http://localhost:4444"
	userToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJlbWFpbCI6ImR1bW15QHBpcC5jb20iLCJuYW1lIjoiVHVjayIsImFkbWluIjp0cnVlLCJleHAiOjUxMjUyMzM0MTIyMX0." +
		"53BBsUCuce2I5ZP98-dTKkk7iNhyWTj3j-vExwgZJQ4"
	adminToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJlbWFpbCI6ImR1bW15QHBpcC5jb20iLCJuYW1lIjoiVHVjayIsImFkbWluIjp0cnVlfQ." +
		"vT7s2Bu7Q1vf1FV86XNW26R-McbslMhnkQw7zvnltNE"
)

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadPath(configPath())

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	return ctx, &Suite{
		T:      t,
		Cfg:    cfg,
		Client: &http.Client{},
	}
}

func FormRequest(
	method string,
	url string,
	body io.Reader,
	admin bool,
) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	if admin {
		req.Header.Add("token", adminToken)
	} else {
		req.Header.Add("token", userToken)
	}
	return req
}

func configPath() string {
	const key = "CONFIG_PATH"

	if v := os.Getenv(key); v != "" {
		return v
	}

	return "../local.yaml"
}
