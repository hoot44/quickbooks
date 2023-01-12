package quickbooks

type RefreshToken struct {
	TokenType              string `json:"token_type"`
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	ExpiresIn              int64  `json:"expires_in"`
	XRefreshTokenExpiresIn int64  `json:"x_refresh_token_expires_in"`
	api                    *api   `json:"-"`
}

func (c *RefreshToken) Query(query string, ifc interface{}) error {
	return c.DoRequest(
		"GET",
		"/v3/company/{realmId}/query",
		map[string]string{
			"query": query,
		},
		nil,
		nil,
		ifc,
	)
}
