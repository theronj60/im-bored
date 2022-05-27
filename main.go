package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// im-bored is main command
// call https://github.com/drewthoennes/Bored-API
// returns a suggestion

type Response struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int     `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float64 `json:"accessibility"`
}

// function for calling random?
func random() {
	response, err := http.Get("http://www.boredapi.com/api/activity/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Activity)
	fmt.Println(responseObject.Link)
}

func main() {
	// welcome message (ascii preferred)
	welcome := "Welcome to I am Bored"
	fmt.Println(welcome)
	// ask for input
	var activity string
	fmt.Println("What kind of acitvity are you looking for? (random is an option)")
	fmt.Scanln(&activity)
	if activity == "random" {
		random()
	}
	// var people int
	// fmt.Println("How many people are there? (Enter a number)")
	// questions acitvity, how many people, etc.
}
