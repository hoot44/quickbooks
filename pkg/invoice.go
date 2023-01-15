package quickbooks

import (
	"errors"
)

type InvoiceQueryResponse struct {
	QueryResponse struct {
		Invoice       []*Invoice
		StartPosition uint64 `json:"startPosition"`
		MaxResults    uint64 `json:"maxResults"`
		TotalCount    uint64 `json:"totalCount"`
	}
	Time string `json:"time"`
}

type InvoiceGetOrCreateResponse struct {
	Invoice Invoice
	Time    string `json:"time"`
}

func (i *InvoiceGetOrCreateResponse) yield(err error) (*Invoice, error) {
	if err == nil && i != nil {
		return &i.Invoice, nil
	}
	if err == nil {
		return nil, errors.New("something weird happened")
	}
	return nil, err
}

func (r *RefreshToken) CreateInvoice(inv *CreateInvoice) (*Invoice, error) {
	i := &InvoiceGetOrCreateResponse{}
	err := r.DoRequest(
		"POST",
		"/v3/company/{realmId}/invoice",
		nil,
		nil,
		map[string]string{
			"body": stringify(inv),
		},
		i,
	)
	return i.yield(err)
}

func (r *RefreshToken) GetInvoice(id string) (*Invoice, error) {
	invoice := &InvoiceGetOrCreateResponse{}
	err := r.DoRequest(
		"GET",
		"/v3/company/{realmId}/invoice/"+id,
		nil,
		nil,
		nil,
		invoice,
	)
	return invoice.yield(err)
}

func (r *RefreshToken) FullUpdateInvoice(i *Invoice) (*Invoice, error) {
	inv := &InvoiceGetOrCreateResponse{}
	err := r.DoRequest(
		"POST",
		"/v3/company/{realmId}/invoice",
		nil,
		nil,
		map[string]string{
			"body": stringify(i),
		},
		inv,
	)
	return inv.yield(err)
}
