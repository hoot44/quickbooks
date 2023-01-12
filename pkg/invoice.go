package quickbooks

import (
	"errors"
)

type InvoiceQueryResponse struct {
	QueryResponse struct {
		Invoice       []Invoice
		StartPosition int64
		MaxResults    int64
		TotalCount    int64
	}
	Time string
}

type InvoiceGetOrCreateResponse struct {
	Invoice Invoice
	Time    string
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
