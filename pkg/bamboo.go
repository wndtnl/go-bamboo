package bamboo

import (
	"github.com/wndtnl/go-bamboo/pkg/client"
	"github.com/wndtnl/go-bamboo/pkg/services"
)

const (
	defaultBaseURL  = "http://localhost:6990/bamboo"
	defaultUsername = "admin"
	defaultPassword = "admin"
)

type Client struct {
	rest *client.Rest

	GlobalVariable *services.GlobalVariableService
}

func NewDefaultClient() (*Client, error) {
	return NewClient(defaultBaseURL, defaultUsername, defaultPassword)
}

func NewClient(baseURL, username, password string) (*Client, error) {

	rest, err := client.NewBasicAuthClient(baseURL, username, password)
	if err != nil {
		return nil, err
	}

	bamboo := &Client{
		rest:           rest,
		GlobalVariable: services.NewGlobalVariableService(rest),
	}

	return bamboo, nil
}
