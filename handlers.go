package main

import (
	"database/sql"
	"encoding/json"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
	"net/http"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func genres(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	// Open database connection
	db, err := sql.Open("sqlite3", "jrdd.db")

	// Make the query
	rows, err := db.Query("SELECT g.name, COUNT(g.ID), SUM(s.length) FROM genres g INNER JOIN songs s ON s.genre = g.ID GROUP BY g.ID")

	// Declare the variables to be used
  	var genre string
	var songs int
	var total_length int

	genres := Genres{}

    for rows.Next(){
    	err = rows.Scan(&genre, &songs, &total_length)	
	    aux_genre := Genre{ Name:genre, Songs: songs, Length: total_length}
	    genres = append(genres, aux_genre)
    }    
	
	// Check por errors
	if err != nil {
		fmt.Println(err)
	}
    
    // Close database connection
    db.Close()

    json.NewEncoder(w).Encode(genres)
}


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

func songs_by_artist(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	artist_name := pat.Param(ctx, "artist")

	// Open database connection
	db, err := sql.Open("sqlite3", "jrdd.db")

	// Make the query
	rows, err := db.Query("SELECT s.song, s.artist, g.name, s.length FROM songs s INNER JOIN genres g ON g.ID = s.genre WHERE s.artist = '" + artist_name + "'")

	// Declare the variables to be used
  	var song string
	var artist string
	var genre string
	var length int


	songs := Songs{}

    for rows.Next(){
    	err = rows.Scan(&song, &artist, &genre, &length)	
	    aux_song := Song{Artist: artist, Genre: genre, Length: length, Name: song}
	    songs = append(songs, aux_song)
    }    
	
	// Check por errors
	if err != nil {
		fmt.Println(err)
	}
    
    // Close database connection
    db.Close()

    json.NewEncoder(w).Encode(songs)

}

func songs_by_genre(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	genre_name := pat.Param(ctx, "genre")

	// Open database connection
	db, err := sql.Open("sqlite3", "jrdd.db")

	// Make the query
	rows, err := db.Query("SELECT s.song, s.artist, g.name, s.length FROM songs s INNER JOIN genres g ON g.ID = s.genre WHERE g.name = '" + genre_name + "'")

	// Declare the variables to be used
  	var song string
	var artist string
	var genre string
	var length int


	songs := Songs{}

    for rows.Next(){
    	err = rows.Scan(&song, &artist, &genre, &length)	
	    aux_song := Song{Artist: artist, Genre: genre, Length: length, Name: song}
	    songs = append(songs, aux_song)
    }    
	
	// Check por errors
	if err != nil {
		fmt.Println(err)
	}
    
    // Close database connection
    db.Close()

    json.NewEncoder(w).Encode(songs)

}



func songs_by_length(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	max_length := pat.Param(ctx, "max_length")
	min_length := pat.Param(ctx, "min_length")

	if max_length < min_length {
		fmt.Fprintf(w, "Error: The maximun length of the song should be greater or equal than the minimum length.")
		return
	}

	// Open database connection
	db, err := sql.Open("sqlite3", "jrdd.db")

	// Make the query
	rows, err := db.Query("SELECT s.song, s.artist, g.name, s.length FROM songs s INNER JOIN genres g ON g.ID = s.genre WHERE s.length <= " + max_length + " AND s.length >= " + min_length + " ORDER BY s.length ASC")

	// Declare the variables to be used
  	var song string
	var artist string
	var genre string
	var length int


	songs := Songs{}

    for rows.Next(){
    	err = rows.Scan(&song, &artist, &genre, &length)
	    aux_song := Song{Artist: artist, Genre: genre, Length: length, Name: song}
	    songs = append(songs, aux_song)
    }    
	
	// Check por errors
	if err != nil {
		fmt.Println(err)
	}
    
    // Close database connection
    db.Close()

    json.NewEncoder(w).Encode(songs)

}
