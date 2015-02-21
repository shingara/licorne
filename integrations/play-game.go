package main

import(
	"encoding/json"
	"fmt"
	"os"
	// "bufio"

	"github.com/parnurzeal/gorequest"
)

const (
)

func BaseUrl() string {
	host := os.Getenv("WEB_1_PORT_9000_TCP_ADDR")
	port := os.Getenv("WEB_1_PORT_9000_TCP_PORT")
	return "http://" + host + ":" + port
}

// This function fetch the content of a URL will return it as an
// array of bytes if retrieved successfully.
func getContent(url string) (string, []error) {
	_, body, errs := gorequest.New().Get(url).End()
	// At this point we're done - simply return the bytes
	return body, errs
}

func postContent(url string, params interface{}) (string, []error) {
	_, body, errs := gorequest.New().Post(url).
		Send(params).
		End()
	if len(errs) > 0 {
		fmt.Println(errs)
	}
	return body, errs
}

type (
	Data struct {
		Games []Game `json:"data"`
	}
  Game struct {
    Name string `json:"name"`
	}
)

func createGame() Game {
	fmt.Println("Create a game")
	body, errs := postContent(
		BaseUrl() + "/v1/games",
		`{"name":"first_game", "players": [ {"name": "cyril"}, { "name": "lena" } ] }`,
	)
	if len(errs) > 0 {
		panic(errs)
	}

	var game Game
	err := json.Unmarshal([]byte(body), &game)
	if err != nil {
		fmt.Println(game.Name)
	}
	return game
}

func main(){

	game := createGame()
	fmt.Println(game)


	content, _ := getContent( BaseUrl() + "/v1/games")
		// if errs != nil {
		// 	// An error occurred while fetching the JSON
		// 	panic(errs)
		// }
		// Fill the record with the data from the JSON

		var data Data
		err := json.Unmarshal([]byte(content), &data)
		if err != nil {
			// An error occurred while converting our JSON to an object
			panic(err)
		}

		// for _, g :=range data.Games {
		// 	fmt.Println(g.Name)
		// }
}
