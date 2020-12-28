package bamboo

const (
	DefaultBaseURL  = "http://localhost:6990/bamboo"
	DefaultUsername = "admin"
	DefaultPassword = "admin"
)

type Client struct {
	rest *Rest

	GlobalVariable   *GlobalVariableService
	LocalAgent       *LocalAgentService
	SharedCredential *SharedCredentialService
	Group            *GroupService
	User             *UserService
	GlobalPermission *GlobalPermissionService
	LinkedRepository *LinkedRepositoryService
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
		rest:             rest,
		GlobalVariable:   NewGlobalVariableService(rest),
		LocalAgent:       NewLocalAgentService(rest),
		SharedCredential: NewSharedCredentialService(rest),
		Group:            NewGroupService(rest),
		User:             NewUserService(rest),
		GlobalPermission: NewGlobalPermissionService(rest),
		LinkedRepository: NewLinkedRepositoryService(rest),
	}

	return bamboo, nil
}
