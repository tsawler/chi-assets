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
	mux := chi.NewRouter()
	fs.WalkDir(assets, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})

	mux.HandleFunc("/", Home)

	mux.Handle("/static/*", http.FileServer(http.FS(assets)))

	return mux
}
