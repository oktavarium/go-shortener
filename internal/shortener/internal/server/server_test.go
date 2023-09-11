package server

import (
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const addr string = "http://localhost:8080/"

func TestServer(t *testing.T) {
	srv := NewServer(addr)
	testSrv := httptest.NewServer(srv.router)
	defer testSrv.Close()

	tests := []struct {
		name       string
		method     string
		url        string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Post request with empty body but with url",
			method:     "POST",
			url:        "wrongurl",
			body:       "",
			wantStatus: 400,
			wantBody:   "",
		},
		{
			name:       "Post request with body",
			method:     "POST",
			url:        "",
			body:       "ya.ru",
			wantStatus: 201,
			wantBody:   "http://localhost:8080/eWEucnU=",
		},
		{
			name:       "Post request with all empty",
			method:     "POST",
			url:        "",
			body:       "",
			wantStatus: 400,
			wantBody:   "",
		},
		{
			name:       "Get request without url",
			method:     "GET",
			url:        "",
			body:       "",
			wantStatus: 400,
			wantBody:   "",
		},
		{
			name:       "Get request with bad url",
			method:     "GET",
			url:        "wrongurl",
			body:       "",
			wantStatus: 400,
			wantBody:   "",
		},
		{
			name:       "Get request with good url",
			method:     "GET",
			url:        "eWEucnU=",
			body:       "",
			wantStatus: 307,
			wantBody:   "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := resty.New()
			r, err := client.R().
				SetBody(test.body).
				Execute(test.method, addr+test.url)
			require.NoError(t, err)
			assert.Equal(t, test.wantStatus, r.StatusCode())
			assert.Equal(t, test.wantBody, string(r.Body()))

		})
	}
}

func newServer(t *testing.T) Server {
	t.Helper()
	return NewServer(":8080")
}
