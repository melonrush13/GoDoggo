package main

import (
	"flag"
	"fmt" // methods for I/O Operations
	// templates for generating HTML output
	"net/http" // library with methods to implement HTTP clients and servers
)

const greeting string = "Hello, World"

type pData struct {
	Name  string
	Breed string
}

func init() {
}

func main() {

	http.HandleFunc("/", myHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)

	//filename := flag.String("file", "gopher.json", "the json file")
	flag.Parse()
	//fmt.Println("Using the story in %s", *filename)
	fmt.Println(greeting)

}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "about page\n")
	pD := pData{
		Name: "Mel",
	}
}
