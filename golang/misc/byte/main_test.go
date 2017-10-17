package main

import "testing"

func TestByte(t *testing.T) {
	msg := []byte("000A0000")
	t.Logf("%b", msg)
	s := []byte("\xd3\xd6\xc7\xd5")
	s1 := []byte("\xd3\xd6\xc7\xd5")
	t.Logf("%b", s)
	t.Logf("%b", s1)
}

func TestByteSlice(t *testing.T) {
	msg := []byte("000A0020LOGN1223023data")
	t.Logf("%s", msg)
	h := msg[:8]
	t.Logf("%s", h)

	a := msg[8:12]
	t.Logf("%s", a)

	b := msg[12 : len(msg)-1]
	t.Logf("%s", b)
}
