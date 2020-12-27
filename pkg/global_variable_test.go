package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GlobalVariable_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable"
	useFileResponse(t, endpoint, "global_variable/all.json")

	var globalVariable = testClient.GlobalVariable;

	variables, err := globalVariable.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, variables, "empty response")
	assert.Equal(t, 2, len(variables))
}

func Test_GlobalVariable_GetOne(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable/3211269"
	useFileResponse(t, endpoint, "global_variable/one.json")

	var globalVariable = testClient.GlobalVariable;

	variable, err := globalVariable.GetOne(3211269)
	assert.Nil(t, err, err)
	assert.NotNil(t, variable, "empty response")
	assert.Equal(t, "Database4", variable.Key)
}

func Test_GlobalVariable_Search(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable/search"
	useFileResponse(t, endpoint, "global_variable/one.json")

	var globalVariable = testClient.GlobalVariable;

	variable, err := globalVariable.Search("Database4")
	assert.Nil(t, err, err)
	assert.NotNil(t, variable, "empty response")
	assert.Equal(t, 3211269, variable.Id)
}

func Test_GlobalVariable_Create(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable"
	useFileResponse(t, endpoint, "global_variable/create.json")

	var globalVariable = testClient.GlobalVariable;

	variable, err := globalVariable.Create("MySecret", "TheSecretValue")
	assert.Nil(t, err, err)
	assert.NotNil(t, variable, "empty response")
	assert.Equal(t, "MySecret", variable.Key)
	assert.Equal(t, "TheSecretValue", variable.Value)
}

func Test_GlobalVariable_Update(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable/3211269"
	useEmptyResponse(endpoint)

	var globalVariable = testClient.GlobalVariable;

	err := globalVariable.Update(3211269, "MySecret", "NewSecretValue")
	assert.Nil(t, err, err)
}

func Test_GlobalVariable_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable/3211269"
	useEmptyResponse(endpoint)

	var globalVariable = testClient.GlobalVariable;

	err := globalVariable.Delete(3211269)
	assert.Nil(t, err, err)
}
