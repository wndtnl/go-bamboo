package bamboo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type SharedCredential struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	// Password
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	// Aws
	AccessKey string `json:"accessKey,omitempty"`
	SecretKey string `json:"secretKey,omitempty"`

	// Ssh
	SshKey        string `json:"sshKey,omitempty"`
	SshPassphrase string `json:"sshPassphrase,omitempty"`
}

var sharedCredentialEndpoint = "shared_credential"

type SharedCredentialService struct {
	rest *Rest
}

func NewSharedCredentialService(rest *Rest) *SharedCredentialService {
	return &SharedCredentialService{
		rest: rest,
	}
}

func (s *SharedCredentialService) GetAll() ([]*SharedCredential, error) {
	return s.GetAllWithContext(context.Background())
}

func (s *SharedCredentialService) GetAllWithContext(ctx context.Context) ([]*SharedCredential, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, sharedCredentialEndpoint, nil)
	if err != nil {
		return nil, err
	}

	sharedCredentials := []*SharedCredential{}
	err = s.rest.Do(req, &sharedCredentials)

	return sharedCredentials, err
}

func (s *SharedCredentialService) GetOne(id int) (*SharedCredential, error) {
	return s.GetOneWithContext(context.Background(), id)
}

func (s *SharedCredentialService) GetOneWithContext(ctx context.Context, id int) (*SharedCredential, error) {

	endpoint := fmt.Sprintf("%s/%d", sharedCredentialEndpoint, id)
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	sharedCredential := new(SharedCredential)
	err = s.rest.Do(req, sharedCredential)

	return sharedCredential, err
}

func (s *SharedCredentialService) Search(key string) (*SharedCredential, error) {
	return s.SearchWithContext(context.Background(), key)
}

func (s *SharedCredentialService) SearchWithContext(ctx context.Context, key string) (*SharedCredential, error) {

	endpoint := fmt.Sprintf("%s/search?name=%s", sharedCredentialEndpoint, url.QueryEscape(key))
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	sharedCredential := new(SharedCredential)
	err = s.rest.Do(req, sharedCredential)

	return sharedCredential, err
}

func (s *SharedCredentialService) Create(sharedCredential *SharedCredential) (*SharedCredential, error) {
	return s.CreateWithContext(context.Background(), sharedCredential)
}

func (s *SharedCredentialService) CreateWithContext(
	ctx context.Context, sharedCredential *SharedCredential) (*SharedCredential, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, sharedCredentialEndpoint, &sharedCredential)
	if err != nil {
		return nil, err
	}

	newSharedCredential := new(SharedCredential)
	err = s.rest.Do(req, newSharedCredential)

	return newSharedCredential, err
}

func (s *SharedCredentialService) Update(id int, sharedCredential *SharedCredential) error {
	return s.UpdateWithContext(context.Background(), id, sharedCredential)
}

func (s *SharedCredentialService) UpdateWithContext(
	ctx context.Context, id int, sharedCredential *SharedCredential) error {

	endpoint := fmt.Sprintf("%s/%d", sharedCredentialEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPut, endpoint, &sharedCredential)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}

func (s *SharedCredentialService) Delete(id int) error {
	return s.DeleteWithContext(context.Background(), id)
}

func (s *SharedCredentialService) DeleteWithContext(ctx context.Context, id int) error {

	endpoint := fmt.Sprintf("%s/%d", sharedCredentialEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}