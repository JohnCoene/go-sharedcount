package sharedcount

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Domains whitelisted domains
type Domains struct {
	Domains                []string `json:"domains"`
	WhitelistDomainEnabled bool     `json:"whitelist_domains_enabled"`
}

// GetDomains returns whitelisted domains
func GetDomains(apikey string) Domains {

	data := new(Domains)

	// Call API
	var client = &http.Client{Timeout: 10 * time.Second}
	endpoint := baseurl + "/domain_whitelist?apikey=" + apikey
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
