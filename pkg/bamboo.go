package bamboo

const (
	DefaultBaseURL  = "http://localhost:6990/bamboo"
	DefaultUsername = "admin"
	DefaultPassword = "admin"
)

type Client struct {
	rest *Rest

	GlobalVariable *GlobalVariableService
	LocalAgent *LocalAgentService
}

func NewDefaultClient() (*Client, error) {
	return NewClient(DefaultBaseURL, DefaultUsername, DefaultPassword)
}

func NewClient(baseURL, username, password string) (*Client, error) {

	rest, err := NewBasicAuthClient(baseURL, username, password)
	if err != nil {
		return nil, err
	}

	bamboo := &Client{
		rest:           rest,
		GlobalVariable: NewGlobalVariableService(rest),
		LocalAgent:     NewLocalAgentService(rest),
	}

	return bamboo, nil
}
