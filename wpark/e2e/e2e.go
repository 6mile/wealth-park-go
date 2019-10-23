package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"

	"github.com/yashmurty/wealth-park/wpark/apiserver"
	"github.com/yashmurty/wealth-park/wpark/backend"
	"github.com/yashmurty/wealth-park/wpark/controller"
)

var (
	initTestData sync.Once
)

func getBackend() *backend.Backend {
	// Create an MySQL based backend.
	be := backend.NewBackendWithMYSQLModels()
	// Create an API server instance.
	be.Server = apiserver.NewServer(apiserver.NewServerArgs{})
	controller.SetupHTTPHandlers(be.Server)
	return be
}

// setupE2ETests *MUST* be run before doing any e2e testings.
func setupE2ETests() {
	initTestData.Do(func() {
		// Recreate all tables.
		getBackend().CreateTables()
	})
}

// CallAPI ...
func CallAPI(method, path, token string, in, out interface{}) (*httptest.ResponseRecorder, apiserver.APIStatus) {
	var j []byte
	var err error
	var status apiserver.APIStatus

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

	// Add token as request header.
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	// Make the call.
	getBackend().Server.Engine.ServeHTTP(w, req)

	// println("Call API Result:")
	// println(core.GetPrettyJSONString(w.Body.Bytes()))

	if w.Code < 300 && out != nil {
		// Unmarshal into non-nil struct.
		err = json.Unmarshal(w.Body.Bytes(), out)
		if err != nil {
			panic(err)
		}
	} else {
		err = json.Unmarshal(w.Body.Bytes(), &status)
		if err != nil {
			panic("Got response: '" + w.Body.String() + "' Code:" + strconv.Itoa(w.Code))
		}
	}

	return w, status
}
