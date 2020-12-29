package bamboo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type LinkedRepositoryPermission struct {
	RepositoryId string   `json:"repositoryId"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Permissions  []string `json:"permissions"`
}

var linkedRepositoryPermissionEndpoint = "linked_repository_permission"

type LinkedRepositoryPermissionService struct {
	rest *Rest
}

func NewLinkedRepositoryPermissionService(rest *Rest) *LinkedRepositoryPermissionService {
	return &LinkedRepositoryPermissionService{
		rest: rest,
	}
}

func (s *LinkedRepositoryPermissionService) GetAll(repositoryId string) ([]*LinkedRepositoryPermission, error) {
	return s.GetAllWithContext(context.Background(), repositoryId)
}

func (s *LinkedRepositoryPermissionService) GetAllWithContext(
	ctx context.Context, repositoryId string) ([]*LinkedRepositoryPermission, error) {

	endpoint := fmt.Sprintf("%s/%s", linkedRepositoryPermissionEndpoint, repositoryId)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	linkedRepositoryPermissions := []*LinkedRepositoryPermission{}
	err = s.rest.Do(req, &linkedRepositoryPermissions)

	return linkedRepositoryPermissions, err
}

func (s *LinkedRepositoryPermissionService) GetOne(repositoryId, permissionType, name string) (*LinkedRepositoryPermission, error) {
	return s.GetOneWithContext(context.Background(), repositoryId, permissionType, name)
}

func (s *LinkedRepositoryPermissionService) GetOneWithContext(
	ctx context.Context, repositoryId, permissionType, name string) (*LinkedRepositoryPermission, error) {

	endpoint := fmt.Sprintf("%s/%s/%s/%s",
		linkedRepositoryPermissionEndpoint, repositoryId, permissionType, url.PathEscape(name))
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	linkedRepositoryPermission := new(LinkedRepositoryPermission)
	err = s.rest.Do(req, linkedRepositoryPermission)

	return linkedRepositoryPermission, err
}

func (s *LinkedRepositoryPermissionService) Upsert(
	linkedRepositoryPermission *LinkedRepositoryPermission) (*LinkedRepositoryPermission, error) {
	return s.UpsertWithContext(context.Background(), linkedRepositoryPermission)
}

func (s *LinkedRepositoryPermissionService) UpsertWithContext(
	ctx context.Context, linkedRepositoryPermission *LinkedRepositoryPermission) (*LinkedRepositoryPermission, error) {

	req, err := s.rest.NewRequestWithContext(
		ctx, http.MethodPost, linkedRepositoryPermissionEndpoint, &linkedRepositoryPermission)
	if err != nil {
		return nil, err
	}

	upsertedLinkedRepositoryPermission := new(LinkedRepositoryPermission)
	err = s.rest.Do(req, upsertedLinkedRepositoryPermission)

	return upsertedLinkedRepositoryPermission, err
}

func (s *LinkedRepositoryPermissionService) Delete(repositoryId, permissionType, name string) error {
	return s.DeleteWithContext(context.Background(), repositoryId, permissionType, name)
}

func (s *LinkedRepositoryPermissionService) DeleteWithContext(ctx context.Context, repositoryId, permissionType, name string) error {

	endpoint := fmt.Sprintf(
		"%s/%s/%s/%s", linkedRepositoryPermissionEndpoint, repositoryId, permissionType, url.PathEscape(name))

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}
