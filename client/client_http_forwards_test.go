package client

import (
	"testing"
)

func isHttpForward(v interface{}) bool {
	_, ok := v.(HttpForward)

	return ok
}

// Tests the ListHttpForwards function
func TestListHttpForwards(t *testing.T) {
	client := getTestClient()

	res, err := client.ListHttpForwards(domainId)

	if err != nil {
		t.Error("Did not expect error")
		t.FailNow()
	}

	if len(*res) < 1 {
		t.Error("Expected lenght to be larger than 1")
	}

	if !isHttpForward((*res)[0]) {
		t.Error("Result was not of type HttpForward")
	}
}

// Tests the CreateHttpForward function
func TestCreateHttpForward(t *testing.T) {
	client := getTestClient()

	res, err := client.CreateHttpForward(domainId, HttpForward{})

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isHttpForward(*res) {
		t.Error("Result was not of type HttpForward")
	}
}

// Tests the GetHttpForwardByHost function
func TestGetHttpForwardByHost(t *testing.T) {
	client := getTestClient()

	res, err := client.GetHttpForwardByHost(domainId, domainName)

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isHttpForward(*res) {
		t.Error("Result was not of type HttpForward")
	}
}

// Tests the UpdateHttpForward function
func TestUpdateHttpForward(t *testing.T) {
	client := getTestClient()

	res, err := client.UpdateHttpForward(domainId, domainName, HttpForward{})

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isHttpForward(*res) {
		t.Error("Result was not of type HttpForward")
	}
}

// Tests the DeleteHttpForward function
func TestDeleteHttpForward(t *testing.T) {
	client := getTestClient()

	res, err := client.DeleteHttpForward(domainId, domainName)

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isHttpForward(*res) {
		t.Error("Result was not of type HttpForward")
	}
}
