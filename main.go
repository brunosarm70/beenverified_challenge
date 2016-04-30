package main

import (
    _ "github.com/mattn/go-sqlite3"
	"net/http"
	"goji.io"
	"goji.io/pat"
)


func main() {
	
	// Available routes
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/songs_by_name/:name"), songs_by_name)
	mux.HandleFuncC(pat.Get("/songs_by_artist/:artist"), songs_by_artist)
	mux.HandleFuncC(pat.Get("/songs_by_genre/:genre"), songs_by_genre)
	mux.HandleFuncC(pat.Get("/songs_by_length/:min_length/:max_length"), songs_by_length)
	mux.HandleFuncC(pat.Get("/genres"), genres)

	// Host listening
	http.ListenAndServe("localhost:8000", mux)
}