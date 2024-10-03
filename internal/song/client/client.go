package client

import (
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"net/http"
	"net/url"
	"time"
)

// ApiClient - структура для клиента Heimdall.
type ApiClient struct {
	Client  *httpclient.Client
	BaseUrl string
}

// NewClient - конструктор *ApiClient с таймаутом и повторными попытками.
func NewClient(url string) *ApiClient {
	backoff := heimdall.NewConstantBackoff(2*time.Second, 5*time.Second)

	retrier := heimdall.NewRetrier(backoff)

	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(5*time.Second),
		httpclient.WithRetryCount(3),
		httpclient.WithRetrier(retrier),
	)

	return &ApiClient{
		Client:  client,
		BaseUrl: url,
	}
}

// Get - метод для выполнения GET запроса с необязательными URL параметрами.
func (c *ApiClient) Get(path string, queryParams ...map[string]string) (*http.Response, error) {
	baseURL, err := url.Parse(c.BaseUrl + path)
	if err != nil {
		return nil, err
	}

	if len(queryParams) > 0 {
		params := url.Values{}
		for key, value := range queryParams[0] {
			params.Add(key, value)
		}
		baseURL.RawQuery = params.Encode()
	}

	response, err := c.Client.Get(baseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}
