package client

func New(token, secret string) *Client {
	return &Client{
		Token:  token,
		Secret: secret,
	}
}
