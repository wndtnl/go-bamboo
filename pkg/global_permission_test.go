package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GlobalPermission_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_permission"
	useFileResponse(t, endpoint, "global_permission/all.json")

	var globalPermissionService = testClient.GlobalPermission

	globalPermissions, err := globalPermissionService.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, globalPermissions, "empty response")
	assert.Equal(t, 2, len(globalPermissions))
}

func Test_GlobalPermission_GetOne_User(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_permission/user/bob"
	useFileResponse(t, endpoint, "global_permission/one_user.json")

	var globalPermissionService = testClient.GlobalPermission

	globalPermission, err := globalPermissionService.GetOne("user", "bob")
	assert.Nil(t, err, err)
	assert.NotNil(t, globalPermission, "empty response")
	assert.Equal(t, "bob", globalPermission.Name)
	assert.Equal(t, 2, len(globalPermission.Permissions))
}

func Test_GlobalPermission_Upsert(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_permission"
	useFileResponse(t, endpoint, "global_permission/one_group.json")

	var globalPermissionService = testClient.GlobalPermission

	globalPermission, err := globalPermissionService.Upsert(&GlobalPermission{
		Name:    "bamboo-admin",
		Type:	 "GROUP",
		Permissions: []string{
			"CREATE",
			"READ",
			"ADMINISTRATION",
			"CREATEREPOSITORY"},
	})
	assert.Nil(t, err, err)
	assert.NotNil(t, globalPermission, "empty response")
	assert.Equal(t, "bamboo-admin", globalPermission.Name)
	assert.Equal(t, 4, len(globalPermission.Permissions))
}

func Test_GlobalPermission_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/global_permission/user/bob"
	useEmptyResponse(endpoint)

	var globalPermissionService = testClient.GlobalPermission

	err := globalPermissionService.Delete("user", "bob")
	assert.Nil(t, err, err)
}