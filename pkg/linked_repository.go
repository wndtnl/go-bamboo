package bamboo

import (
	"context"
	"fmt"
	"net/http"
)

type LinkedRepository struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	RepositoryUrl string `json:"repositoryUrl"`
	Branch        string `json:"branch"`

	AuthType string `json:"authType"`

	// AuthType: PASSWORD
	Username string `json:"username"`
	Password string `json:"password"`

	// AuthType: SSH
	SshKey        string `json:"sshKey"`
	SshPassphrase string `json:"sshPassphrase"`

	// AuthType: [PASSWORD|SSH]_SHARED
	SharedCredentialId int `json:"sharedCredentialId"`

	// Options
	UseShallowClones     bool `json:"useShallowClones"`
	UseRemoteAgentCache  bool `json:"useRemoteAgentCache"`
	UseSubmodules        bool `json:"useSubmodules"`
	VerboseLogs          bool `json:"verboseLogs"`
	FetchWholeRepository bool `json:"fetchWholeRepository"`
	UseLFS               bool `json:"useLFS"`

	CommandTimeout int `json:"commandTimeout"`

	QuietPeriodEnabled    bool `json:"quietPeriodEnabled"`
	QuietPeriodWaitTime   int  `json:"quietPeriodWaitTime"`
	QuietPeriodMaxRetries int  `json:"quietPeriodMaxRetries"`

	FilterPattern      string `json:"filterPattern"`
	FilterPatternRegex string `json:"filterPatternRegex"`
	ChangeSetRegex     string `json:"changeSetRegex"`
}

var linkedRepositoryEndpoint = "linked_repository"

type LinkedRepositoryService struct {
	rest *Rest
}

func NewLinkedRepositoryService(rest *Rest) *LinkedRepositoryService {
	return &LinkedRepositoryService{
		rest: rest,
	}
}

func (s *LinkedRepositoryService) GetAll() ([]*LinkedRepository, error) {
	return s.GetAllWithContext(context.Background())
}

func (s *LinkedRepositoryService) GetAllWithContext(ctx context.Context) ([]*LinkedRepository, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, linkedRepositoryEndpoint, nil)
	if err != nil {
		return nil, err
	}

	linkedRepositories := []*LinkedRepository{}
	err = s.rest.Do(req, &linkedRepositories)

	return linkedRepositories, err
}

func (s *LinkedRepositoryService) GetOne(id int) (*LinkedRepository, error) {
	return s.GetOneWithContext(context.Background(), id)
}

func (s *LinkedRepositoryService) GetOneWithContext(ctx context.Context, id int) (*LinkedRepository, error) {

	endpoint := fmt.Sprintf("%s/%d", linkedRepositoryEndpoint, id)
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	linkedRepository := new(LinkedRepository)
	err = s.rest.Do(req, linkedRepository)

	return linkedRepository, err
}

func (s *LinkedRepositoryService) Create(linkedRepository *LinkedRepository) (*LinkedRepository, error) {
	return s.CreateWithContext(context.Background(), linkedRepository)
}

func (s *LinkedRepositoryService) CreateWithContext(
	ctx context.Context, linkedRepository *LinkedRepository) (*LinkedRepository, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, linkedRepositoryEndpoint, &linkedRepository)
	if err != nil {
		return nil, err
	}

	newLinkedRepository := new(LinkedRepository)
	err = s.rest.Do(req, newLinkedRepository)

	return newLinkedRepository, err
}

func (s *LinkedRepositoryService) Update(id int, linkedRepository *LinkedRepository) error {
	return s.UpdateWithContext(context.Background(), id, linkedRepository)
}

func (s *LinkedRepositoryService) UpdateWithContext(
	ctx context.Context, id int, linkedRepository *LinkedRepository) error {

	endpoint := fmt.Sprintf("%s/%d", linkedRepositoryEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPut, endpoint, &linkedRepository)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}

func (s *LinkedRepositoryService) Delete(id int) error {
	return s.DeleteWithContext(context.Background(), id)
}

func (s *LinkedRepositoryService) DeleteWithContext(ctx context.Context, id int) error {

	endpoint := fmt.Sprintf("%s/%d", linkedRepositoryEndpoint, id)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}
