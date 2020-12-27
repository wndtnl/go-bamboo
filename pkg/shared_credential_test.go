package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SharedCredential_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/shared_credential"
	useFileResponse(t, endpoint, "shared_credential/all.json")

	var sharedCredential = testClient.SharedCredential

	credentials, err := sharedCredential.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, credentials, "empty response")
	assert.Equal(t, 3, len(credentials))
}

func Test_SharedCredential_GetOne_Password(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/shared_credential/3473412"
	useFileResponse(t, endpoint, "shared_credential/one_password.json")

	var sharedCredential = testClient.SharedCredential

	credential, err := sharedCredential.GetOne(3473412)
	assert.Nil(t, err, err)
	assert.NotNil(t, credential, "empty response")
	assert.Equal(t, "yUlOmTpWohqvQ4WiPQ9h", credential.Username)
	assert.Equal(t, "0vfEHCoEdTyFg1Mqoxjd", credential.Password)
}

func Test_SharedCredential_Search_Password(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/shared_credential/search"
	useFileResponse(t, endpoint, "shared_credential/one_password.json")

	var sharedCredential = testClient.SharedCredential

	credential, err := sharedCredential.Search("LHnR8yFw33o4gfYgdSlZ")
	assert.Nil(t, err, err)
	assert.NotNil(t, credential, "empty response")
	assert.Equal(t, 3473412, credential.Id)
}

func Test_SharedCredential_Create_Password(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/shared_credential"
	useFileResponse(t, endpoint, "shared_credential/one_password.json")

	var sharedCredential = testClient.SharedCredential

	credential, err := sharedCredential.Create(&SharedCredential{
		Name:     "LHnR8yFw33o4gfYgdSlZ",
		Type:     "PASSWORD",
		Username: "yUlOmTpWohqvQ4WiPQ9h",
		Password: "0vfEHCoEdTyFg1Mqoxjd",
	})
	assert.Nil(t, err, err)
	assert.NotNil(t, credential, "empty response")
	assert.Equal(t, 3473412, credential.Id)
}

func Test_SharedCredential_Update_Password(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/shared_credential/3473412"
	useEmptyResponse(endpoint)

	var sharedCredential = testClient.SharedCredential

	err := sharedCredential.Update(3473412, &SharedCredential{
		Id:       3473412,
		Name:     "Agent2",
		Type:     "PASSWORD",
		Username: "username",
		Password: "password",
	})
	assert.Nil(t, err, err)
}

func Test_SharedCredential_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/shared_credential/3473412"
	useEmptyResponse(endpoint)

	var sharedCredential = testClient.SharedCredential

	err := sharedCredential.Delete(3473412)
	assert.Nil(t, err, err)
}
