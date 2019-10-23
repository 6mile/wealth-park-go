package e2e

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoot(t *testing.T) {
	// Should access root (/) successfully.
	{
		backend := getBackend()

		req, _ := http.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()
		backend.Server.Engine.ServeHTTP(resp, req)

		require.Contains(t, resp.Body.String(), "good")
	}
}
