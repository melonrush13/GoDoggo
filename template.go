package main

import (
	"html/template"
	"log"
	"net/http"
)

type dogData struct {
	Name  string
	Breed string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	data := dogData{
		Name:  "Bear",
		Breed: "Wheaton Terrier",
	}

	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		log.Print("TEMPLATE PARSING ERROR ", err)
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Print("EXECUTE ERROR ", err)
	}
}
