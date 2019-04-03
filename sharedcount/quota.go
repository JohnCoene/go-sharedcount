package sharedcount

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Quota list of rate limited-related information
type Quota struct {
	QuotaUsedToday      int    `json:"quota_used_today"`
	Plan                string `json:"plan"`
	QuotaRemainigToday  int    `json:"quota_remaining_today"`
	QuotaAllocatedToday int    `json:"quota_allocated_today"`
}

// GetQuota Return rate limit information
func (apikey *APIKey) GetQuota() Quota {

	data := new(Quota)

	// Call API
	var client = &http.Client{Timeout: 10 * time.Second}
	endpoint := baseurl + "/quota?apikey=" + apikey.Key
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
