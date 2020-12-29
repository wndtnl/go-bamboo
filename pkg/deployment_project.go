package bamboo

import (
	"context"
	"fmt"
	"net/http"
)

type DeploymentProject struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PlanKey     string `json:"planKey"`
}

var deploymentProjectEndpoint = "deployment_project"

type DeploymentProjectService struct {
	rest *Rest
}

func NewDeploymentProjectService(rest *Rest) *DeploymentProjectService {
	return &DeploymentProjectService{
		rest: rest,
	}
}

func (s *DeploymentProjectService) GetAll() ([]*DeploymentProject, error) {
	return s.GetAllWithContext(context.Background())
}

func (s *DeploymentProjectService) GetAllWithContext(ctx context.Context) ([]*DeploymentProject, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, deploymentProjectEndpoint, nil)
	if err != nil {
		return nil, err
	}

	deploymentProjects := []*DeploymentProject{}
	err = s.rest.Do(req, &deploymentProjects)

	return deploymentProjects, err
}

func (s *DeploymentProjectService) GetOne(id string) (*DeploymentProject, error) {
	return s.GetOneWithContext(context.Background(), id)
}

func (s *DeploymentProjectService) GetOneWithContext(ctx context.Context, id string) (*DeploymentProject, error) {

	endpoint := fmt.Sprintf("%s/%s", deploymentProjectEndpoint, id)
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	deploymentProject := new(DeploymentProject)
	err = s.rest.Do(req, deploymentProject)

	return deploymentProject, err
}

func (s *DeploymentProjectService) Create(deploymentProject *DeploymentProject) (*DeploymentProject, error) {
	return s.CreateWithContext(context.Background(), deploymentProject)
}

func (s *DeploymentProjectService) CreateWithContext(
	ctx context.Context, deploymentProject *DeploymentProject) (*DeploymentProject, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, deploymentProjectEndpoint, &deploymentProject)
	if err != nil {
		return nil, err
	}

	newDeploymentProject := new(DeploymentProject)
	err = s.rest.Do(req, newDeploymentProject)

	return newDeploymentProject, err
}

func (s *DeploymentProjectService) Update(id string, deploymentProject *DeploymentProject) error {
	return s.UpdateWithContext(context.Background(), id, deploymentProject)
}

func (s *DeploymentProjectService) UpdateWithContext(
	ctx context.Context, id string, deploymentProject *DeploymentProject) error {

	endpoint := fmt.Sprintf("%s/%s", deploymentProjectEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPut, endpoint, &deploymentProject)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}

func (s *DeploymentProjectService) Delete(id string) error {
	return s.DeleteWithContext(context.Background(), id)
}

func (s *DeploymentProjectService) DeleteWithContext(ctx context.Context, id string) error {

	endpoint := fmt.Sprintf("%s/%s", deploymentProjectEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}
