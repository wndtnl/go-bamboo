package services

import (
	"context"
	"fmt"
	"github.com/wndtnl/go-bamboo/pkg/client"
	"net/http"
)

type GlobalVariable struct {
	Id    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

var globalVariableEndpoint = "global_variable"

type GlobalVariableService struct {
	rest *client.Rest
}

func NewGlobalVariableService(rest *client.Rest) *GlobalVariableService {
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

func (s *GlobalVariableService) GetOne(id int) (*GlobalVariable, error) {
	return s.GetOneWithContext(context.Background(), id)
}

func (s *GlobalVariableService) GetOneWithContext(ctx context.Context, id int) (*GlobalVariable, error) {

	endpoint := fmt.Sprintf("%s/%d", globalVariableEndpoint, id)
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

	endpoint := fmt.Sprintf("%s/search?key=%s", globalVariableEndpoint, key)
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	globalVariable := new(GlobalVariable)
	err = s.rest.Do(req, globalVariable)

	return globalVariable, err
}

func (s *GlobalVariableService) Create(key, value string) (*GlobalVariable, error) {
	return s.CreateWithContext(context.Background(), key, value)
}

func (s *GlobalVariableService) CreateWithContext(ctx context.Context, key, value string) (*GlobalVariable, error) {

	variable := &GlobalVariable{
		Key:   key,
		Value: value,
	}

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, globalVariableEndpoint, &variable)
	if err != nil {
		return nil, err
	}

	newVariable := new(GlobalVariable)
	err = s.rest.Do(req, newVariable)

	return newVariable, err
}

func (s *GlobalVariableService) Update(id int, key, value string) error {
	return s.UpdateWithContext(context.Background(), id, key, value)
}

func (s *GlobalVariableService) UpdateWithContext(ctx context.Context, id int, key, value string) error {

	variable := &GlobalVariable{
		Id:    id,
		Key:   key,
		Value: value,
	}

	endpoint := fmt.Sprintf("%s/%d", globalVariableEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPut, endpoint, &variable)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}

func (s *GlobalVariableService) Delete(id int) error {
	return s.DeleteWithContext(context.Background(), id)
}

func (s *GlobalVariableService) DeleteWithContext(ctx context.Context, id int) error {

	endpoint := fmt.Sprintf("%s/%d", globalVariableEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}
