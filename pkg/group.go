package bamboo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Group struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

var groupEndpoint = "group"

type GroupService struct {
	rest *Rest
}

func NewGroupService(rest *Rest) *GroupService {
	return &GroupService{
		rest: rest,
	}
}

func (s *GroupService) GetAll() ([]*Group, error) {
	return s.GetAllWithContext(context.Background())
}

func (s *GroupService) GetAllWithContext(ctx context.Context) ([]*Group, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, groupEndpoint, nil)
	if err != nil {
		return nil, err
	}

	groups := []*Group{}
	err = s.rest.Do(req, &groups)

	return groups, err
}

func (s *GroupService) GetOne(name string) (*Group, error) {
	return s.GetOneWithContext(context.Background(), name)
}

func (s *GroupService) GetOneWithContext(ctx context.Context, name string) (*Group, error) {

	endpoint := fmt.Sprintf("%s/%s", groupEndpoint, url.PathEscape(name))
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	group := new(Group)
	err = s.rest.Do(req, group)

	return group, err
}

func (s *GroupService) Create(group *Group) (*Group, error) {
	return s.CreateWithContext(context.Background(), group)
}

func (s *GroupService) CreateWithContext(ctx context.Context, group *Group) (*Group, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, groupEndpoint, &group)
	if err != nil {
		return nil, err
	}

	newGroup := new(Group)
	err = s.rest.Do(req, newGroup)

	return newGroup, err
}

func (s *GroupService) Update(name string, group *Group) error {
	return s.UpdateWithContext(context.Background(), name, group)
}

func (s *GroupService) UpdateWithContext(ctx context.Context, name string, group *Group) error {

	endpoint := fmt.Sprintf("%s/%s", groupEndpoint, url.PathEscape(name))

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPut, endpoint, &group)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}

func (s *GroupService) Delete(name string) error {
	return s.DeleteWithContext(context.Background(), name)
}

func (s *GroupService) DeleteWithContext(ctx context.Context, name string) error {

	endpoint := fmt.Sprintf("%s/%s", groupEndpoint, name)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}
