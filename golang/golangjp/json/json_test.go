package json

import "testing"

func TestHello(t *testing.T) {
	greeting := Hello()
	if greeting != "hello!" {
		t.Errorf("want hello! but %s", greeting)
	}
}
