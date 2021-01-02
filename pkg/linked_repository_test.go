package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LinkedRepository_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository"
	useFileResponse(t, endpoint, "linked_repository/all.json")

	var linkedRepositoryService = testClient.LinkedRepository

	linkedRepositories, err := linkedRepositoryService.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, linkedRepositories, "empty response")
	assert.Equal(t, 5, len(linkedRepositories))
}

func Test_LinkedRepository_GetOne(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository/3309579"
	useFileResponse(t, endpoint, "linked_repository/one.json")

	var linkedRepositoryService = testClient.LinkedRepository

	linkedRepository, err := linkedRepositoryService.GetOne("3309579")
	assert.Nil(t, err, err)
	assert.NotNil(t, linkedRepository, "empty response")
	assert.Equal(t, "0j7hmptQcLYSgV1LdLD1", linkedRepository.Name)
}

func Test_LinkedRepository_Search(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository/search"
	useFileResponse(t, endpoint, "linked_repository/one.json")

	var linkedRepositoryService = testClient.LinkedRepository

	linkedRepository, err := linkedRepositoryService.Search("0j7hmptQcLYSgV1LdLD1")
	assert.Nil(t, err, err)
	assert.NotNil(t, linkedRepository, "empty response")
	assert.Equal(t, "3309579", linkedRepository.Id)
}

func Test_LinkedRepository_Create_Git_None(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository"
	useFileResponse(t, endpoint, "linked_repository/one.json")

	var linkedRepositoryService = testClient.LinkedRepository

	linkedRepository, err := linkedRepositoryService.Create(&LinkedRepository{
		Name:     "0j7hmptQcLYSgV1LdLD1",
		Type:     "GIT",
		AuthType: "NONE",
	})
	assert.Nil(t, err, err)
	assert.NotNil(t, linkedRepository, "empty response")
	assert.Equal(t, "0j7hmptQcLYSgV1LdLD1", linkedRepository.Name)
}

func Test_LinkedRepository_Update_Git_None(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository/3309579"
	useEmptyResponse(endpoint)

	var linkedRepositoryService = testClient.LinkedRepository

	err := linkedRepositoryService.Update("3309579", &LinkedRepository{
		Id:       "3309579",
		Name:     "updated-name",
		Type:     "GIT",
		AuthType: "NONE",
	})
	assert.Nil(t, err, err)
}

func Test_LinkedRepository_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/linked_repository/3309579"
	useEmptyResponse(endpoint)

	var linkedRepositoryService = testClient.LinkedRepository

	err := linkedRepositoryService.Delete("3309579")
	assert.Nil(t, err, err)
}
