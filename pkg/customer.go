package quickbooks

import "errors"

type CustomerQueryResponse struct {
	QueryResponse struct {
		Customer      []Customer
		StartPosition int64
		MaxResults    int64
		TotalCount    int64
	}
	Time string
}

type CustomerGetOrCreateResponse struct {
	Customer Customer
	Time     string
}

func (c *CustomerGetOrCreateResponse) yield(err error) (*Customer, error) {
	if err == nil && c != nil {
		return &c.Customer, nil
	}
	if err == nil {
		return nil, errors.New("something weird happened")
	}
	return nil, err
}

func (r *RefreshToken) CreateCustomer(c *CreateCustomer) (*Customer, error) {
	customer := &CustomerGetOrCreateResponse{}
	err := r.DoRequest(
		"POST",
		"/v3/company/{realmId}/customer",
		nil,
		nil,
		map[string]string{
			"body": stringify(c),
		},
		customer,
	)
	return customer.yield(err)
}

func (r *RefreshToken) GetCustomer(id string) (*Customer, error) {
	customer := &CustomerGetOrCreateResponse{}
	err := r.DoRequest(
		"GET",
		"/v3/company/{realmId}/customer/"+id,
		nil,
		nil,
		nil,
		customer,
	)
	return customer.yield(err)
}

func (r *RefreshToken) FullUpdateCustomer(c *Customer) (*Customer, error) {
	c.CurrencyRef = nil
	customer := &CustomerGetOrCreateResponse{}
	err := r.DoRequest("POST",
		"/v3/company/{realmId}/customer",
		nil,
		nil,
		map[string]string{
			"body": stringify(c),
		},
		customer,
	)
	return customer.yield(err)
}
