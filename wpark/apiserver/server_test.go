package apiserver

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicServer(t *testing.T) {
	var empty interface{}

	t.Run("should succeed and get root endpoint", func(t *testing.T) {
		rootResp := APIStatus{}
		w, _ := CallAPI("GET", "/", &empty, &rootResp)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, http.StatusOK, rootResp.Code)
	})

	t.Run("should succeed and get fail endpoint", func(t *testing.T) {
		w, status := CallAPI("GET", "/fail", &empty, &empty)
		require.Equal(t, http.StatusInternalServerError, w.Code)
		require.Contains(t, status.Message, "all hell")
	})

}
