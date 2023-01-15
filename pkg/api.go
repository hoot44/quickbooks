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

type api struct {
	environment                     ENV
	client                          *http.Client
	discovery                       *Discovery
	clientId, clientSecret, realmId string
}

var discoveries map[ENV]*Discovery = map[ENV]*Discovery{}

func NewClient(clientId, clientSecret, realmId string, environment ENV) (*api, error) {
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
		err = deserialize(response, &d)
		if err != nil {
			return nil, err
		}
		discoveries[environment] = &d
		disco = &d
	}
	return &api{
		environment:  environment,
		clientId:     clientId,
		clientSecret: clientSecret,
		discovery:    disco,
		client:       &http.Client{},
		realmId:      realmId,
	}, nil
}

func (c *api) Refresh(token *RefreshToken) (*RefreshToken, error) {
	authStr := base64.StdEncoding.EncodeToString([]byte(c.clientId + ":" + c.clientSecret))
	request, err := Request(
		"POST",
		c.discovery.TokenEndpoint,
		nil,
		map[string]string{
			"Authorization": "Basic " + authStr,
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
	err = deserialize(response, responseToken)
	if err != nil {
		return nil, err
	}

	responseToken.api = c

	return responseToken, nil
}

func deserialize(response *http.Response, ifc interface{}) (e error) {
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
	fmt.Printf("body = %s\n\n", body)
	if response.StatusCode != http.StatusOK {
		qbe := &QuickbooksError{}
		if err = json.Unmarshal(body, qbe); err != nil {
			return err
		}
		if reqid, ok := response.Header["Intuit_tid"]; ok && len(reqid) > 0 {
			fmt.Printf("XXXX\n\n\n")
			qbe.IntuitTid = reqid[0]
		}
		return qbe
	}

	err = json.Unmarshal(body, ifc)
	fmt.Printf("err = %+v\n", err)
	if err != nil {
		qbe := &QuickbooksError{}
		err = json.Unmarshal(body, qbe)
		fmt.Printf("err2 = %+v,qbe=%+v\n", err, qbe)
		if err == nil && qbe.Error() != "" {
			return qbe
		}
		return err
	}

	return nil
}

func stringify(ifc interface{}) (s string) {
	bs, err := json.Marshal(ifc)
	if err != nil {
		panic(err)
	}
	fmt.Printf("stringify %v\n%s\n", ifc, string(bs))
	return string(bs)
}

func (r *RefreshToken) DoRequest(method, uri string, qs, headers, data map[string]string, ifc interface{}) error {
	request, err := r.Request(method, uri, qs, headers, data)
	if err != nil {
		return err
	}

	response, err := r.api.client.Do(request)
	if err != nil {
		return err
	}

	err = deserialize(response, ifc)
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
	u.Path = strings.Replace(uri, "{realmId}", r.api.realmId, -1)

	if headers == nil {
		headers = map[string]string{}
	}
	headers["Authorization"] = "Bearer " + r.AccessToken
	return Request(method, u.String(), qs, headers, data)
}

func Request(method, uri string, qs, headers, data map[string]string) (*http.Request, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	renderedQS := ""
	if qs != nil && len(qs) > 0 {
		qd := url.Values{}
		for k, v := range qs {
			qd.Add(k, v)
		}
		renderedQS = qd.Encode()
		u.RawQuery = renderedQS
	}

	fmt.Printf("%s %s?%s HTTP/1.0\n", method, uri, renderedQS)
	var requestBody string
	if data != nil {
		if bod, ok := data["body"]; len(data) == 1 && ok {
			if _, ok := headers["Content-Type"]; !ok {
				fmt.Println("Content-Type: application/json")
				headers["Content-Type"] = "application/json"
			}
			requestBody = bod
		} else {
			requestData := url.Values{}
			if data != nil {
				for k, v := range data {
					requestData.Set(k, v)
				}
			}
			requestBody = requestData.Encode()
		}
	}

	request, err := http.NewRequest(method, u.String(), bytes.NewBufferString(requestBody))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	fmt.Println("Accept: application/json")
	if headers != nil {
		for k, v := range headers {
			fmt.Printf("%s: %s\n", k, v)
			request.Header.Add(k, v)
		}
	}
	fmt.Printf("%s\n\n", requestBody)

	return request, err
}
