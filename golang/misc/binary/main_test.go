package main

import "testing"

func TestBigEndian(t *testing.T) {
	seq := []byte{0x12, 0x34, 0x56, 0x78}
	var i uint32
	i = ((uint32(seq[0]) << 24) | (uint32(seq[1]) << 16) | (uint32(seq[2]) << 8) | uint32(seq[3]))
	t.Logf("%d, %x", i, i)
}
