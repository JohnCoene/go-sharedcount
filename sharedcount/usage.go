package sharedcount

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Usage usage data of apikey
type Usage map[string]Day

// Day daily consumption
type Day struct {
	AllowOverages int
	Date          string
	Queries       int
	Quota         int
}

// GetUsage returns usage data of apikey
func GetUsage(apikey string) Usage {

	data := new(Usage)

	// Call API
	var client = &http.Client{Timeout: 10 * time.Second}
	endpoint := baseurl + "/usage?apikey=" + apikey
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
