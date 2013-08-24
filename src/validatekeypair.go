package main

import (
	"log"
	"os"
	"webdevutils"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		log.Fatal("Usage: validatekeypair CERT KEY")
	}
	cert := args[1]
	key := args[2]
	err := webdevutils.ValidateX509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Valid X509 key pair")
	}
}
