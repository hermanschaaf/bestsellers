package bestsellers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestListNames(t *testing.T) {
	var dummyListNamesResponse = `{
		"status":"OK",
		"copyright":"copyright",
		"num_results":30,
		"results": [{
			"list_name":"Combined Print and E-Book Fiction",
			"display_name":"Combined Print & E-Book Fiction",
			"list_name_encoded":"combined-print-and-e-book-fiction",
			"oldest_published_date":"2011-02-13",
			"newest_published_date":"2014-08-31",
			"updated":"WEEKLY"
		}
	]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(dummyListNamesResponse))

		wantURL := "/svc/books/v2/lists/names?api-key=test-api-key"
		if r.URL.String() != wantURL {
			t.Errorf("Request URL = %q, want %q", r.URL, wantURL)
		}
	}))
	defer ts.Close()

	c := NewClient("test-api-key")
	c.rootURL = ts.URL

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

// func TestLists(t *testing.T) {
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte(dummyListNamesResponse))

// 		wantURL := "/svc/books/v2/lists/names?api-key=test-api-key"
// 		if r.URL.String() != wantURL {
// 			t.Errorf("Request URL = %q, want %q", r.URL, wantURL)
// 		}
// 	}))
// 	defer ts.Close()

// 	c := NewClient("test-api-key")
// 	c.rootURL = ts.URL

// 	_, err := c.ListNames()
// 	if err != nil {
// 		t.Fatalf("Error: %v", err)
// 	}
// }
