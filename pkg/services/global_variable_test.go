package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/wndtnl/go-bamboo/pkg/client"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testMux    = http.NewServeMux()
	testServer *httptest.Server

	rest                  *client.Rest
	globalVariableService *GlobalVariableService
)

func setup() {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	rest, _ = client.NewBasicAuthClient(testServer.URL, "", "")

	globalVariableService = NewGlobalVariableService(rest)
}

func teardown() {
	testServer.Close()
}

func useFileResponse(t *testing.T, endpoint, file string) {
	raw, err := ioutil.ReadFile("../../testdata/" + file)
	assert.Nil(t, err, err)

	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(raw)
	})
}

func useEmptyResponse(t *testing.T, endpoint string) {
	testMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(""))
	})
}

func Test_GlobalVariable_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable"
	useFileResponse(t, endpoint, "global_variable/all.json")

	variables, err := globalVariableService.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, variables, "empty response")
	assert.Equal(t, 2, len(variables))
}

func Test_GlobalVariable_GetOne(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable/3211269"
	useFileResponse(t, endpoint, "global_variable/one.json")

	variable, err := globalVariableService.GetOne(3211269)
	assert.Nil(t, err, err)
	assert.NotNil(t, variable, "empty response")
	assert.Equal(t, "Database4", variable.Key)
}

func Test_GlobalVariable_Search(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable/search"
	useFileResponse(t, endpoint, "global_variable/one.json")

	variable, err := globalVariableService.Search("Database4")
	assert.Nil(t, err, err)
	assert.NotNil(t, variable, "empty response")
	assert.Equal(t, 3211269, variable.Id)
}

func Test_GlobalVariable_Create(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable"
	useFileResponse(t, endpoint, "global_variable/create.json")

	variable, err := globalVariableService.Create("MySecret", "TheSecretValue")
	assert.Nil(t, err, err)
	assert.NotNil(t, variable, "empty response")
	assert.Equal(t, "MySecret", variable.Key)
	assert.Equal(t, "TheSecretValue", variable.Value)
}

func Test_GlobalVariable_Update(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable/3211269"
	useEmptyResponse(t, endpoint)

	err := globalVariableService.Update(3211269, "MySecret", "NewSecretValue")
	assert.Nil(t, err, err)
}

func Test_GlobalVariable_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable/3211269"
	useEmptyResponse(t, endpoint)

	err := globalVariableService.Delete(3211269)
	assert.Nil(t, err, err)
}
