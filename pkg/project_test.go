package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Project_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/project"
	useFileResponse(t, endpoint, "project/all.json")

	var projectService = testClient.Project

	projects, err := projectService.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, projects, "empty response")
	assert.Equal(t, 3, len(projects))
}

func Test_Project_GetOne(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/project/3440642"
	useFileResponse(t, endpoint, "project/one.json")

	var projectService = testClient.Project

	project, err := projectService.GetOne(3440642)
	assert.Nil(t, err, err)
	assert.NotNil(t, project, "empty response")
	assert.Equal(t, "PWT", project.Key)
}

func Test_Project_Search(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/project/search"
	useFileResponse(t, endpoint, "project/one.json")

	var projectService = testClient.Project

	project, err := projectService.Search("PWT")
	assert.Nil(t, err, err)
	assert.NotNil(t, project, "empty response")
	assert.Equal(t, 3440642, project.Id)
}

func Test_Project_Create(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/project"
	useFileResponse(t, endpoint, "project/one.json")

	var projectService = testClient.Project

	project, err := projectService.Create(&Project{
		Name: "P6PrnP0TUJG46DK8Fpyz",
		Key: "PWT",
		Description: "0YTPeDwGUfxJAHMSjG82ULqTkw1b6hhcHonRcthM",
	})
	assert.Nil(t, err, err)
	assert.NotNil(t, project, "empty response")
	assert.Equal(t, "PWT", project.Key)
	assert.Equal(t, 3440642, project.Id)
}

func Test_Project_Update(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/project/3440642"
	useEmptyResponse(endpoint)

	var projectService = testClient.Project

	err := projectService.Update(3440642, &Project{
		Name: "updated-name",
		Key: "PWT",
		Description: "0YTPeDwGUfxJAHMSjG82ULqTkw1b6hhcHonRcthM",
	})
	assert.Nil(t, err, err)
}

func Test_Project_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/project/3440642"
	useEmptyResponse(endpoint)

	var projectService = testClient.Project

	err := projectService.Delete(3440642)
	assert.Nil(t, err, err)
}
