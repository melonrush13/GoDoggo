package main

import (
	"encoding/json"
	"fmt" // methods for I/O Operations
	"io/ioutil"

	// library with methods to implement HTTP clients and servers
	"os"
)

const greeting string = "Hello, World"

// Dogs struct which contains
// an array of dogs
type Dogs struct {
	Dogs []struct {
		Name   string `json:"name"`
		Breed  string `json:"breed"`
		Age    int    `json:"age"`
		Owners struct {
			Dad string `json:"dad"`
			Mom string `json:"mom"`
		} `json:"owners"`
	} `json:"dogs"`
}

func init() {
}

func main() {
	//open jsonFile
	jsonFile, err := os.Open("dogs.json")
	// if error on os.Open, handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("opened dogs.json")
	defer jsonFile.Close()

	// read our opened file as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initalize dog array
	var dogs Dogs

	// unmarshal bytearray
	json.Unmarshal(byteValue, &dogs)

	//iterate through json
	for i := 0; i < len(dogs.Dogs); i++ {
		fmt.Println("Dog Name: " + dogs.Dogs[i].Name)
		fmt.Println("Dog Breed: " + dogs.Dogs[i].Breed)
		fmt.Println("Dog Owners: "+dogs.Dogs[i].Owners.Dad, dogs.Dogs[i].Owners.Mom)
	}
}
