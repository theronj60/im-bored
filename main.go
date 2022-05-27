package main

import (
	"fmt"
    "log"
	"io/ioutil"
    "net/http"
    "os"
)

// im-bored is main command
// call https://github.com/drewthoennes/Bored-API
// returns a suggestion

func main() {
	// welcome message (ascii preferred)
	welcome := "Welcome to I am Bored"
	fmt.Println(welcome)
	// ask for input
	var activity string
	fmt.Println("What kind of acitvity are you looking for? (random is an option)")
	fmt.Scanln(&activity)
	if activity == "random" {
		response, err := http.Get("http://www.boredapi.com/api/activity/")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	}
	// var people int
	// fmt.Println("How many people are there? (Enter a number)")
	// questions acitvity, how many people, etc.
}
