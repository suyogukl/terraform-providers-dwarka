package dwarka

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Dwarka URL
const HostURL string = "http://localhost:1410"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
}

// NewClient -
func NewClient(host *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Dwarka URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, expectedStatusCodes ...int) ([]byte, error) {
	expectedStatusCodes = append(expectedStatusCodes, http.StatusOK)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	statusCode := res.StatusCode

	for _, expectedStatusCode := range expectedStatusCodes {
		if statusCode == expectedStatusCode {
			return body, err
		}
	}
	return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
}
