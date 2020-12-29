package bamboo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type GlobalVariable struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

var globalVariableEndpoint = "global_variable"

type GlobalVariableService struct {
	rest *Rest
}

func NewGlobalVariableService(rest *Rest) *GlobalVariableService {
	return &GlobalVariableService{
		rest: rest,
	}
}

func (s *GlobalVariableService) GetAll() ([]*GlobalVariable, error) {
	return s.GetAllWithContext(context.Background())
}

func (s *GlobalVariableService) GetAllWithContext(ctx context.Context) ([]*GlobalVariable, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, globalVariableEndpoint, nil)
	if err != nil {
		return nil, err
	}

	globalVariables := []*GlobalVariable{}
	err = s.rest.Do(req, &globalVariables)

	return globalVariables, err
}

func (s *GlobalVariableService) GetOne(id string) (*GlobalVariable, error) {
	return s.GetOneWithContext(context.Background(), id)
}

func (s *GlobalVariableService) GetOneWithContext(ctx context.Context, id string) (*GlobalVariable, error) {

	endpoint := fmt.Sprintf("%s/%s", globalVariableEndpoint, id)
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	globalVariable := new(GlobalVariable)
	err = s.rest.Do(req, globalVariable)

	return globalVariable, err
}

func (s *GlobalVariableService) Search(key string) (*GlobalVariable, error) {
	return s.SearchWithContext(context.Background(), key)
}

func (s *GlobalVariableService) SearchWithContext(ctx context.Context, key string) (*GlobalVariable, error) {

	endpoint := fmt.Sprintf("%s/search?key=%s", globalVariableEndpoint, url.QueryEscape(key))
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	globalVariable := new(GlobalVariable)
	err = s.rest.Do(req, globalVariable)

	return globalVariable, err
}

func (s *GlobalVariableService) Create(globalVariable *GlobalVariable) (*GlobalVariable, error) {
	return s.CreateWithContext(context.Background(), globalVariable)
}

func (s *GlobalVariableService) CreateWithContext(ctx context.Context, globalVariable *GlobalVariable) (*GlobalVariable, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, globalVariableEndpoint, &globalVariable)
	if err != nil {
		return nil, err
	}

	newVariable := new(GlobalVariable)
	err = s.rest.Do(req, newVariable)

	return newVariable, err
}

func (s *GlobalVariableService) Update(id string, globalVariable *GlobalVariable) error {
	return s.UpdateWithContext(context.Background(), id, globalVariable)
}

func (s *GlobalVariableService) UpdateWithContext(ctx context.Context, id string, globalVariable *GlobalVariable) error {

	endpoint := fmt.Sprintf("%s/%s", globalVariableEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPut, endpoint, &globalVariable)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}

func (s *GlobalVariableService) Delete(id string) error {
	return s.DeleteWithContext(context.Background(), id)
}

func (s *GlobalVariableService) DeleteWithContext(ctx context.Context, id string) error {

	endpoint := fmt.Sprintf("%s/%s", globalVariableEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}
