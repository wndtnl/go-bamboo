package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_User_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/user"
	useFileResponse(t, endpoint, "user/all.json")

	var userService = testClient.User

	users, err := userService.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, users, "empty response")
	assert.Equal(t, 3, len(users))
}

func Test_User_GetOne(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/user/admin"
	useFileResponse(t, endpoint, "user/one.json")

	var userService = testClient.User

	user, err := userService.GetOne("admin")
	assert.Nil(t, err, err)
	assert.NotNil(t, user, "empty response")
	assert.Equal(t, "admin", user.Username)
	assert.Equal(t, "admin@example.com", user.Email)
}

func Test_User_Create(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/user"
	useFileResponse(t, endpoint, "user/one.json")

	var userService = testClient.User

	user, err := userService.Create(&User{
		Username:    "admin",
		FullName: 	"Admin",
		Email: "admin@example.com",
		Password: "sUp3rS3crEt!",
		JabberAddress: "",
		Enabled: true,
	})
	assert.Nil(t, err, err)
	assert.NotNil(t, user, "empty response")
	assert.Equal(t, "admin", user.Username)
	assert.Equal(t, "admin@example.com", user.Email)
}

func Test_User_Update(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/user/admin"
	useEmptyResponse(endpoint)

	var userService = testClient.User

	err := userService.Update("admin", &User{
		Username:    "admin",
		FullName: 	"The Admin",
		Email: "admin@example.com",
		Password: "",
		JabberAddress: "",
		Enabled: false,
	})
	assert.Nil(t, err, err)
}

func Test_User_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/user/admin"
	useEmptyResponse(endpoint)

	var userService = testClient.User

	err := userService.Delete("admin")
	assert.Nil(t, err, err)
}
