// PasswordState Client

package passwordstateclient

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// ApiUrl - Placeholder URL
const ApiUrl string = "http://passwordstate"

// Client -
type Client struct {
	ApiUrl    	string
	HTTPClient 	*http.Client
	ApiKey		string
}

// Function - Create Client
func NewClient (api_url, api_key *string) (*Client, error) {

	// Set Client Config
	psClient := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Placeholder URL
		ApiUrl: ApiUrl,
	}

	// If the api url is provided, set the url to the provided one
	if api_url != nil {
		psClient.ApiUrl = *api_url
	}

	// If the api key is provided, set the key to the provided one
	if api_key != nil {
		psClient.ApiKey = *api_key
	}

	// If api key is not provided, return empty client
	if api_key == nil {
		return &psClient, nil
	}

	// Return Client
	return &psClient, nil
		
}

// Function - Do/send the API Request
func (psClient *Client) doRequest(req *http.Request, api_key *string) ([]byte, error) {
	// Set the API key
	key := psClient.ApiKey

	// Check if the API key is null if not set key
	if api_key != nil {
		key = *api_key
	}

	// Add Request Headers
	req.Header.Add("APIKey", key)

	// Perform the API HTTP Request
	res, err := psClient.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Read and Check the Response Body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Check the Status Code and return error if not OK
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s, url: %q", res.StatusCode, body, res.Request.URL)
	}

	// Return the Body
	return body, err
}
