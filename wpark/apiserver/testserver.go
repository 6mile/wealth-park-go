package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/yashmurty/wealth-park/wpark/config"
	"github.com/yashmurty/wealth-park/wpark/pkg/util"
	"go.uber.org/zap"
)

// APITestServer is used in controller tests.
type APITestServer struct {
	Server *Server
}

var (
	testServer      *APITestServer
	setupTestServer sync.Once
)

// GetTestServer returns an APITestServer singleton instance.
func GetTestServer(setupRoutes ...func(*Server)) *APITestServer {
	setupTestServer.Do(func() {
		c := config.GetInstance()

		// Create the test server.
		ts := &APITestServer{
			Server: NewServer(NewServerArgs{
				ServerID: c.ServerID,
				Addr:     c.GetAddr(),
			}),
		}

		// Setup additional routes.
		for _, f := range setupRoutes {
			f(ts.Server)
		}

		testServer = ts
	})
	return testServer
}

// CallAPI calls an API endpoint during testing.
func CallAPI(method, path string, in, out interface{}) (*httptest.ResponseRecorder, APIStatus) {
	// Marshal only non-nil payloads, this is useful when we don’t wanna pass
	// an inbound payload, e.g. GET /something/:id
	var j []byte
	var err error
	var status APIStatus

	if in != nil {
		j, err = json.Marshal(in)
		if err != nil {
			panic(err)
		}
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, path, bytes.NewReader(j))
	if err != nil {
		panic(err)
	}

	// Make the call.
	GetTestServer().Server.Engine.ServeHTTP(w, req)

	if w.Code < 300 && out != nil {
		// Unmarshal into non-nil struct.
		err = json.Unmarshal(w.Body.Bytes(), out)
		if err != nil {
			log.With(
				zap.String("from_json_string", w.Body.String()),
				zap.String("into_struct", util.GetJSON(out)),
			).Error("could not unmarshal server response")
			panic(err)
		}
	} else {
		err = json.Unmarshal(w.Body.Bytes(), &status)
		if err != nil {
			log.With(
				zap.String("from_json_string", w.Body.String()),
				zap.String("into_struct", util.GetJSON(out)),
			).Error("could not unmarshal server response")
			panic(err)
		}
	}

	return w, status
}
