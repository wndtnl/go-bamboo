package bamboo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type GlobalPermission struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Permissions []string `json:"permissions"`
}

var globalPermissionEndpoint = "global_permission"

type GlobalPermissionService struct {
	rest *Rest
}

func NewGlobalPermissionService(rest *Rest) *GlobalPermissionService {
	return &GlobalPermissionService{
		rest: rest,
	}
}

func (s *GlobalPermissionService) GetAll() ([]*GlobalPermission, error) {
	return s.GetAllWithContext(context.Background())
}

func (s *GlobalPermissionService) GetAllWithContext(ctx context.Context) ([]*GlobalPermission, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, globalPermissionEndpoint, nil)
	if err != nil {
		return nil, err
	}

	globalPermissions := []*GlobalPermission{}
	err = s.rest.Do(req, &globalPermissions)

	return globalPermissions, err
}

func (s *GlobalPermissionService) GetOne(permissionType, name string) (*GlobalPermission, error) {
	return s.GetOneWithContext(context.Background(), permissionType, name)
}

func (s *GlobalPermissionService) GetOneWithContext(
	ctx context.Context, permissionType, name string) (*GlobalPermission, error) {

	endpoint := fmt.Sprintf("%s/%s/%s", globalPermissionEndpoint, permissionType, url.PathEscape(name))
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	globalPermission := new(GlobalPermission)
	err = s.rest.Do(req, globalPermission)

	return globalPermission, err
}

func (s *GlobalPermissionService) Upsert(globalPermission *GlobalPermission) (*GlobalPermission, error) {
	return s.UpsertWithContext(context.Background(), globalPermission)
}

func (s *GlobalPermissionService) UpsertWithContext(
	ctx context.Context, globalPermission *GlobalPermission) (*GlobalPermission, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, globalPermissionEndpoint, &globalPermission)
	if err != nil {
		return nil, err
	}

	upsertedGlobalPermission := new(GlobalPermission)
	err = s.rest.Do(req, upsertedGlobalPermission)

	return upsertedGlobalPermission, err
}

func (s *GlobalPermissionService) Delete(permissionType, name string) error {
	return s.DeleteWithContext(context.Background(), permissionType, name)
}

func (s *GlobalPermissionService) DeleteWithContext(ctx context.Context, permissionType, name string) error {

	endpoint := fmt.Sprintf("%s/%s/%s", globalPermissionEndpoint, permissionType, url.PathEscape(name))

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}
