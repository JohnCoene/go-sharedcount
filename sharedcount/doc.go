/*
Package sharedcount provides an interface to the SharedCount API.
Install with
	go get github.com/JohnCoene/go-sharedcount
Get a free API key on SharedCount's website then pass API key to various functions;
always as first argument.
	key := &sharedcount.APIKey{Key: "myAPIkey"}
	data := key.GetURL("https://golang.org/")
*/
package sharedcount
