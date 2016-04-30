package main

// Types  ==========================================================
type Song struct {
    Name      string    `json:"song"`
    Artist    string    `json:"artist"`
    Genre     string 	`json:"genre"`
    Length    int		`json:"length"`
}

type Songs []Song

type Genre struct {
	Name 	  string    `json:"genre"`
	Songs     int       `json:"songs"`
	Length    int       `json:"length"`
}

type Genres []Genre
// =================================================================


