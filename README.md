New York Times Best Seller List Go API Client
---------------------------------------------

[![Build Status](https://travis-ci.org/hermanschaaf/bestsellers.svg?branch=master)](https://travis-ci.org/hermanschaaf/bestsellers)

A simple Go API Client for the New York Times Best Seller List. This is still under development, so use at your own discretion. If you would like to contribute to this library, you are very welcome. When doing so, please adhere to these guidelines:

 - Run `goimports` on files before making a PR ([goimports](https://godoc.org/code.google.com/p/go.tools/cmd/goimports))
 - All new code must be accompanied by a test
 - If it's a very significant departure from what is already here, first open a Github issue for discussion and to get a go-ahead

For more information, read the docs below:

 - [Go Client Docs](http://godoc.org/github.com/hermanschaaf/bestsellers)
 - [API Docs](http://developer.nytimes.com/docs/best_sellers_api)

### Example usage

```
package main

import (
	"fmt"

	"github.com/hermanschaaf/bestsellers"
)

func main() {
	c := bestsellers.NewClient("your-api-key")

	// get the best seller lists
	lists, err := c.ListNames()
	if err != nil {
		panic(err)
	}

	// print the first three list names
	for _, r := range lists.Results[:3] {
		fmt.Printf("%q\n", r)
	}
}
```

This prints something like the following:

```
{"Combined Print and E-Book Fiction" "Combined Print & E-Book Fiction" "combined-print-and-e-book-fiction" "2011-02-13" "2014-08-31" "WEEKLY"}
{"Combined Print and E-Book Nonfiction" "Combined Print & E-Book Nonfiction" "combined-print-and-e-book-nonfiction" "2011-02-13" "2014-08-31" "WEEKLY"}
{"Hardcover Fiction" "Hardcover Fiction" "hardcover-fiction" "2008-06-08" "2014-08-31" "WEEKLY"}
```