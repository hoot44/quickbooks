package quickbooks

type ClientQueryResponse struct {
	QueryResponse struct {
		Customer []Client
	}
}

func (c *RefreshToken) QueryCustomer(query string) ([]Client, error) {
	request, err := c.Request("GET", "/v3/company/{realmID}/query", map[string]string{
		"query": query,
	}, nil, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.api.client.Do(request)
	if err != nil {
		return nil, err
	}

	var responseToken = &ClientQueryResponse{}
	err = deserialize[*ClientQueryResponse](response, responseToken)
	if err != nil {
		return nil, err
	}

	for _, cc := range responseToken.QueryResponse.Customer {
		cc.rt = c
	}

	return responseToken.QueryResponse.Customer, nil
}
