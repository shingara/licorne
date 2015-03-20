package client

import(
	"encoding/json"

	"fmt"
)

type (
	Data struct {
		Games []Game `json:"data"`
	}

	Player struct {
		Name string `json:"name"`
		Color string `json:"color"`
	}

  Game struct {
		ID string
    Name string `json:"name"`
		Player []Player `json:"players"`
	}
)

func parseJsonGame(body string) Game {
	var game Game
	err := json.Unmarshal([]byte(body), &game)
	if err != nil {
		fmt.Println(game.Name)
	}
	return game
}

func CreateGame() Game {
	fmt.Println("Create a game")
	body, errs := PostContent(
		BaseUrl() + "/v1/games",
		`{"name":"first_game", "players": [ {"name": "cyril"}, { "name": "lena" } ] }`,
	)
	if len(errs) > 0 {
		panic(errs)
	}
	game := parseJsonGame(body)
	return game
}

func GetGame(id string) Game {
	fmt.Println("get game " + id)
	body, errs := GetContent(
		BaseUrl() + "/v1/games/" + id,
	)
	if len(errs) > 0 {
		panic(errs)
	}
	game := parseJsonGame(body)
	return game
}
