package client

import (
	"bytes"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
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

func (s *Client) Request(options RequestOptions) ([]byte, error) {
	reqUrl := baseUrl + options.endpoint

	if options.params != nil {
		v, _ := query.Values(options.params)
		reqUrl += v.Encode()
	}

	var msg *bytes.Buffer
	if options.message != nil {
		msg = bytes.NewBuffer(options.message)
	}

	req, err := http.NewRequest(options.method, reqUrl, msg)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(s.Token, s.Secret)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || 300 <= res.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}

	return body, nil
}

func (s *Client) Get(endpoint string, params interface{}) ([]byte, error) {
	return s.Request(RequestOptions{
		method:   "GET",
		endpoint: endpoint,
		message:  nil,
		params:   params,
	})
}

func (s *Client) Post(endpoint string, message []byte) ([]byte, error) {
	return s.Request(RequestOptions{
		method:   "POST",
		endpoint: endpoint,
		message:  message,
		params:   nil,
	})
}

func (s *Client) Put(endpoint string, message []byte) ([]byte, error) {
	return s.Request(RequestOptions{
		method:   "PUT",
		endpoint: endpoint,
		message:  message,
		params:   nil,
	})
}

func (s *Client) Delete(endpoint string) ([]byte, error) {
	return s.Request(RequestOptions{
		method:   "DELETE",
		endpoint: endpoint,
		message:  nil,
		params:   nil,
	})
}
