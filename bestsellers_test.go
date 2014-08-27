package bestsellers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListNames(t *testing.T) {
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

	_, err := c.ListNames()
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
}

var dummyListNamesResponse = `{"status":"OK","copyright":"Copyright (c) 2014 The New York Times Company.  All Rights Reserved.","num_results":30,"results":[{"list_name":"Combined Print and E-Book Fiction","display_name":"Combined Print & E-Book Fiction","list_name_encoded":"combined-print-and-e-book-fiction","oldest_published_date":"2011-02-13","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Combined Print and E-Book Nonfiction","display_name":"Combined Print & E-Book Nonfiction","list_name_encoded":"combined-print-and-e-book-nonfiction","oldest_published_date":"2011-02-13","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Hardcover Fiction","display_name":"Hardcover Fiction","list_name_encoded":"hardcover-fiction","oldest_published_date":"2008-06-08","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Hardcover Nonfiction","display_name":"Hardcover Nonfiction","list_name_encoded":"hardcover-nonfiction","oldest_published_date":"2008-06-08","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Trade Fiction Paperback","display_name":"Paperback Trade Fiction","list_name_encoded":"trade-fiction-paperback","oldest_published_date":"2008-06-08","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Mass Market Paperback","display_name":"Paperback Mass-Market Fiction","list_name_encoded":"mass-market-paperback","oldest_published_date":"2008-06-08","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Paperback Nonfiction","display_name":"Paperback Nonfiction","list_name_encoded":"paperback-nonfiction","oldest_published_date":"2008-06-08","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"E-Book Fiction","display_name":"E-Book Fiction","list_name_encoded":"e-book-fiction","oldest_published_date":"2011-02-13","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"E-Book Nonfiction","display_name":"E-Book Nonfiction","list_name_encoded":"e-book-nonfiction","oldest_published_date":"2011-02-13","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Hardcover Advice","display_name":"Hardcover Advice & Misc.","list_name_encoded":"hardcover-advice","oldest_published_date":"2008-06-08","newest_published_date":"2013-04-21","updated":"WEEKLY"},{"list_name":"Paperback Advice","display_name":"Paperback Advice & Misc.","list_name_encoded":"paperback-advice","oldest_published_date":"2008-06-08","newest_published_date":"2013-04-21","updated":"WEEKLY"},{"list_name":"Advice How-To and Miscellaneous","display_name":"Advice, How-To & Miscellaneous","list_name_encoded":"advice-how-to-and-miscellaneous","oldest_published_date":"2013-04-28","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Picture Books","display_name":"Children's Picture Books","list_name_encoded":"picture-books","oldest_published_date":"2008-06-08","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Chapter Books","display_name":"Children's Chapter Books","list_name_encoded":"chapter-books","oldest_published_date":"2008-06-08","newest_published_date":"2012-12-09","updated":"WEEKLY"},{"list_name":"Childrens Middle Grade","display_name":"Children's Middle Grade","list_name_encoded":"childrens-middle-grade","oldest_published_date":"2012-12-16","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Young Adult","display_name":"Young Adult","list_name_encoded":"young-adult","oldest_published_date":"2012-12-16","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Paperback Books","display_name":"Children's Paperback Books","list_name_encoded":"paperback-books","oldest_published_date":"2008-06-08","newest_published_date":"2012-12-09","updated":"WEEKLY"},{"list_name":"Series Books","display_name":"Children's Series","list_name_encoded":"series-books","oldest_published_date":"2008-06-08","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Hardcover Graphic Books","display_name":"Hardcover Graphic Books","list_name_encoded":"hardcover-graphic-books","oldest_published_date":"2009-03-15","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Paperback Graphic Books","display_name":"Paperback Graphic Books","list_name_encoded":"paperback-graphic-books","oldest_published_date":"2009-03-15","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Manga","display_name":"Manga","list_name_encoded":"manga","oldest_published_date":"2009-03-15","newest_published_date":"2014-08-31","updated":"WEEKLY"},{"list_name":"Combined Print Fiction","display_name":"Combined Hardcover & Paperback Fiction","list_name_encoded":"combined-print-fiction","oldest_published_date":"2011-02-13","newest_published_date":"2013-05-12","updated":"WEEKLY"},{"list_name":"Combined Print Nonfiction","display_name":"Combined Hardcover & Paperback Nonfiction","list_name_encoded":"combined-print-nonfiction","oldest_published_date":"2011-02-13","newest_published_date":"2013-05-12","updated":"WEEKLY"},{"list_name":"Hardcover Business Books","display_name":"Hardcover Business Books","list_name_encoded":"hardcover-business-books","oldest_published_date":"2011-07-03","newest_published_date":"2013-10-13","updated":"MONTHLY"},{"list_name":"Paperback Business Books","display_name":"Paperback Business Books","list_name_encoded":"paperback-business-books","oldest_published_date":"2011-07-03","newest_published_date":"2013-10-13","updated":"MONTHLY"},{"list_name":"Business Books","display_name":"Business Books","list_name_encoded":"business-books","oldest_published_date":"2013-11-03","newest_published_date":"2014-08-10","updated":"MONTHLY"},{"list_name":"Hardcover Political Books","display_name":"Political Books","list_name_encoded":"hardcover-political-books","oldest_published_date":"2011-07-03","newest_published_date":"2014-08-10","updated":"MONTHLY"},{"list_name":"Science","display_name":"Science","list_name_encoded":"science","oldest_published_date":"2013-04-14","newest_published_date":"2014-08-10","updated":"MONTHLY"},{"list_name":"Food and Fitness","display_name":"Food and Fitness","list_name_encoded":"food-and-fitness","oldest_published_date":"2013-09-01","newest_published_date":"2014-08-10","updated":"MONTHLY"},{"list_name":"Sports","display_name":"Sports","list_name_encoded":"sports","oldest_published_date":"2014-03-02","newest_published_date":"2014-08-10","updated":"MONTHLY"}]}`