package main

import (
	"flag"
	"log"
	"os"
	"webdevutils"
)

func main() {
	pwd := os.Getenv("PWD")

	var ipAddr, serveDir, keyfile, certfile string
	var serveTls bool
	flag.StringVar(&ipAddr, "listen", ":8000",
		"Listen IP and port address, e.g. 0.0.0.0:8000")
	flag.StringVar(&serveDir, "serve", pwd, "Directory to serve file from")
	flag.BoolVar(&serveTls, "tls", false, "Whether to server using TLS")
	flag.StringVar(&keyfile, "key", "",
		"The key file to use for TLS serving")
	flag.StringVar(&certfile, "cert", "",
		"The cert file to use for TLS serving")
	flag.Parse()

	log.Println("Listening on: ", ipAddr)
	log.Println("Serving from: ", serveDir)

	if serveTls {
		if len(keyfile) == 0 || len(certfile) == 0 {
			log.Fatal("Must provide both -key and -cert options for TLS mode.")
		}

		log.Println("Using protocol: TLS")
		log.Println("Using keyfile: ", keyfile)
		log.Println("Using crtfile: ", certfile)

		err := webdevutils.StaticServerTLS(ipAddr, certfile, keyfile, serveDir)
		if err != nil {
			log.Panic(err)
		}
	} else {
		err := webdevutils.StaticServer(ipAddr, serveDir)
		if err != nil {
			log.Panic(err)
		}
	}
}
