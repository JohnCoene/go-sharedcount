package sharedcount

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const BASEURL = "https://api.sharedcount.com/v1.0"

// Shared sharedcount object
type Shared struct {
	StumbleUpon   int      `json:"StumbleUpon"`
	Pinterest     int      `json:"Pinterest"`
	LinkedIn      int      `json:"LinkedIn"`
	Facebook      Facebook `json:"Facebook"`
	GooglePlusOne int      `json:"GooglePlusOne"`
}

// Facebook sharedcount data for Facebook
type Facebook struct {
	TotalCount         int `json:"total_count"`
	CommentCount       int `json:"comment_count"`
	ReactionCount      int `json:"reaction_count"`
	ShareCount         int `json:"share_count"`
	CommentPluginCount int `json:"comment_plugin_count"`
	OgObject           int `json:"og_object"`
}

// GetUrl returns sharedcount data for specific url
func GetUrl(apikey, url string) Shared {
	// initialise data
	data := new(Shared)

	// Call API
	var client = &http.Client{Timeout: 10 * time.Second}
	endpoint := BASEURL + "?apikey=" + apikey + "&url=" + url
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
