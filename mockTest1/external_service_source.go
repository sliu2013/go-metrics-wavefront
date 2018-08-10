package mockTest1

import "net/http"

//type Getter interface {
//	Get(url string) (resp *http.Response, err error)
//}
//
//type ExternalServiceClient struct {
//	HTTPClient Getter
//}

func (c *ExternalServiceClient) Call() (result *Result, err error) {
	// ...
	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient.Get()
	}

	resp, err := httpClient.Get(url)
	// ...

	http.DefaultClient.Get(url)
}