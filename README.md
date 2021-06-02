![GitHub Workflow Status](https://img.shields.io/github/workflow/status/halvardssm/go-domeneshop-client/CI?logo=github&style=flat-square)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/halvardssm/go-domeneshop-client?logo=go&style=flat-square)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/halvardssm/go-domeneshop-client?logo=go&style=flat-square)

# Go Domeneshop Client

This is a domeneshop client for Go following the [API specs](https://api.domeneshop.no/docs/). 

> 🚨 As the Domeneshop API is not yet stable, this implementation can also not be considered stable. As soon as a stable release of their API is out, this will be updated accordingly. Versions will follow the Domeneshop API versions.

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
func New(token, secret string) *Client
    Creates a new Domeneshop Client

func (s *Client) CreateDnsRecord(domainId int, newRecord Dns) (*DnsSaveResponse, error)
   // https://api.domeneshop.no/docs/#tag/dns/paths/~1domains~1{domainId}~1dns/post

func (s *Client) CreateHttpForward(domainId int, newForward HttpForward) (*HttpForward, error)
    // https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1/post

func (s *Client) Delete(endpoint string) ([]byte, error)
    // Performs a DELETE request

func (s *Client) DeleteDnsRecord(domainId, recordId int) (*DnsSaveResponse, error)
    // https://api.domeneshop.no/docs/#tag/dns/paths/~1domains~1{domainId}~1dns~1{recordId}/delete

func (s *Client) DeleteHttpForward(domainId int, host string) (*HttpForward, error)
    // https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1{host}/delete

func (s *Client) Get(endpoint string, params interface{}) ([]byte, error)
    // Performs a GET request

func (s *Client) GetDnsRecordById(domainId, recordId int) (*Dns, error)
    // https://api.domeneshop.no/docs/#operation/getRecordById

func (s *Client) GetDomainById(domainId int) (*Domain, error)
    // https://api.domeneshop.no/docs/#tag/domains/paths/~1domains~1{domainId}/get

func (s *Client) GetHttpForwardByHost(domainId int, host string) (*HttpForward, error)
    // https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1{host}/get

func (s *Client) GetInvoiceByInvoiceNumber(invoiceId int) (*Invoice, error)
    // https://api.domeneshop.no/docs/#operation/findInvoiceByNumber

func (s *Client) ListDnsRecords(domainId int, params ListDnsRecordsParams) (*[]Dns, error)
    // https://api.domeneshop.no/docs/#operation/getDnsRecords

func (s *Client) ListDomains(params ListDomainsParams) (*[]Domain, error)
    // https://api.domeneshop.no/docs/#operation/getDomains

func (s *Client) ListHttpForwards(domainId int) (*[]HttpForward, error)
    // https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1/get

func (s *Client) ListInvoices(params ListInvoicesParams) (*[]Invoice, error)
    // https://api.domeneshop.no/docs/#operation/getInvoices

func (s *Client) Post(endpoint string, message []byte) ([]byte, error)
    // Performs a POST request

func (s *Client) Put(endpoint string, message []byte) ([]byte, error)
    // Performs a PUT request

func (s *Client) Request(options RequestOptions) ([]byte, error)
    // Base Request func, used by higher level request interfaces

func (s *Client) UpdateDDns(params UpdateDDnsParams) (*interface{}, error)
    // https://api.domeneshop.no/docs/#tag/ddns/paths/~1dyndns~1update/get As the
    // api is not ready, assume that it is successfull if no error is returned

func (s *Client) UpdateDnsRecord(domainId, recordId int, updatedRecord Dns) (*DnsSaveResponse, error)
    // https://api.domeneshop.no/docs/#tag/dns/paths/~1domains~1{domainId}~1dns~1{recordId}/put

func (s *Client) UpdateHttpForward(domainId int, host string, updatedForward HttpForward) (*HttpForward, error)
    // https://api.domeneshop.no/docs/#tag/forwards/paths/~1domains~1{domainId}~1forwards~1{host}/put
```
