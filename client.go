package bamboo

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultBaseURL = "http://localhost:6990/bamboo/rest/api/latest/"
)

type SimpleCredentials struct {
	Username string
	Password string
}

type Client struct {
	client      *http.Client
	BaseURL     *url.URL
	SimpleCreds *SimpleCredentials
}

func (c *Client) SetURL(desiredURL string) error {
	newURL, err := url.Parse(desiredURL)
	if err != nil {
		return err
	}

	if newURL.Scheme == "" {
		return errors.New("URL scheme was blank")
	}

	if !strings.HasSuffix(newURL.Path, "/rest/api/latest/") {
		newURL.Path += "/rest/api/latest/"
	}
	c.BaseURL = newURL
	return nil
}

func NewSimpleClient(httpClient *http.Client, username, password string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	client := &Client{
		client: httpClient,
		BaseURL: baseURL,
		SimpleCreds: &SimpleCredentials{
			Username: username,
			Password: password,
		},
	}

	return client;
}