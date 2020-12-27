package bamboo

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testMux    *http.ServeMux
	testServer *httptest.Server
	testClient *Client
)

func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	testClient, _ = NewClient(testServer.URL, DefaultUsername, DefaultPassword)
}

func teardown() {
	testServer.Close()
}

func useFileResponse(t *testing.T, endpoint, file string) {
	raw, err := ioutil.ReadFile("../testdata/" + file)
	assert.Nil(t, err, err)

	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	})
}

func useEmptyResponse(endpoint string) {
	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(""))
	})
}

func Test_NewDefaultClient(t *testing.T) {
	bamboo, err := NewDefaultClient()
	assert.Nil(t, err, err)
	assert.NotNil(t, bamboo, "Could not initialize client")
}
