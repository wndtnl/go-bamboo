package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DeploymentProject_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/deployment_project"
	useFileResponse(t, endpoint, "deployment_project/all.json")

	var deploymentProjectService = testClient.DeploymentProject

	deploymentProjects, err := deploymentProjectService.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, deploymentProjects, "empty response")
	assert.Equal(t, 2, len(deploymentProjects))
}

func Test_DeploymentProject_GetOne(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/deployment_project/4063233"
	useFileResponse(t, endpoint, "deployment_project/one.json")

	var deploymentProjectService = testClient.DeploymentProject

	deploymentProject, err := deploymentProjectService.GetOne("4063233")
	assert.Nil(t, err, err)
	assert.NotNil(t, deploymentProject, "empty response")
	assert.Equal(t, "Website", deploymentProject.Name)
}

func Test_DeploymentProject_Create(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/deployment_project"
	useFileResponse(t, endpoint, "deployment_project/one.json")

	var deploymentProjectService = testClient.DeploymentProject

	deploymentProject, err := deploymentProjectService.Create(&DeploymentProject{
		Name:        "Website",
		Description: "",
		PlanKey:     "KOU-WEB",
	})
	assert.Nil(t, err, err)
	assert.NotNil(t, deploymentProject, "empty response")
	assert.Equal(t, "Website", deploymentProject.Name)
	assert.Equal(t, "4063233", deploymentProject.Id)
}

func Test_DeploymentProject_Update(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/deployment_project/4063233"
	useEmptyResponse(endpoint)

	var deploymentProjectService = testClient.DeploymentProject

	err := deploymentProjectService.Update("4063233", &DeploymentProject{
		Id:          "4063233",
		Name:        "updated-name",
		Description: "",
		PlanKey:     "KOU-WEB",
	})
	assert.Nil(t, err, err)
}

func Test_DeploymentProject_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/deployment_project/4063233"
	useEmptyResponse(endpoint)

	var deploymentProjectService = testClient.DeploymentProject

	err := deploymentProjectService.Delete("4063233")
	assert.Nil(t, err, err)
}
