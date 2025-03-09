package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	// Note that no error handling is necessary, as Read always succeeds.
	secret := make([]byte, 32)
	fmt.Println(secret)
	_, err := rand.Read(secret)
	if err != nil {
		log.Fatal("Error generating secret:", err)
	}
	encodedSecret := base64.StdEncoding.EncodeToString(secret)
    fmt.Println("SECRET_KEY=", encodedSecret)
}
