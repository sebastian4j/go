package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title:  "Titulo 1",
		Year:   1942,
		Color:  false,
		Actors: []string{"A1", "A2"},
	},
	{
		Title:  "Titulo 2",
		Year:   1999,
		Color:  true,
		Actors: []string{"B1", "B2"},
	},
}

func main() {
	data, err := json.Marshal(movies)
	fmt.Println(string(data))
	if err != nil {
		log.Fatalf("json marshaling error, %s", err)
	}

	data, _ = json.MarshalIndent(movies, "", " ")
	fmt.Printf("%s\n", data)

	// obtener solo el Title de cada json object
	var titles []struct {
		Title string
	}

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("error en unmarshal: %s", err)
	}
	fmt.Println(titles)
}
