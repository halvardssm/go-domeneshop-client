package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

const baseUrl string = "https://api.domeneshop.no/v0/"

type Client struct {
	Token  string
	Secret string
}

type RequestOptions struct {
	method   string
	endpoint string
	message  []byte
	params   interface{}
}

// Base Request func, used by higher level request interfaces
func (s *Client) Request(options RequestOptions) ([]byte, error) {
	// Creates the base request url
	reqUrl := baseUrl + options.endpoint

	// Parses the query params into an url encoded string
	v, _ := query.Values(options.params)
	params := v.Encode()

	// If url encoded string is set, it will add it to the url
	if len(params) > 0 {
		reqUrl += "?" + params
	}

	fmt.Println(reqUrl)

	// Creates a request
	req, err := http.NewRequest(options.method, reqUrl, bytes.NewBuffer(options.message))

	// Checks for errors
	if err != nil {
		return nil, err
	}

	// Sets authentication
	req.SetBasicAuth(s.Token, s.Secret)

	// Creates client and performs request
	client := &http.Client{}
	res, err := client.Do(req)

	// Checks for errors
	if err != nil {
		return nil, err
	}

	// Fetches response
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	// Checks for errors
	if err != nil {
		return nil, err
	}

	// Checks for errors
	if res.StatusCode < 200 || 300 <= res.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	// If everything succeeded, it will retur the body
	return body, nil
}

// Performs a GET request
func (s *Client) Get(endpoint string, params interface{}) ([]byte, error) {
	return s.Request(RequestOptions{
		method:   "GET",
		endpoint: endpoint,
		message:  nil,
		params:   params,
	})
}

// Performs a POST request
func (s *Client) Post(endpoint string, message []byte) ([]byte, error) {
	return s.Request(RequestOptions{
		method:   "POST",
		endpoint: endpoint,
		message:  message,
		params:   nil,
	})
}

// Performs a PUT request
func (s *Client) Put(endpoint string, message []byte) ([]byte, error) {
	return s.Request(RequestOptions{
		method:   "PUT",
		endpoint: endpoint,
		message:  message,
		params:   nil,
	})
}

// Performs a DELETE request
func (s *Client) Delete(endpoint string) ([]byte, error) {
	return s.Request(RequestOptions{
		method:   "DELETE",
		endpoint: endpoint,
		message:  nil,
		params:   nil,
	})
}
