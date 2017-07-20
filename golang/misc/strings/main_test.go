package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	s := []string{"path", "to", "file.json"}
	t.Logf("/%s", strings.Join(s, "/"))
}

type testData struct {
	In  string
	Out string
}

func TestParseInt(t *testing.T) {
	data := []struct {
		In  string
		Out int64
	}{
		{In: "983", Out: 983},
	}

	for _, d := range data {
		n, err := strconv.ParseInt(d.In, 10, 64)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%d", n)
	}
}

func TestGetLastFour(t *testing.T) {
	data := []testData{
		{"hello, world!", "rld!"},
		{"1234567890", "7890"},
		{"098", "098"},
		{"", ""},
	}

	for _, d := range data {
		s := getLastFour(d.In)
		if s != d.Out {
			t.Errorf("want %s got %s", d.Out, s)
		}
	}
}

func TestGetLastFour2(t *testing.T) {
	data := []testData{
		{"hello, world!", "rld!"},
		{"1234567890", "7890"},
		{"098", "s is less than four chars"},
		{"", "s is less than four chars"},
	}

	for _, d := range data {
		s, err := getLastFour2(d.In)
		if err != nil {
			if d.Out != err.Error() {
				t.Error(err)
			}
		} else {
			if s != d.Out {
				t.Errorf("want %s got %s", d.Out, s)
			}
		}
	}
}

func TestTremSpace(t *testing.T) {
	data := map[string]string{
		"test   ":     "test",
		"   test    ": "test",
		"   test":     "test",
		"  　　test":    "test",
		"  --test":    "--test",
	}

	for k, v := range data {
		s := strings.TrimSpace(k)
		if s != v {
			t.Errorf("want %s got %s", v, s)
		}
	}
}

func TestParseFloat(t *testing.T) {
	n := "1000.01111234445"
	f, err := strconv.ParseFloat(n, 64)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%f", f)
}

func TestTrim(t *testing.T) {
	s10 := "          "
	s5 := "     "
	cases := []struct {
		s string
		e string
	}{
		{s: "12345" + s10 + s10, e: "12345"},
		{s: "12345     " + s5 + s10, e: "12345"},
		{s: "1234567890" + s10 + s5, e: "1234567890"},
		{s: "あいうえお" + s10 + s10, e: "あいうえお"},
		{s: "あいうえお　　　　　" + s10 + s10, e: "あいうえお"},
		{s: "あいうえおかきくけこたちつてと" + s10, e: "あいうえおかきくけこたちつてと"},
	}

	for _, d := range cases {
		a := trim(d.s)
		if a != d.e {
			t.Errorf("want '%s' got '%s'", d.e, a)
		}
		t.Logf("|%s|->|%s|", d.s, a)
	}
}

func TestLower(t *testing.T) {
	t.Log(strings.ToLower("AAAAA"))
}

func TestTrimer(t *testing.T) {
	cases := []struct {
		str      string
		expected string
	}{
		{str: "test program              ", expected: "test program"},
		{str: "test program だよ             ", expected: "test program だよ"},
		{str: "あいうえお　だよ             ", expected: "あいうえお　だよ"},
	}

	for _, c := range cases {
		res := strings.TrimSpace(c.str)
		t.Logf("%s|", res)
		if res != c.expected {
			t.Errorf("want |%s| got |%s|", c.expected, res)
		}
	}
}

func TestCreateJSON(t *testing.T) {
	a := createJSON("1000")
	t.Logf("%s", a)
	b := createJSON2("1000")
	t.Logf("%s", b)
}
