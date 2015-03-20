package client

import (
	"fmt"
	"time"
	"encoding/json"
)

type (
	PlayerStatus struct {
		PlayerId string
		Position int
		NbDiamond int
	}

	Turn struct {
		GameId string
		Number int

		StartTurn time.Time
		PlayerTurn string
		EndTurn time.Time

		PlayersEndStatus []PlayerStatus
	}
)

func parseJsonTurn(body string) (turn Turn) {
	err := json.Unmarshal([]byte(body), &turn)
	if err != nil {
		fmt.Println(turn.Number)
	}
	return turn
}

func GetTurn(game_id string) Turn {
	fmt.Println("get Turn")
	body, errs := GetContent(
		BaseUrl() + "/v1/games/" + game_id + "/turn",
	)
	if len(errs) > 0 {
		panic(errs)
	}
	return parseJsonTurn(body)
}
