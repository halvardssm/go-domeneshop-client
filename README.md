# Go Domeneshop Client

This is a domeneshop client for Go following the [api specs](https://api.domeneshop.no/docs/)

## Usage

```go
import ( 
    Domeneshop "github.com/halvardssm/go-domeneshop-client/client"
)

client := Domeneshop.New("myToken", "mySecret")

res,err := client.ListDomains({Tld: ".no" })
```

## Available methods

```go
func (s *Client) CreateDnsRecord(domainId int, newRecord Dns) (*DnsCreateResponse, error)
func (s *Client) CreateHttpForward(domainId int, newForward HttpForward) (*HttpForward, error)
func (s *Client) Delete(endpoint string) ([]byte, error)
func (s *Client) DeleteDnsRecord(domainId, recordId int) (*DnsCreateResponse, error)
func (s *Client) DeleteHttpForward(domainId int, host string) (*HttpForward, error)
func (s *Client) Get(endpoint string, params interface{}) ([]byte, error)
func (s *Client) GetDnsRecordByHost(domainId int, host string) (*HttpForward, error)
func (s *Client) GetDnsRecordById(domainId, recordId int) (*Dns, error)
func (s *Client) GetDomainById(domainId int) (*Domain, error)
func (s *Client) GetInvoiceByInvoiceNumber(invoiceId int) (*Invoice, error)
func (s *Client) ListDnsRecords(domainId int, params ListDnsRecordsParams) (*[]Dns, error)
func (s *Client) ListDomains(params ListDomainsParams) (*[]Domain, error)
func (s *Client) ListHttpForwards(domainId int) (*[]HttpForward, error)
func (s *Client) ListInvoices(params ListInvoicesParams) (*[]Invoice, error)
func (s *Client) Post(endpoint string, message []byte) ([]byte, error)
func (s *Client) Put(endpoint string, message []byte) ([]byte, error)
func (s *Client) Request(options RequestOptions) ([]byte, error)
func (s *Client) UpdateDnsRecord(domainId, recordId int, updatedRecord Dns) (*DnsCreateResponse, error)
func (s *Client) UpdateHttpForward(domainId int, host string, updatedForward HttpForward) (*HttpForward, error)
```
