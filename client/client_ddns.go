package client

import (
	"encoding/json"
)

type UpdateDDnsParams struct {
	HostName string `url:"hostname"`
	MyIp     string `url:"myip,omitempty"`
}

// https://api.domeneshop.no/docs/#tag/ddns/paths/~1dyndns~1update/get
// As the api is not ready, assume that it is successfull if no error is returned
func (s *Client) UpdateDDns(params UpdateDDnsParams) (*interface{}, error) {
	bytes, err := s.Get("dyndns/update", params)
	if err != nil {
		return nil, err
	}
	var data interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
