package client

import (
	"testing"
)

func isInvoice(v interface{}) bool {
	_, ok := v.(Invoice)

	return ok
}

// Tests the ListInvoices function
func TestListInvoices(t *testing.T) {
	client := getTestClient()

	res, err := client.ListInvoices(ListInvoicesParams{})

	if err != nil {
		t.Error("Did not expect error")
		t.FailNow()
	}

	if len(*res) < 1 {
		t.Error("Expected lenght to be larger than 1")
	}

	if !isInvoice((*res)[0]) {
		t.Error("Result was not of type Invoice")
	}
}

// Tests the GetInvoiceByInvoiceNumber function
func TestGetInvoiceByInvoiceNumber(t *testing.T) {
	client := getTestClient()

	res, err := client.GetInvoiceByInvoiceNumber(invoiceId)

	if err != nil {
		t.Error("Did not expect error")
	}

	if !isInvoice(*res) {
		t.Error("Result was not of type Invoice")
	}
}
