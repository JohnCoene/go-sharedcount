package sharedcount

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Status API status
type Status struct {
	Status string `json:"status"`
}

// GetStatus returns API status
func (apikey *APIKey) GetStatus() Status {

	data := new(Status)

	// Call API
	var client = &http.Client{Timeout: 10 * time.Second}
	endpoint := baseurl + "/status?apikey=" + apikey.Key
	resp, err := client.Get(endpoint)
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close() // close response

	// Unmarshall
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Panic(err)
	}

	return (*data)
}
