package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateCode() string {
	return fmt.Sprintf("%d", rand.Intn(100))
}
