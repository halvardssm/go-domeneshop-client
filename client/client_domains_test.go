package client

import (
	"testing"
)

func isDomain(v interface{}) bool {
	_, ok := v.(Domain)

	return ok
}

// Tests the ListDomains function
func TestListDomains(t *testing.T) {
	client := getTestClient()

	res, err := client.ListDomains(ListDomainsParams{})

	if err != nil {
		t.Error("Did not expect error")
		t.FailNow()
	}

	if len(*res) < 1 {
		t.Error("Expected lenght to be larger than 1")
	}

	// t.Log((*res)[0].Id)

	if !isDomain((*res)[0]) {
		t.Error("Result was not of type Domain")
	}
}

// Test the GetDomainById function
func TestGetDomainById(t *testing.T) {
	client := getTestClient()

	res, err := client.GetDomainById(domainId)

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isDomain(*res) {
		t.Error("Result was not of type Domain")
	}
}
