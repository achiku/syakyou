package main

import (
	"path/filepath"
	"testing"
)

func TestExt(t *testing.T) {
	cases := []struct {
		Path string
		Ext  string
	}{
		{Path: "hoge/fuga/file.txt", Ext: ".txt"},
		{Path: "hoge/fuga/file.jpg", Ext: ".jpg"},
	}
	for _, c := range cases {
		ext := filepath.Ext(c.Path)
		t.Logf("%s", ext)
		if ext != c.Ext {
			t.Errorf("want %s got %s", c.Ext, ext)
		}
	}
}
