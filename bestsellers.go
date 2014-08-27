package bestsellers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey  string
	rootURL string
}

func NewClient(APIKey string) *Client {
	rootURL := "http://api.nytimes.com/"
	return &Client{APIKey, rootURL}
}

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

type ListNamesResponse struct {
	Status     string            `json:"status"`
	Copyright  string            `json:"copyright"`
	NumResults int               `json:"num_results"`
	Results    []listNamesResult `json:"results"`
}

type listNamesResult struct {
	ListName            string `json:"list_name"`
	DisplayName         string `json:"display_name"`
	ListNameEncoded     string `json:"list_name_encoded"`
	OldestPublishedDate string `json:"oldest_published_date"`
	NewestPublishedDate string `json:"newest_published_date"`
	Updated             string `json:"updated"`
}

// ListNames returns the response for /svc/books/v2/lists/names
func (c *Client) ListNames() (*ListNamesResponse, error) {
	u, err := c.getURL("svc/books/v2/lists/names")
	if err != nil {
		return nil, err
	}
	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	var listNames ListNamesResponse
	err = json.Unmarshal(contents, &listNames)
	return &listNames, err
}
