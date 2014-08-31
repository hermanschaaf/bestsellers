package bestsellers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func setupTestServer(t *testing.T, wantURL string, dummyResponse []byte) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(dummyResponse)

		if r.URL.String() != wantURL {
			t.Errorf("Request URL = %q, want %q", r.URL, wantURL)
		}
	}))
	return ts
}

func TestListNames(t *testing.T) {
	dummyListNamesResponse, err := ioutil.ReadFile("testdata/listnames.json")
	if err != nil {
		t.Fatal("Error reading json testdata:", err)
	}

	ts := setupTestServer(t, "/svc/books/v2/lists/names?api-key=test-api-key", dummyListNamesResponse)
	defer ts.Close()

	// create a new API client
	c := NewClient("test-api-key")
	c.rootURL = ts.URL

	// get the available list names
	got, err := c.ListNames()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	oldestDate, _ := time.Parse(dateFmt, "2011-02-13")
	newestDate, _ := time.Parse(dateFmt, "2014-08-31")
	want := ListNamesResponse{
		BaseResponse: BaseResponse{
			Status:     "OK",
			Copyright:  "copyright",
			NumResults: 30,
		},
		Results: []ListNamesResult{
			ListNamesResult{
				ListName:            "Combined Print and E-Book Fiction",
				DisplayName:         "Combined Print & E-Book Fiction",
				ListNameEncoded:     "combined-print-and-e-book-fiction",
				OldestPublishedDate: Time(oldestDate),
				NewestPublishedDate: Time(newestDate),
				Updated:             UpdateType(Weekly),
			},
		},
	}

	if !reflect.DeepEqual(got.BaseResponse, want.BaseResponse) {
		t.Errorf("got BaseResponse = %q, want %q", got.BaseResponse, want.BaseResponse)
	}

	if len(got.Results) != len(want.Results) {
		t.Fatalf("got len(Results) = %d, want %d", len(got.Results), len(want.Results))
	}

	if !reflect.DeepEqual(got.Results[0], want.Results[0]) {
		t.Errorf("got Results[0] = %q, want %q", got.Results[0], want.Results[0])
	}
}

func TestLists(t *testing.T) {
	dummyListResponse, err := ioutil.ReadFile("testdata/lists.json")
	if err != nil {
		t.Fatal("Error reading json testdata:", err)
	}

	ts := setupTestServer(t, "/svc/books/v2/lists/hardcover-nonfiction?api-key=test-api-key", dummyListResponse)
	defer ts.Close()

	// create a new API client
	c := NewClient("test-api-key")
	c.rootURL = ts.URL

	// get the hardcover-fiction list, with 0 offset
	got, err := c.Lists("hardcover-nonfiction", 0)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	// verify that the response was correct and the JSON correctly marshaled
	wantBaseResponse := BaseResponse{
		Status:     "OK",
		Copyright:  "Copyright (c) 2014 The New York Times Company.  All Rights Reserved.",
		NumResults: 25,
	}

	if got.BaseResponse != wantBaseResponse {
		t.Errorf("Got BaseResponse = %v, want %v", got.BaseResponse, wantBaseResponse)
	}

	if len(got.Results) != 2 {
		t.Fatalf("Got len(Results) = %d, want %d", len(got.Results), 2)
	}

	wantISBNs := []ISBN{
		ISBN{ISBN10: "1595231129", ISBN13: "9781595231123"},
		ISBN{ISBN10: "1611763398", ISBN13: "9781611763393"},
	}

	if !reflect.DeepEqual(got.Results[0].ISBNs, wantISBNs) {
		t.Error("got ISBNS = %v, want %v", got.Results[0].ISBNs, wantISBNs)
	}

	wantDetails := []BookDetails{
		BookDetails{
			Title:            "ONE NATION",
			Description:      "Carson, a retired pediatric neurosurgeon, now a Fox News contributor, offers solutions to problems.",
			Contributor:      "by Ben Carson with Candy Carson",
			Author:           "Ben Carson with Candy Carson",
			ContributorNote:  "",
			Price:            0,
			AgeGroup:         "",
			Publisher:        "Sentinel",
			PrimaryISBN13:    "9781595231123",
			PrimaryISBN10:    "1595231129",
			BookImage:        "http://du.ec2.nytimes.com.s3.amazonaws.com/prd/books/9781595231123.jpg",
			AmazonProductURL: "http://www.amazon.com/One-Nation-What-Americas-Future/dp/1595231129?tag=thenewyorktim-20",
		},
	}

	if !reflect.DeepEqual(got.Results[0].BookDetails, wantDetails) {
		t.Errorf("got BookDetails = %v, want %v", got.Results[0].BookDetails, wantDetails)
	}

	if got.Results[0].DisplayName != "Hardcover Nonfiction" {
		t.Errorf("got Results[0].DisplayName = %q, want %q", got.Results[0].DisplayName, "Hardcover Nonfiction")
	}

	if got.Results[0].Updated != Weekly {
		t.Errorf("got Results[0].Updated = %v, want %v", got.Results[0].Updated, "<weekly>")
	}

	if got.Results[0].Asterisk != Bool(false) {
		t.Errorf("got Results[0].Asterisk = %v, want %v", got.Results[0].Asterisk, Bool(false))
	}

	if got.Results[0].Dagger != Bool(false) {
		t.Errorf("got Results[0].Dagger = %v, want %v", got.Results[0].Dagger, Bool(false))
	}
}

func TestListsOffset(t *testing.T) {
	dummyListResponse, err := ioutil.ReadFile("testdata/lists.json")
	if err != nil {
		t.Fatal("Error reading json testdata:", err)
	}

	ts := setupTestServer(t, "/svc/books/v2/lists/ebook-fiction?api-key=test-api-key&offset=10", dummyListResponse)
	defer ts.Close()

	// create a new API client
	c := NewClient("test-api-key")
	c.rootURL = ts.URL

	// get the ebook-fiction list, with 10 offset
	_, err = c.Lists("ebook-fiction", 10)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

func TestListsByDate(t *testing.T) {
	dummyListResponse, err := ioutil.ReadFile("testdata/lists.json")
	if err != nil {
		t.Fatal("Error reading json testdata:", err)
	}

	ts := setupTestServer(t, "/svc/books/v2/lists/2011-02-13/ebook-fiction?api-key=test-api-key&offset=10", dummyListResponse)
	defer ts.Close()

	// create a new API client
	c := NewClient("test-api-key")
	c.rootURL = ts.URL

	// get the ebook-fiction list, with 10 offset
	date, _ := time.Parse(dateFmt, "2011-02-13")
	_, err = c.ListsByDate("ebook-fiction", date, 10)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}
