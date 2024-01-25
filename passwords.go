// Functions for api/passwords

package passwordstateclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Function - Get Password Data from the API
func (psClient *Client) GetPassword(password_id int64) ([]Password, error) {

	// Convert Int to String
	id := strconv.Itoa(int(password_id))

	// Check if API URL is empty
	if psClient.ApiUrl == ""{
		return nil, fmt.Errorf("define API URL")
	}

	// Check if API Key is empty
	if psClient.ApiKey == "" {
		return nil, fmt.Errorf("define API Key")
	}

	// Check if Password ID is empty
	if id == "" {
		return nil, fmt.Errorf("define Password ID (PID)")
	}

	// Join URL and Password ID into one string
	finalApiUrl := psClient.ApiUrl + "/" + id

	// Configure the API HTTP Request
	req, err := http.NewRequest("GET", finalApiUrl, nil)
	if err != nil {
		return nil, err
	}

	// Perform the API HTTP Request
	body, err := psClient.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	// Pharse Response Body, configure based on model, and return data
	passwords := []Password{}
	err = json.Unmarshal(body, &passwords)
	if err != nil {
		return nil, err
	}

	// Return Final Data
	return passwords, nil
}
