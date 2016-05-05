package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


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



func genres_test(pUrl string){
	url := pUrl

	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}

	var g Genres
	err = json.Unmarshal(body, &g)
	if err != nil{
		panic(err)
	}

	var json_marshal, err2 = json.MarshalIndent(g, "", "  ")
    if err2 != nil {
		fmt.Println(err2)
	}
    fmt.Println(string(json_marshal) + "\n\n")
}


func songs_test(pUrl string){
	url := pUrl

	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}

	var s Songs
	err = json.Unmarshal(body, &s)
	if err != nil{
		panic(err)
	}

	var json_marshal, err2 = json.MarshalIndent(s, "", "  ")
    if err2 != nil {
		fmt.Println(err2)
	}
    fmt.Println(string(json_marshal) + "\n\n")
}



func main(){

	fmt.Println("Test #1: Genres")
	test1 := "http://localhost:8000/genres"
	fmt.Println("url: " + test1)
	genres_test(test1)

	fmt.Println("Test #2: Songs by name")
	test2 := "http://localhost:8000/songs_by_name/Horacio"
	fmt.Println("url: " + test2)
	songs_test(test2)

	fmt.Println("Test #3: Songs by artist")
	test3 := "http://localhost:8000/songs_by_artist/Colornoise"
	fmt.Println("url: " + test3)
	songs_test(test3)

	fmt.Println("Test #4: Songs by genre")
	test4 := "http://localhost:8000/songs_by_genre/Pop"
	fmt.Println("url: " + test4)
	songs_test(test4)

	fmt.Println("Test #5: Songs by length")
	test5 := "http://localhost:8000/songs_by_length/200/300"
	fmt.Println("url: " + test5)
	songs_test(test5)

}