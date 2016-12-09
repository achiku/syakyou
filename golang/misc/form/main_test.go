package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestForm(t *testing.T) {
	ts := httptest.NewServer(newMux(t))
	defer ts.Close()

	post1 := `order%5Bamount%5D=10000.0&order%5Bbutton_name%5D=%E3%83%9C%E3%82%BF%E3%83%B3+%231&order%5Bconfirmed_at%5D=&order%5Bcreated_at%5D=2016-09-30+02%3A36%3A26+%2B0000&order%5Bcurrency%5D=JPY&order%5Bcustom%5D=525&order%5Bcustomer%5D%5Baddress%5D=&order%5Bcustomer%5D%5Bemail%5D=&order%5Bcustomer%5D%5Bname%5D=&order%5Bevent%5D%5Btype%5D=received&order%5Bid%5D=C-XuxcMDUg&order%5Bstatus%5D=received&order%5Btotal_btc%5D%5Bamount%5D=0.1&order%5Btotal_btc%5D%5Bcurrency%5D=BTC&order%5Btotal_native%5D%5Bamount%5D=10000.0&order%5Btotal_native%5D%5Bcurrency%5D=JPY`
	post2 := `orderAmount=10000&orderButtonName=testtest`
	url := ts.URL + "/post"

	cases := []struct {
		body string
	}{
		{body: post1},
		{body: post2},
	}

	for _, c := range cases {
		req, err := http.NewRequest("POST", url, strings.NewReader(c.body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		client := http.Client{}
		res, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%+v", res)
	}
}
