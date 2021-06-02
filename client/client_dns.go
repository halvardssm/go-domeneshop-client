package client

import (
	"encoding/json"
	"fmt"
)

type Dns struct {
	Host     string `json:"host"`
	Ttl      int16  `json:"ttl"`
	Type     string `json:"type"`
	Data     string `json:"data"`
	Priority int16  `json:"priority"`
	Weight   int16  `json:"weight"`
	Port     int16  `json:"port"`
}

type DnsSaveResponse struct {
	Id int `json:"id"`
}

type ListDnsRecordsParams struct {
	Host string `url:"host,omitempty"`
	Type string `url:"type,omitempty"`
}

// https://api.domeneshop.no/docs/#operation/getDnsRecords
func (s *Client) ListDnsRecords(domainId int, params ListDnsRecordsParams) (*[]Dns, error) {
	url := fmt.Sprintf("domains/%d/dns", domainId)
	bytes, err := s.Get(url, params)
	if err != nil {
		return nil, err
	}
	var data []Dns
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#tag/dns/paths/~1domains~1{domainId}~1dns/post
func (s *Client) CreateDnsRecord(domainId int, newRecord Dns) (*DnsSaveResponse, error) {
	url := fmt.Sprintf("domains/%d/dns", domainId)
	message, err := json.Marshal(newRecord)
	if err != nil {
		return nil, err
	}
	bytes, err := s.Post(url, message)
	if err != nil {
		return nil, err
	}
	var data DnsSaveResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#operation/getRecordById
func (s *Client) GetDnsRecordById(domainId, recordId int) (*Dns, error) {
	url := fmt.Sprintf("domains/%d/dns/%d", domainId, recordId)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data Dns
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#tag/dns/paths/~1domains~1{domainId}~1dns~1{recordId}/put
func (s *Client) UpdateDnsRecord(domainId, recordId int, updatedRecord Dns) (*DnsSaveResponse, error) {
	url := fmt.Sprintf("domains/%d/dns/%d", domainId, recordId)
	message, err := json.Marshal(updatedRecord)
	if err != nil {
		return nil, err
	}
	bytes, err := s.Put(url, message)
	if err != nil {
		return nil, err
	}
	var data DnsSaveResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#tag/dns/paths/~1domains~1{domainId}~1dns~1{recordId}/delete
func (s *Client) DeleteDnsRecord(domainId, recordId int) (*DnsSaveResponse, error) {
	url := fmt.Sprintf("domains/%d/dns/%d", domainId, recordId)
	bytes, err := s.Delete(url)
	if err != nil {
		return nil, err
	}
	var data DnsSaveResponse
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
