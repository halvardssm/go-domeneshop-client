package client

import (
	"testing"
)

func isDnsRecord(v interface{}) bool {
	_, ok := v.(Dns)

	return ok
}

func isDnsSaveResponse(v interface{}) bool {
	_, ok := v.(DnsSaveResponse)

	return ok
}

// Tests the ListDomains function
func TestListDnsRecords(t *testing.T) {
	client := getTestClient()

	res, err := client.ListDnsRecords(domainId, ListDnsRecordsParams{})

	if err != nil {
		t.Error("Did not expect error")
		t.FailNow()
	}

	if len(*res) < 1 {
		t.Error("Expected lenght to be larger than 1")
	}

	if !isDnsRecord((*res)[0]) {
		t.Error("Result was not of type Dns")
	}
}

// Tests the CreateDnsRecord function
func TestCreateDnsRecord(t *testing.T) {
	client := getTestClient()

	res, err := client.CreateDnsRecord(domainId, Dns{})

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isDnsSaveResponse(res) {
		t.Error("Result was not of type DnsSaveResponse")
	}
}

// Tests the GetDnsRecordById function
func TestGetDnsRecordById(t *testing.T) {
	client := getTestClient()

	res, err := client.GetDnsRecordById(domainId, recordId)

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isDnsRecord(res) {
		t.Error("Result was not of type Dns")
	}
}

// Tests the UpdateDnsRecord function
func TestUpdateDnsRecord(t *testing.T) {
	client := getTestClient()

	res, err := client.UpdateDnsRecord(domainId, recordId, Dns{})

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isDnsSaveResponse(res) {
		t.Error("Result was not of type DnsSaveResponse")
	}
}

// Tests the DeleteDnsRecord function
func TestDeleteDnsRecord(t *testing.T) {
	client := getTestClient()

	res, err := client.DeleteDnsRecord(domainId, recordId)

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isDnsSaveResponse(res) {
		t.Error("Result was not of type DnsSaveResponse")
	}
}
