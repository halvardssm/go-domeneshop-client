package client

import (
	"encoding/json"
	"fmt"
)

type Invoice struct {
	Id         int    `json:"id"`
	Type       string `json:"type"`
	Amount     int    `json:"amount"`
	Currency   string `json:"currency"`
	DueDate    string `json:"due_date"`
	IssuedDate string `json:"issued_date"`
	PaidDate   string `json:"paid_date"`
	Status     string `json:"status"`
	Url        string `json:"url"`
}

type ListInvoicesParams struct {
	Status string `url:"status,omitempty"`
}

// https://api.domeneshop.no/docs/#operation/getInvoices
func (s *Client) ListInvoices(params ListInvoicesParams) (*[]Invoice, error) {
	url := "invoice"
	bytes, err := s.Get(url, params)
	if err != nil {
		return nil, err
	}
	var data []Invoice
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// https://api.domeneshop.no/docs/#operation/findInvoiceByNumber
func (s *Client) GetInvoiceByInvoiceNumber(invoiceId int) (*Invoice, error) {
	url := fmt.Sprintf("invoice/%d", invoiceId)
	bytes, err := s.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var data Invoice
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
