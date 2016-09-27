package main

import "testing"

func TestSign(t *testing.T) {
	r := req{
		Key:    "hoge",
		Secret: "foo",
		Nonce:  100,
		URL:    "https://coincheck.com/api/ec/buttons",
		Body:   "hoge=foo&bar=boo",
	}
	expected := "6fb5a7f2a872adebd343fa8b463b3e2fe06cb088519f4836cda9a2b16bc30b97"

	s, err := sign(r)
	if err != nil {
		t.Fatal(err)
	}
	if s != expected {
		t.Errorf("want %s got %s", expected, s)
	}
}
