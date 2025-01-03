package client

import (
	"Banner_Infrastructure/internal/configure"
	"context"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"testing"
)

type Suite struct {
	*testing.T
	Cfg    *configure.Config
	Client *http.Client
}

var (
	BaseURL   = "http://localhost:4444"
	UserToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJlbWFpbCI6ImR1bW15QHBvcC5jb3JuIiwibmFtZSI6InNpbGx5IiwiYWRtaW4iOmZhbHNlLCJleHAiOjUxMjUyMzM0MTIyMX0." +
		"NJPL563Qey8-WqVvZ_WO-IHCxUUCDicJpmfG-CTCGAM"
	AdminToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJlbWFpbCI6ImR1bW15QHBpcC5jb20iLCJuYW1lIjoiVHVjayIsImFkbWluIjp0cnVlfQ." +
		"vT7s2Bu7Q1vf1FV86XNW26R-McbslMhnkQw7zvnltNE"
	ExpiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
		"eyJlbWFpbCI6ImR1bW15QHBvcC5jb3JuIiwibmFtZSI6InNpbGx5IiwiYWRtaW4iOmZhbHNlLCJleHAiOjE3MTI3NjQ2ODF9." +
		"N9napZzgIolAdZu7Hee9oAjRRuSR6VqcqSilRfoidnk"
)

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := configure.MustLoadPath(configPath())

	BaseURL = address(cfg)

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
	token string,
) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", token)

	return req
}

func configPath() string {
	const key = "CONFIG_PATH"

	if v := os.Getenv(key); v != "" {
		return v
	}

	return "../../config/stage.yaml"
}

func address(cfg *configure.Config) string {
	return net.JoinHostPort(`http://localhost`, strconv.Itoa(cfg.Port))
}
