[![](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/JohnCoene/go-sharedcount/sharedcount) [![Go Report Card](https://goreportcard.com/badge/github.com/JohnCoene/go-sharedcount)](https://goreportcard.com/report/github.com/JohnCoene/go-sharedcount) 

# sharedcount

Go interface to the [SharedCount](https://www.sharedcount.com) API.

## Docs

Documentation is on [Godoc](https://godoc.org/github.com/JohnCoene/go-sharedcount/sharedcount).

## Examples

_API key available for free on the [website](https://www.sharedcount.com)._ Set yourself up with a key then call methods to get data.

```go
package main

import (
	"fmt"
	"go-sharedcount/sharedcount"
)

func main() {

	key := &sharedcount.APIKey{Key: "myKeyGoesHere"} // setup

  // Get API usage
	quota := key.GetURL("https://golang.org")
	fmt.Println(quota)
}
```