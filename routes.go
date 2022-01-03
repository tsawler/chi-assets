package main

import (
	"embed"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io/fs"
	"net/http"
)

//go:embed static
var assets embed.FS

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func routes() *chi.Mux {
	// create the router
	mux := chi.NewRouter()

	// this just prints the full path of all assets to the console
	fs.WalkDir(assets, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})

	// a sample handler for the home page
	mux.HandleFunc("/", Home)

	// handle serving embedded assets
	mux.Handle("/static/*", http.FileServer(http.FS(assets)))

	return mux
}
