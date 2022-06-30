package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/manifoldco/promptui"
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

func prompt() string {
	prompt := promptui.Select{
		Label: "Which activity type should we search for?",
		Items: []string{"random", "education", "recreational", "social", "diy", "charity", "cooking", "relaxation", "music", "busywork"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return result
}

// function for calling random
func random() {
	responseUrl := "http://www.boredapi.com/api/activity/"
	
	response := getResponse(responseUrl)
	fmt.Println(response.Activity)
	fmt.Println(response.Link)
}

func getResponse(query string) Response {
	queryResponse, err := http.Get(query)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	queryData, err := ioutil.ReadAll(queryResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseQuery Response
	json.Unmarshal(queryData, &responseQuery)

	return responseQuery
}

func main() {
	// welcome message (ascii preferred)
	welcome := "\nWelcome to I am Bored\n"
	fmt.Println(welcome)

	// ask for input
	// prompt func
	activityType := prompt()

	if activityType == "random" {
		random()
		return
	}
	
	query := "http://www.boredapi.com/api/activity?type=" + activityType

	// @TODO create response function, accepts variable returns string
	// @TODO create loop to continue looking for something to do until one is found

	response := getResponse(query)
	
	fmt.Println(response.Activity)
}
