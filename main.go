package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {

	cert, err := tls.LoadX509KeyPair("certs/server-cert.pem", "certs/server-key.pem")
	if err != nil {
		log.Fatalf("Failed to load key pair: %s", err)
	}

	// Configurer le serveur HTTPS
	server := &http.Server{
		Addr:    ":8082",
		Handler: http.HandlerFunc(handler),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	log.Println("Starting server at https://localhost:8443")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("ListenAndServeTLS failed: %s", err)
	}
}

// Exemple de gestionnaire HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
