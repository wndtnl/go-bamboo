package bamboo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type User struct {
	Username      string `json:"username"`
	FullName      string `json:"fullName"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	JabberAddress string `json:"jabberAddress"`
	Enabled       bool   `json:"enabled"`
}

var userEndpoint = "user"

type UserService struct {
	rest *Rest
}

func NewUserService(rest *Rest) *UserService {
	return &UserService{
		rest: rest,
	}
}

func (s *UserService) GetAll() ([]*User, error) {
	return s.GetAllWithContext(context.Background())
}

func (s *UserService) GetAllWithContext(ctx context.Context) ([]*User, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, userEndpoint, nil)
	if err != nil {
		return nil, err
	}

	users := []*User{}
	err = s.rest.Do(req, &users)

	return users, err
}

func (s *UserService) GetOne(name string) (*User, error) {
	return s.GetOneWithContext(context.Background(), name)
}

func (s *UserService) GetOneWithContext(ctx context.Context, name string) (*User, error) {

	endpoint := fmt.Sprintf("%s/%s", userEndpoint, url.PathEscape(name))
	req, err := s.rest.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	user := new(User)
	err = s.rest.Do(req, user)

	return user, err
}

func (s *UserService) Create(user *User) (*User, error) {
	return s.CreateWithContext(context.Background(), user)
}

func (s *UserService) CreateWithContext(ctx context.Context, user *User) (*User, error) {

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPost, userEndpoint, &user)
	if err != nil {
		return nil, err
	}

	newUser := new(User)
	err = s.rest.Do(req, newUser)

	return newUser, err
}

func (s *UserService) Update(name string, user *User) error {
	return s.UpdateWithContext(context.Background(), name, user)
}

func (s *UserService) UpdateWithContext(ctx context.Context, name string, user *User) error {

	endpoint := fmt.Sprintf("%s/%s", userEndpoint, url.PathEscape(name))

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodPut, endpoint, &user)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}

func (s *UserService) Delete(name string) error {
	return s.DeleteWithContext(context.Background(), name)
}

func (s *UserService) DeleteWithContext(ctx context.Context, name string) error {

	endpoint := fmt.Sprintf("%s/%s", userEndpoint, name)

	req, err := s.rest.NewRequestWithContext(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	return s.rest.Do(req, nil)
}