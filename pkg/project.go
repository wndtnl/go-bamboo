package bamboo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Key         string `json:"key"`
	Description string `json:"description"`
}

var projectEndpoint = "project"

type ProjectService struct {
	rest *Rest
}

func NewProjectService(rest *Rest) *ProjectService {
	return &ProjectService{
		rest: rest,
	}
}

func (s *ProjectService) GetAll() ([]*Project, error) {
	return s.GetAllWithContext(context.Background())
}

func (s *ProjectService) GetAllWithContext(ctx context.Context) ([]*Project, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, projectEndpoint, nil)
	if err != nil {
		return nil, err
	}

	projects := []*Project{}
	err = s.rest.Do(req, &projects)

	return projects, err
}

func (s *ProjectService) GetOne(id int) (*Project, error) {
	return s.GetOneWithContext(context.Background(), id)
}

func (s *ProjectService) GetOneWithContext(ctx context.Context, id int) (*Project, error) {

	endpoint := fmt.Sprintf("%s/%d", projectEndpoint, id)
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	project := new(Project)
	err = s.rest.Do(req, project)

	return project, err
}

func (s *ProjectService) Search(key string) (*Project, error) {
	return s.SearchWithContext(context.Background(), key)
}

func (s *ProjectService) SearchWithContext(ctx context.Context, key string) (*Project, error) {

	endpoint := fmt.Sprintf("%s/search?key=%s", projectEndpoint, url.QueryEscape(key))
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	project := new(Project)
	err = s.rest.Do(req, project)

	return project, err
}

func (s *ProjectService) Create(project *Project) (*Project, error) {
	return s.CreateWithContext(context.Background(), project)
}

func (s *ProjectService) CreateWithContext(ctx context.Context, project *Project) (*Project, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, projectEndpoint, &project)
	if err != nil {
		return nil, err
	}

	newProject := new(Project)
	err = s.rest.Do(req, newProject)

	return newProject, err
}

func (s *ProjectService) Update(id int, project *Project) error {
	return s.UpdateWithContext(context.Background(), id, project)
}

func (s *ProjectService) UpdateWithContext(ctx context.Context, id int, project *Project) error {

	endpoint := fmt.Sprintf("%s/%d", projectEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPut, endpoint, &project)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}

func (s *ProjectService) Delete(id int) error {
	return s.DeleteWithContext(context.Background(), id)
}

func (s *ProjectService) DeleteWithContext(ctx context.Context, id int) error {

	endpoint := fmt.Sprintf("%s/%d", projectEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}
