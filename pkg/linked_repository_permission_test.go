package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LinkedRepositoryPermission_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository_permission/4980737"
	useFileResponse(t, endpoint, "linked_repository_permission/all.json")

	var linkedRepositoryPermissionService = testClient.LinkedRepositoryPermission

	linkedRepositoryPermissions, err := linkedRepositoryPermissionService.GetAll(4980737)
	assert.Nil(t, err, err)
	assert.NotNil(t, linkedRepositoryPermissions, "empty response")
	assert.Equal(t, 3, len(linkedRepositoryPermissions))
}

func Test_LinkedRepositoryPermission_GetOne_Group(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository_permission/4980737/group/bamboo-admin"
	useFileResponse(t, endpoint, "linked_repository_permission/one_group.json")

	var linkedRepositoryPermissionService = testClient.LinkedRepositoryPermission

	linkedRepositoryPermission, err :=
		linkedRepositoryPermissionService.GetOne(4980737, "group", "bamboo-admin")
	assert.Nil(t, err, err)
	assert.NotNil(t, linkedRepositoryPermission, "empty response")
	assert.Equal(t, "bamboo-admin", linkedRepositoryPermission.Name)
	assert.Equal(t, 2, len(linkedRepositoryPermission.Permissions))
}

func Test_LinkedRepositoryPermission_Upsert(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository_permission"
	useFileResponse(t, endpoint, "linked_repository_permission/one_group.json")

	var linkedRepositoryPermissionService = testClient.LinkedRepositoryPermission

	linkedRepositoryPermission, err := linkedRepositoryPermissionService.Upsert(&LinkedRepositoryPermission{
		RepositoryId: 4980737,
		Name: "bamboo-admin",
		Type: "GROUP",
		Permissions: []string{
			"READ",
			"ADMINISTRATION",
		},
	})
	assert.Nil(t, err, err)
	assert.NotNil(t, linkedRepositoryPermission, "empty response")
	assert.Equal(t, "bamboo-admin", linkedRepositoryPermission.Name)
	assert.Equal(t, 2, len(linkedRepositoryPermission.Permissions))
}

func Test_LinkedRepositoryPermission_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository_permission/4980737/group/bamboo-admin"
	useEmptyResponse(endpoint)

	var linkedRepositoryPermissionService = testClient.LinkedRepositoryPermission

	err := linkedRepositoryPermissionService.Delete(4980737, "group", "bamboo-admin")
	assert.Nil(t, err, err)
}
