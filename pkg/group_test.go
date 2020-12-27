package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Group_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/group"
	useFileResponse(t, endpoint, "group/all.json")

	var groupService = testClient.Group

	groups, err := groupService.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, groups, "empty response")
	assert.Equal(t, 3, len(groups))
}

func Test_Group_GetOne(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/group/bamboo-admin"
	useFileResponse(t, endpoint, "group/one.json")

	var groupService = testClient.Group

	group, err := groupService.GetOne("bamboo-admin")
	assert.Nil(t, err, err)
	assert.NotNil(t, group, "empty response")
	assert.Equal(t, "bamboo-admin", group.Name)
	assert.Equal(t, 1, len(group.Members))
}

func Test_Group_Create(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/group"
	useFileResponse(t, endpoint, "group/one.json")

	var groupService = testClient.Group

	group, err := groupService.Create(&Group{
		Name:    "bamboo-admin",
		Members: []string{"admin"},
	})
	assert.Nil(t, err, err)
	assert.NotNil(t, group, "empty response")
	assert.Equal(t, "bamboo-admin", group.Name)
	assert.Equal(t, 1, len(group.Members))
}

func Test_Group_Update(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/group/bamboo-admin"
	useEmptyResponse(endpoint)

	var groupService = testClient.Group

	err := groupService.Update("bamboo-admin", &Group{
		Name:    "bamboo-admin",
		Members: []string{"admin", "jane-doe"},
	})
	assert.Nil(t, err, err)
}

func Test_Group_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/group/bamboo-admin"
	useEmptyResponse(endpoint)

	var groupService = testClient.Group

	err := groupService.Delete("bamboo-admin")
	assert.Nil(t, err, err)
}
