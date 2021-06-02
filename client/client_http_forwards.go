package client

import (
	"encoding/json"
	"fmt"
)

type HttpForward struct {
	Host  string `json:"host"`
	Frame bool   `json:"frame"`
	Url   string `json:"url"`
}

// https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1/get
func (s *Client) ListHttpForwards(domainId int) (*[]HttpForward, error) {
	url := fmt.Sprintf("domains/%d/forwards", domainId)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data []HttpForward
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1/post
func (s *Client) CreateHttpForward(domainId int, newForward HttpForward) (*HttpForward, error) {
	url := fmt.Sprintf("domains/%d/forwards", domainId)
	message, err := json.Marshal(newForward)
	if err != nil {
		return nil, err
	}
	bytes, err := s.Post(url, message)
	if err != nil {
		return nil, err
	}
	var data HttpForward
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1{host}/get
func (s *Client) GetHttpForwardByHost(domainId int, host string) (*HttpForward, error) {
	url := fmt.Sprintf("domains/%d/forwards/%s", domainId, host)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data HttpForward
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1{host}/put
func (s *Client) UpdateHttpForward(domainId int, host string, updatedForward HttpForward) (*HttpForward, error) {
	url := fmt.Sprintf("domains/%d/forwards/%s", domainId, host)
	message, err := json.Marshal(updatedForward)
	if err != nil {
		return nil, err
	}
	bytes, err := s.Put(url, message)
	if err != nil {
		return nil, err
	}
	var data HttpForward
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1{host}/delete
func (s *Client) DeleteHttpForward(domainId int, host string) (*HttpForward, error) {
	url := fmt.Sprintf("domains/%d/forwards/%s", domainId, host)
	bytes, err := s.Delete(url)
	if err != nil {
		return nil, err
	}
	var data HttpForward
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
