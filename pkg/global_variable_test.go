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

	var globalVariableService = testClient.GlobalVariable

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

	var globalVariableService = testClient.GlobalVariable

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

	var globalVariableService = testClient.GlobalVariable

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

	var globalVariableService = testClient.GlobalVariable

	variable, err := globalVariableService.Create(&GlobalVariable{
		Key:   "MySecret",
		Value: "TheSecretValue",
	})
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

	var globalVariableService = testClient.GlobalVariable

	err := globalVariableService.Update(3211269, &GlobalVariable{
		Id:    3211269,
		Key:   "MySecret",
		Value: "TheSecretValue",
	})
	assert.Nil(t, err, err)
}

func Test_GlobalVariable_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_variable/3211269"
	useEmptyResponse(endpoint)

	var globalVariableService = testClient.GlobalVariable

	err := globalVariableService.Delete(3211269)
	assert.Nil(t, err, err)
}
