package bamboo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LocalAgent_GetAll(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/local_agent"
	useFileResponse(t, endpoint, "local_agent/all.json")

	var localAgentService = testClient.LocalAgent

	agents, err := localAgentService.GetAll()
	assert.Nil(t, err, err)
	assert.NotNil(t, agents, "empty response")
	assert.Equal(t, 3, len(agents))
}

func Test_LocalAgent_GetOne(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/local_agent/3375106"
	useFileResponse(t, endpoint, "local_agent/one.json")

	var localAgentService = testClient.LocalAgent

	agent, err := localAgentService.GetOne("3375106")
	assert.Nil(t, err, err)
	assert.NotNil(t, agent, "empty response")
	assert.Equal(t, "Agent1", agent.Name)
}

func Test_LocalAgent_Search(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/local_agent/search"
	useFileResponse(t, endpoint, "local_agent/one.json")

	var localAgentService = testClient.LocalAgent

	agent, err := localAgentService.Search("Agent1")
	assert.Nil(t, err, err)
	assert.NotNil(t, agent, "empty response")
	assert.Equal(t, "3375106", agent.Id)
}

func Test_LocalAgent_Create(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/local_agent"
	useFileResponse(t, endpoint, "local_agent/create.json")

	var localAgentService = testClient.LocalAgent

	agent, err := localAgentService.Create(&LocalAgent{
		Name:        "Local Agent 3",
		Description: "Managed By API",
		Enabled:     false,
	})
	assert.Nil(t, err, err)
	assert.NotNil(t, agent, "empty response")
	assert.Equal(t, "Local Agent 3", agent.Name)
	assert.Equal(t, "Managed By API", agent.Description)
}

func Test_LocalAgent_Update(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/local_agent/3375106"
	useEmptyResponse(endpoint)

	var localAgentService = testClient.LocalAgent

	err := localAgentService.Update("3375106", &LocalAgent{
		Id:          "3375106",
		Name:        "Agent2",
		Description: "New Description",
		Enabled:     true,
	})
	assert.Nil(t, err, err)
}

func Test_LocalAgent_Delete(t *testing.T) {
	setup()
	defer teardown()

	endpoint := "/rest/tpb/1.0/local_agent/3375106"
	useEmptyResponse(endpoint)

	var localAgentService = testClient.LocalAgent

	err := localAgentService.Delete("3375106")
	assert.Nil(t, err, err)
}
