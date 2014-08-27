New York Times Best Seller List Go API Client
=============================================

[Go Client Docs](http://godoc.org/github.com/hermanschaaf/bestsellers)
[API Docs](http://developer.nytimes.com/docs/best_sellers_api)

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