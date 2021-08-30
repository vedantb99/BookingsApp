package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{

	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-sa", "/search-availability", "POST", []postData{
		{key: "start", value: "2020:01:02"},
		{key: "end", value: "2021:01:05"},
	}, http.StatusOK},
	{"post-sa-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020:01:02"},
		{key: "end", value: "2021:01:05"},
	}, http.StatusOK},
	{"make-resrvation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Vedant"},
		{key: "last_name", value: "B"},
		{key: "email", value: "vedant@mkcl.org"},
		{key: "phone", value: "2021123123"},
	}, http.StatusOK},
}

func TestHtandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()
	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, we expected %d but got back %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, we expected %d but got back %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
