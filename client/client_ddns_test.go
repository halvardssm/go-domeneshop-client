package client

import (
	"testing"
)

// Tests the UpdateDDns function
func TestUpdateDDns(t *testing.T) {
	client := getTestClient()

	_, err := client.UpdateDDns(UpdateDDnsParams{HostName: domainName})

	if err != nil {
		t.Error("Did not expect error")
	}
}
