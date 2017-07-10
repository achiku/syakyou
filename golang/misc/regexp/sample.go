package main

import "regexp"

var r = regexp.MustCompile(`(^[a-z]+_[a-z]+):(.*)`)
