package bamboo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type LocalAgent struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

var localAgentEndpoint = "local_agent"

type LocalAgentService struct {
	rest *Rest
}

func NewLocalAgentService(rest *Rest) *LocalAgentService {
	return &LocalAgentService{
		rest: rest,
	}
}

func (s *LocalAgentService) GetAll() ([]*LocalAgent, error) {
	return s.GetAllWithContext(context.Background())
}

func (s *LocalAgentService) GetAllWithContext(ctx context.Context) ([]*LocalAgent, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, localAgentEndpoint, nil)
	if err != nil {
		return nil, err
	}

	localAgents := []*LocalAgent{}
	err = s.rest.Do(req, &localAgents)

	return localAgents, err
}

func (s *LocalAgentService) GetOne(id int) (*LocalAgent, error) {
	return s.GetOneWithContext(context.Background(), id)
}

func (s *LocalAgentService) GetOneWithContext(ctx context.Context, id int) (*LocalAgent, error) {

	endpoint := fmt.Sprintf("%s/%d", localAgentEndpoint, id)
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	localAgent := new(LocalAgent)
	err = s.rest.Do(req, localAgent)

	return localAgent, err
}

func (s *LocalAgentService) Search(name string) (*LocalAgent, error) {
	return s.SearchWithContext(context.Background(), name)
}

func (s *LocalAgentService) SearchWithContext(ctx context.Context, name string) (*LocalAgent, error) {

	endpoint := fmt.Sprintf("%s/search?name=%s", localAgentEndpoint, url.QueryEscape(name))
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	localAgent := new(LocalAgent)
	err = s.rest.Do(req, localAgent)

	return localAgent, err
}

func (s *LocalAgentService) Create(localAgent *LocalAgent) (*LocalAgent, error) {
	return s.CreateWithContext(context.Background(), localAgent)
}

func (s *LocalAgentService) CreateWithContext(ctx context.Context, localAgent *LocalAgent) (*LocalAgent, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, localAgentEndpoint, &localAgent)
	if err != nil {
		return nil, err
	}

	newAgent := new(LocalAgent)
	err = s.rest.Do(req, newAgent)

	return newAgent, err
}

func (s *LocalAgentService) Update(id int, localAgent *LocalAgent) error {
	return s.UpdateWithContext(context.Background(), id, localAgent)
}

func (s *LocalAgentService) UpdateWithContext(ctx context.Context, id int, localAgent *LocalAgent) error {

	endpoint := fmt.Sprintf("%s/%d", localAgentEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPut, endpoint, &localAgent)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}

func (s *LocalAgentService) Delete(id int) error {
	return s.DeleteWithContext(context.Background(), id)
}

func (s *LocalAgentService) DeleteWithContext(ctx context.Context, id int) error {

	endpoint := fmt.Sprintf("%s/%d", localAgentEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}