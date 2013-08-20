package main

import (
	"flag"
	"log"
	"net/http"
	"os"
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

		_, err := os.Stat(keyfile)
		if err != nil {
			log.Fatal("keyfile: ", err)
		}

		_, err = os.Stat(certfile)
		if err != nil {
			log.Fatal("certfile: ", err)
		}
		log.Println("Using protocol: TLS")
		log.Println("Using keyfile: ", keyfile)
		log.Println("Using crtfile: ", certfile)

		log.Panic(http.ListenAndServeTLS(ipAddr, certfile, keyfile,
			http.FileServer(http.Dir(serveDir))))
	} else {
		log.Panic(http.ListenAndServe(ipAddr,
			http.FileServer(http.Dir(serveDir))))
	}
}
