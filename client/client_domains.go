package client

import (
	"encoding/json"
	"fmt"
)

type DomainServices struct {
	Registrar bool   `json:"registrar"`
	Dns       bool   `json:"dns"`
	Email     bool   `json:"email"`
	Webhotel  string `json:"webhotel"`
}

type Domain struct {
	Id             int            `json:"id"`
	Domain         string         `json:"domain"`
	ExpiryDate     string         `json:"expiry_date"`
	RegisteredDate string         `json:"registered_date"`
	Renew          bool           `json:"renew"`
	Registrant     string         `json:"registrant"`
	Status         string         `json:"status"`
	Nameservers    []string       `json:"nameservers"`
	Services       DomainServices `json:"services"`
}

type ListDomainsParams struct {
	Tld string `url:"tld,omitempty"`
}

// https://api.domeneshop.no/docs/#operation/getDomains
func (s *Client) ListDomains(params ListDomainsParams) (*[]Domain, error) {
	bytes, err := s.Get("domains", params)
	if err != nil {
		return nil, err
	}
	var data []Domain
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#tag/domains/paths/~1domains~1{domainId}/get
func (s *Client) GetDomainById(domainId int) (*Domain, error) {
	url := fmt.Sprintf("domains/%d", domainId)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data Domain
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
