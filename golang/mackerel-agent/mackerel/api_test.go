package mackerel

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mackerelio/mackerel-agent/version"
)

func TestNewAPI(t *testing.T) {
	api, err := NewAPI(
		"http://example.com",
		"dummy-key",
		true,
	)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}

	fmt.Println(api.BaseURL.String())
	if api.BaseURL.String() != "http://example.com" {
		t.Error("should return URL")
	}

	fmt.Println(api.APIKey)
	if api.APIKey != "dummy-key" {
		t.Error("should return api key")
	}

	fmt.Println(api.Verbose)
	if api.Verbose != true {
		t.Error("should return verbose value")
	}
}

func TestUrlFor(t *testing.T) {
	api, _ := NewAPI(
		"http://example.com",
		"dummy-key",
		true,
	)

	fmt.Println(api.urlFor("/").String())
	if api.urlFor("/").String() != "http://example.com/" {
		t.Error("should return http://example/")
	}

	fmt.Println(api.urlFor("/path/to/api").String())
	if api.urlFor("/path/to/api").String() != "http://example.com/path/to/api" {
		t.Error("should return http://example.com/path/to/api")
	}

	// query string don't work since using url.Path
	fmt.Println(api.urlFor("/path/to/api/?key1=value&key2=value").String())
}

func TestDo(t *testing.T) {
	version.VERSION = "1.0.0"
	version.GITCOMMIT = "1234beaf"
	handler := func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Header)

		userAgent := "mackerel-agent/1.0.0 (Revision 1234beaf)"
		if req.Header.Get("X-Api-Key") != "dummy-key" {
			t.Error("X-Api-Key header should contains passed key")
		}

		if h := req.Header.Get("User-Agent"); h != userAgent {
			t.Errorf("User-Agent should be '%s' but %s", userAgent, h)
		}
	}

	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		handler(res, req)
	}))
	defer ts.Close()

	api, _ := NewAPI(
		ts.URL,
		"dummy-key",
		false,
	)

	req, _ := http.NewRequest("GET", api.urlFor("/").String(), nil)
	api.do(req)
}
