package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

type req struct {
	Key    string
	Secret string
	Nonce  int64
	URL    string
	Body   string
}

func sign(r req) (string, error) {
	message := fmt.Sprintf("%d%s%s", r.Nonce, r.URL, r.Body)
	log.Printf("Message: %s", message)
	log.Printf("Key: %s", r.Secret)
	sig := hmac.New(sha256.New, []byte(r.Secret))
	sig.Write([]byte(message))
	log.Printf("%x", sig.Sum(nil))
	return hex.EncodeToString(sig.Sum(nil)), nil
}
