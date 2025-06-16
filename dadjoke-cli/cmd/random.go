/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long:  `This command fetches a random dad joke from the icanhazdadjoke api`,

	Run: func(cmd *cobra.Command, args []string) {
		jokeTerm, _ := cmd.Flags().GetString("term")

		if jokeTerm != "" {
			getRandomJokeWithTerm(jokeTerm)
		} else {
			getRandomJoke()
		}

	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.PersistentFlags().String("term", "", "A search term for a dad joke.")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

type SearchResult struct {
	Results    json.RawMessage `json:"results"`
	SearchTerm string          `json:"search_term"`
	Status     int             `json:"status"`
	TotalJokes int             `json:"total_jokes"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}
	fmt.Println(string(joke.Joke))
}

func getRandomJokeWithTerm(jokeTerm string) {
	total, results := getJokeDataWithTerm(jokeTerm)
	randomiseJokeList(total, results)
}

func randomiseJokeList (length int, jokeList []Joke) {
	rand.Seed(time.Now().Unix())

	min := 0
	max := length - 1

	if length <= 0 {
		err := fmt.Errorf("No jokes found with this term")
		fmt.Println(err.Error())
	} else {
		randomNum := min + rand.Intn(max - min)
		fmt.Println(jokeList[randomNum].Joke)
	}
}


func getJokeDataWithTerm(jokeTerm string) (totalJokes int, jokeList []Joke){
	url := fmt.Sprintf("https://icanhazdadjoke.com/search?term=%s", jokeTerm)

	responseBytes := getJokeData(url)
	jokeListRaw := SearchResult{}

	if err := json.Unmarshal(responseBytes, &jokeListRaw); err != nil {
		log.Printf("could not unmarshal responseBytes - %v", err)
	}

	jokes := []Joke{}
	if err := json.Unmarshal(jokeListRaw.Results, &jokes); err != nil {
		log.Printf("could not unmarshal jokeListRaw results - %v", err)
	}

	return jokeListRaw.TotalJokes, jokes
}



func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("could not request a dadjoke - %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/example/dadjoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request - %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Could not read response body - %v", err)
	}
	return responseBytes
}

// 🔧 Goal of This CLI
// You're building a Command-Line Interface tool that, when someone types:

// bash
// Copy
// Edit
// dadjoke random
// …it fetches a random dad joke from the internet and shows it in the terminal. This is done using the icanhazdadjoke.com API, which is a free API for dad jokes.

// 📦 Project Structure (using Cobra)
// Cobra gives you a standard structure to build CLI tools in Go. It uses:

// main.go: Starts the CLI.

// cmd/root.go: Defines the root command (dadjoke).

// cmd/random.go: Defines a subcommand (random) — what runs when you type dadjoke random.

// Your random.go file is responsible for getting and printing the joke.

// 🧠 Basic Concepts & Workflow
// Let’s break your code and explain what each part is doing.

// 1. Defining the random command
// go
// Copy
// Edit
// var randomCmd = &cobra.Command{
// 	Use:   "random",
// 	Short: "Get a random dad joke",
// 	Long:  `This command fetches a random dad joke from the icanhazdadjoke api`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		getRandomJoke()
// 	},
// }
// 🧠 What’s happening?
// This defines a subcommand named random.

// Use: The command name.

// Short & Long: Help text shown when you use --help.

// Run: This function runs when you type dadjoke random. Here, it calls getRandomJoke().

// 2. Registering the Command
// go
// Copy
// Edit
// func init() {
// 	rootCmd.AddCommand(randomCmd)
// }
// 🧠 What’s happening?
// You’re attaching this random command to the root command (dadjoke).

// So now you can run it from the terminal like:

// bash
// Copy
// Edit
// dadjoke random
// 3. Making the API Request: getRandomJoke()
// go
// Copy
// Edit
// func getRandomJoke() {
// 	url := "https://icanhazdadjoke.com/"
// 	responseBytes := getJokeData(url)

// 	joke := Joke{}
// 	if err := json.Unmarshal(responseBytes, &joke); err != nil {
// 		log.Printf("Could not unmarshal response - %v", err)
// 	}
// 	fmt.Println(joke.Joke)
// }
// 🧠 What’s happening?
// Sets the API URL.

// Calls getJokeData(url) to make the HTTP request and get the response.

// Converts the JSON response into a Go struct (type-safe data).

// Prints the joke.

// 4. Making HTTP GET Request: getJokeData()
// go
// Copy
// Edit
// func getJokeData(baseAPI string) []byte {
// 	request, err := http.NewRequest(http.MethodGet, baseAPI, nil)
// 	if err != nil {
// 		log.Printf("could not request a dadjoke - %v", err)
// 	}

// 	request.Header.Add("Accept", "application/json")
// 	request.Header.Add("User-Agent", "Dadjoke CLI (github.com/example/dadjoke)")

// 	response, err := http.DefaultClient.Do(request)
// 	if err != nil {
// 		log.Printf("Could not make a request - %v", err)
// 	}

// 	responseBytes, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		log.Println("Could not read response body - %v", err)
// 	}
// 	return responseBytes
// }
// 🧠 What’s happening?
// Builds an HTTP GET request to the dad joke API.

// Sets headers so the API sends a JSON response.

// Sends the request and gets the response.

// Reads the response body into a byte slice ([]byte).

// Returns this response data to be used in getRandomJoke().

// 5. The Joke Struct
// go
// Copy
// Edit
// type Joke struct {
// 	ID     string `json:"id"`
// 	Joke   string `json:"joke"`
// 	Status int    `json:"status"`
// }
// 🧠 What’s happening?
// This maps the JSON response to Go variables.

// For example, the API returns:

// json
// Copy
// Edit
// {
//   "id": "123",
//   "joke": "Why did the chicken cross the road? To get to the other side.",
//   "status": 200
// }
// Go uses json.Unmarshal to convert this into the Joke struct.

// 🧪 Complete Example Flow
// Let’s simulate what happens when you type:

// bash
// Copy
// Edit
// dadjoke random
// Cobra sees you typed the random command.

// It runs the Run: function → calls getRandomJoke().

// getRandomJoke() calls the joke API using getJokeData().

// The API returns JSON like:

// json
// Copy
// Edit
// {"id":"abc123", "joke":"I'm afraid for the calendar. Its days are numbered.", "status":200}
// JSON is parsed into your Go Joke struct.

// joke.Joke is printed:
// ✅ "I'm afraid for the calendar. Its days are numbered."
