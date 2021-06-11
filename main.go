package main


import (
	"log"
	"net/http"
	"os"

	"github.com/nathanramli/go-rest-api/home"
	"github.com/nathanramli/go-rest-api/server"
)

/**
	Untuk set environment variable pada windows. Gunakan:
	setx ENV_VARIABLE_NAME "value"
	Note: Buka command prompt baru untuk melihat hasil
 */

var (
	CertFile	= "./certs/localhost.crt"
	KeyFile 	= "./certs/localhost.key"
	ServiceAddr = "127.0.0.1:8080"
)

func main() {
	logger := log.New(os.Stdout, "backend_", log.LstdFlags | log.Lshortfile)
	h := home.NewHandlers(logger)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, ServiceAddr)

	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		log.Fatalf("%v", err)
	}
}