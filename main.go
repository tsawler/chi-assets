package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	srv := &http.Server{
		Addr:              ":4000",
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      600 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
