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

func (c *CustomerQueryResponse) yield(err error) ([]Customer, error) {
	if err == nil && c != nil {
		return c.Customer, nil
	}
	if err == nil {
		return nil, errors.New("something weird happened")
	}
	return nil, err
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
		stringify(c),
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
		"",
		customer,
	)
	return customer.yield(err)
}

func (r *RefreshToken) FetchCustomers(id string) ([]Customer, error) {
	cqr := &CustomerQueryResponse{}
	err := r.DoRequest("POST",
		"/v3/company/"+id+"/query?query=SELECT * FROM CUSTOMER",
		nil,
		nil,
		nil,
		cqr,
	)
	return cqr.yield(err)
}

func (r *RefreshToken) FullUpdateCustomer(c *Customer) (*Customer, error) {
	c.CurrencyRef = nil
	customer := &CustomerGetOrCreateResponse{}
	err := r.DoRequest("POST",
		"/v3/company/{realmId}/customer",
		nil,
		nil,
		stringify(c),
		customer,
	)
	return customer.yield(err)
}
