package client

import (
// "os"
)

// Replace these value if you are testing locally, with a domain you own
const domainId = 1236524
const recordId = 1236524
const domainName = "example.com"
const invoiceId = 1234

func getTestClient() *Client {
	token := "OjnLOx0HtOr2iiLD"                                                  //os.Getenv("DOMENESHOP_TOKEN")
	secret := "fhpDZVdbe0GhVfixH5a4LxvFTvZOvWxMRSBLlmCKKlORRixMWwkfnEmwazVskWe5" //os.Getenv("DOMENESHOP_SECRET")
	client := New(token, secret)

	return client
}
