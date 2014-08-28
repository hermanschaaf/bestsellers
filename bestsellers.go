// Package bestsellers provides a simplified interface to the New York
// Times Best Sellers List API.
package bestsellers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// dateFmt describes the date format as specified by the API (YYYY-MM-DD)
const dateFmt = "2006-01-02"
const (
	Daily = iota
	Weekly
	Monthly
	Yearly
	Never
)

// BaseResponse is the basic response returned by all endpoints in the API.
// It includes the status of the request, a copyright notice and the number
// of results returned for the query.
type BaseResponse struct {
	Status     string `json:"status"`
	Copyright  string `json:"copyright"`
	NumResults int    `json:"num_results"`
}

// Client is the main struct used to interface with the API.
// API methods are implemented as methods on this struct, and so
// the first step of any interaction with the API client must be
// to insantiate this struct. This can be done using the NewClient
// function.
type Client struct {
	apiKey  string
	rootURL string
}

// NewClient returns a new Client, which can be used to interface
// with the API. It takes only an APIKey string as single argument.
func NewClient(APIKey string) *Client {
	rootURL := "http://api.nytimes.com/"
	return &Client{APIKey, rootURL}
}

// getURL takes a resource path as a string and returns a
// url.URL customized for the client settings and specified
// path on the API root.
func (c *Client) getURL(path string) (*url.URL, error) {
	u, err := url.Parse(c.rootURL)
	if err != nil {
		return nil, err
	}

	u.Path = path
	params := u.Query()
	params.Add("api-key", c.apiKey)
	u.RawQuery = params.Encode()

	return u, err
}

// get takes a path string and performs a GET request to the specified
// path for this client, and returns the result as a byte slice, or an
// not-nil error if something went wrong during the request.
func (c *Client) get(path string, offset int) ([]byte, error) {
	u, err := c.getURL(path)
	if err != nil {
		return []byte{}, err
	}

	// add offset parameter if set
	if offset > 0 {
		params := u.Query()
		params.Add("offset", strconv.Itoa(offset))
		u.RawQuery = params.Encode()
	}

	response, err := http.Get(u.String())
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	return contents, err
}

// ListNamesResponse describes the response returned by ListNames.
// It includes meta-information returned by the API, such as status
// and copyright information, as well as the number of results returned.
type ListNamesResponse struct {
	BaseResponse
	Results []listNamesResult `json:"results"`
}

// jsonTime allows us to parse dates from the JSON response
type jsonTime time.Time

func (j *jsonTime) UnmarshalJSON(b []byte) error {
	d, err := time.Parse(dateFmt, strings.Trim(string(b), `"`))
	if err != nil {
		return err
	}
	*j = jsonTime(d)
	return nil
}

type updateType uint8

func (u *updateType) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	switch s {
	case "DAILY":
		*u = Daily
	case "WEEKLY":
		*u = Weekly
	case "MONTHLY":
		*u = Monthly
	case "YEARLY":
		*u = Yearly
	default:
		*u = Never
	}
	return nil
}

// listNamesResult describes the form of a single ListName result returned
// by ListNames. DisplayName contains a human-formatted description of the list name,
// and ListNameEncoded is the API-friendly name that should be used when calling
// other API methods, such as Lists and ListsByDate.
type listNamesResult struct {
	ListName            string     `json:"list_name"`
	DisplayName         string     `json:"display_name"`
	ListNameEncoded     string     `json:"list_name_encoded"`
	OldestPublishedDate jsonTime   `json:"oldest_published_date"`
	NewestPublishedDate jsonTime   `json:"newest_published_date"`
	Updated             updateType `json:"updated"`
}

// ListNames returns the response for /svc/books/v2/lists/names
func (c *Client) ListNames() (*ListNamesResponse, error) {
	content, err := c.get("svc/books/v2/lists/names", 0)
	if err != nil {
		return nil, err
	}

	var listNames ListNamesResponse
	err = json.Unmarshal(content, &listNames)
	return &listNames, err
}

type ListsResponse struct {
}

func (c *Client) lists(path string, offset int) (*ListsResponse, error) {
	content, err := c.get(path, offset)
	if err != nil {
		return nil, err
	}

	var lists ListsResponse
	err = json.Unmarshal(content, &lists)
	return &lists, err
}

// Lists returns the reponse for /svc/books/v2/lists/{list-name}.
func (c *Client) Lists(listName string, offset int) (*ListsResponse, error) {
	p := fmt.Sprintf("/svc/books/v2/lists/%s")
	return c.lists(p, offset)
}

// ListsByDate returns the reponse for /svc/books/v2/lists/{date}/{list-name}.
func (c *Client) ListsByDate(listName string, date time.Time, offset int) (*ListsResponse, error) {
	p := fmt.Sprintf("/svc/books/v2/lists/%s/%s", date.Format(dateFmt), listName)
	return c.lists(p, offset)
}
