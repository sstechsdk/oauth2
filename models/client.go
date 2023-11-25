package models

// Client client model
type Client struct {
	ID               string
	Secret           string
	Domain           string
	Public           bool
	UserID           string
	CallbackUrl      string
	ServerPublicKey  string
	ServerPrivateKey string
	ClientPublicKey  string
	AppName          string
}

// GetID client id
func (c *Client) GetID() string {
	return c.ID
}

// GetSecret client secret
func (c *Client) GetSecret() string {
	return c.Secret
}

// GetDomain client domain
func (c *Client) GetDomain() string {
	return c.Domain
}

// IsPublic public
func (c *Client) IsPublic() bool {
	return c.Public
}

// GetUserID user id
func (c *Client) GetUserID() string {
	return c.UserID
}

// GetCallbackUrl client id
func (c *Client) GetCallbackUrl() string {
	return c.CallbackUrl
}

// GetServerPublicKey client domain
func (c *Client) GetServerPublicKey() string {
	return c.ServerPublicKey
}

// GetServerPrivateKey client domain
func (c *Client) GetServerPrivateKey() string {
	return c.ServerPrivateKey
}

// GetClientPublicKey user id
func (c *Client) GetClientPublicKey() string {
	return c.ClientPublicKey
}

// GetAppName user id
func (c *Client) GetAppName() string {
	return c.AppName
}
