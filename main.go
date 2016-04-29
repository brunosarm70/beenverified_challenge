package main

import (
	"database/sql"
	"encoding/json"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
	"net/http"
	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)


// Types  ==========================================================
type Song struct {
    Name      string    `json:"song"`
    Artist    string    `json:"artist"`
    Genre     string 	`json:"genre"`
    Length    int		`json:"length"`
}

type Songs []Song
// =================================================================



func songs_by_name(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Get the name of the song
	name := pat.Param(ctx, "name")

	// Open database connection
	db, err := sql.Open("sqlite3", "jrdd.db")

	// Make the query
	rows, err := db.Query("SELECT s.song, s.artist, g.name, s.length FROM songs s INNER JOIN genres g ON g.ID = s.genre WHERE s.song = '" + name + "'")

	// Declare the variables to be used
  	var song string
	var artist string
	var genre string
	var length int

	// Create the empty list of "Song"s
	songs := Songs{}

    for rows.Next(){
    	// Assign the columns to its variables
    	err = rows.Scan(&song, &artist, &genre, &length)	
    	// Create a Song and append it to the list
	    aux_song := Song{Artist: artist, Genre: genre, Length: length, Name: song}
	    songs = append(songs, aux_song)
    }    
	
	// Check por errors
	if err != nil {
		fmt.Println(err)
	}
	
    // Close database connection
    db.Close()

    // Send the list of json's
    json.NewEncoder(w).Encode(songs)
}




func main() {
	
	// Available routes
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/songs_by_name/:name"), songs_by_name)

	// Host listening
	http.ListenAndServe("localhost:8000", mux)
}