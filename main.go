package main

import (
	"crypto/tls"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

const domain = "http2.uutispuro.fi"

func main() {
	certCache := "tmp/certs"
	key, cert := "", ""

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache(certCache),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world with let's encrypt"))
	})

	server := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	server.ListenAndServeTLS(key, cert)
}
