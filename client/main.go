package client

// Creates a new Domeneshop Client
func New(token, secret string) *Client {
	return &Client{
		Token:  token,
		Secret: secret,
	}
}
