package bamboo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type BasicAuthCredentials struct {
	Username string
	Password string
}

type Rest struct {
	httpClient  *http.Client
	baseURL     *url.URL
	credentials *BasicAuthCredentials
}

func NewBasicAuthClient(baseURL, username, password string) (*Rest, error) {

	httpClient := &http.Client{
		Timeout: time.Minute * 5,
	}

	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(parsedUrl.Path, "/rest/tpb/1.0/") {
		parsedUrl.Path += "/rest/tpb/1.0/"
	}

	client := &Rest{
		httpClient: httpClient,
		baseURL:    parsedUrl,
		credentials: &BasicAuthCredentials{
			Username: username,
			Password: password,
		},
	}

	return client, nil
}

func (r *Rest) NewRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	return r.NewRequestWithContext(context.Background(), method, endpoint, body)
}

func (r *Rest) NewRequestWithContext(
	ctx context.Context, method, endpoint string, body interface{}) (*http.Request, error) {

	// Relative endpoint
	relEndpoint, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	relEndpoint.Path = strings.TrimLeft(relEndpoint.Path, "/")

	// Absolute endpoint
	absEndpoint := r.baseURL.ResolveReference(relEndpoint).String()

	// Body
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Request
	req, err := http.NewRequestWithContext(ctx, method, absEndpoint, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Authentication
	if r.credentials.Username != "" {
		req.SetBasicAuth(r.credentials.Username, r.credentials.Password)
	}

	return req, nil
}

func (r *Rest) Do(req *http.Request, payload interface{}) error {

	res, err := r.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if !isSuccess(res) {
		return toError(res)
	}

	return toPayload(res, payload)
}

func isSuccess(r *http.Response) bool {
	status := r.StatusCode
	return 200 <= status && status <= 299
}

func toPayload(res *http.Response, payload interface{}) error {
	var err error
	if payload != nil {
		err = json.NewDecoder(res.Body).Decode(payload)

	}
	return err
}

type Error interface {
	Status() int
	error
}

type apiError struct {
	StatusCode int
	Content string
}

func toError(res *http.Response) error {
	a := &apiError{}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	content := string(bodyBytes)

	a.Content = content
	a.StatusCode = res.StatusCode

	return a
}

func (a *apiError) Error() string {
	if a.Content == "" {
		return fmt.Sprintf("request failed with status %d", a.StatusCode)
	} else {
		return fmt.Sprintf("request failed with status %d and content: %s", a.StatusCode, a.Content)
	}
}

func (a *apiError) Status() int {
	return a.StatusCode
}
