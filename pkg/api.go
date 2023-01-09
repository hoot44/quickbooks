package quickbooks

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type RefreshToken struct {
	TokenType              string `json:"token_type"`
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	ExpiresIn              int64  `json:"expires_in"`
	XRefreshTokenExpiresIn int64  `json:"x_refresh_token_expires_in"`
	realmID                string `json:"-"`
	api                    *api   `json:"-"`
}

type api struct {
	environment                        ENV
	client                             *http.Client
	discovery                          *Discovery
	clientID, clientSecret, clientAuth string
}

var discoveries map[ENV]*Discovery = map[ENV]*Discovery{}

func NewClient(clientID, clientSecret string, environment ENV) (*api, error) {
	authStr := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	var disco *Discovery
	if d, ok := discoveries[environment]; ok {
		disco = d
	} else {
		request, err := Request("GET",
			string(environment),
			nil,
			nil,
			nil,
		)
		if err != nil {
			return nil, err
		}
		response, err := (&http.Client{}).Do(request)
		if err != nil {
			return nil, err
		}

		var d Discovery
		err = deserialize[*Discovery](response, &d)
		if err != nil {
			return nil, err
		}
		discoveries[environment] = &d
		disco = &d
	}
	return &api{
		environment:  environment,
		clientID:     clientID,
		clientSecret: clientSecret,
		clientAuth:   authStr,
		discovery:    disco,
		client:       &http.Client{},
	}, nil
}

func (c *api) Refresh(realm string, token *RefreshToken) (*RefreshToken, error) {
	request, err := Request(
		"POST",
		c.discovery.TokenEndpoint,
		nil,
		map[string]string{
			"Authorization": "Basic " + c.clientAuth,
			"Content-Type":  "application/x-www-form-urlencoded",
		},
		map[string]string{
			"grant_type":    "refresh_token",
			"refresh_token": token.RefreshToken,
		},
	)

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	var responseToken = &RefreshToken{}
	err = deserialize[*RefreshToken](response, responseToken)
	if err != nil {
		return nil, err
	}

	responseToken.api = c
	responseToken.realmID = realm

	return responseToken, nil
}

func deserialize[T any](response *http.Response, ifc T) (e error) {
	defer response.Body.Close()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered %v\n", r)
			e = errors.New(fmt.Sprintf("%+v", r))
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("body =\n%s\n\n", body)

	if response.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	err = json.Unmarshal(body, ifc)
	if err != nil {
		return err
	}

	return nil
}

func (r *RefreshToken) Request(method, uri string, qs, headers, data map[string]string) (*http.Request, error) {
	u, err := url.Parse("https://sandbox-quickbooks.api.intuit.com")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(uri, "{realmID}", r.realmID, -1)

	if headers == nil {
		headers = map[string]string{}
	}
	headers["Authorization"] = "Bearer " + r.AccessToken
	if method != "GET" {
		headers["Content-Type"] = "application/json"
	}
	return Request(method, u.String(), qs, headers, data)
}

func Request(method, uri string, qs, headers, data map[string]string) (*http.Request, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if qs != nil && len(qs) > 0 {
		qd := url.Values{}
		for k, v := range qs {
			qd.Add(k, v)
		}
		u.RawQuery = qd.Encode()
	}

	requestData := url.Values{}
	if data != nil {
		for k, v := range data {
			requestData.Set(k, v)
		}
	}

	request, err := http.NewRequest(method, u.String(), bytes.NewBufferString(requestData.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}

	return request, err
}
