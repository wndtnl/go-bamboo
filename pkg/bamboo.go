package bamboo

import (
	"github.com/wndtnl/go-bamboo/pkg/client"
	"github.com/wndtnl/go-bamboo/pkg/services"
)

const (
	DefaultBaseURL  = "http://localhost:6990/bamboo"
	DefaultUsername = "admin"
	DefaultPassword = "admin"
)

type Client struct {
	rest *client.Rest

	GlobalVariable *services.GlobalVariableService
}

func NewDefaultClient() (*Client, error) {
	return NewClient(DefaultBaseURL, DefaultUsername, DefaultPassword)
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
