package main

import(
	"fmt"
	"encoding/json"

	"licorne/integrations/client"
)

// This function fetch the content of a URL will return it as an
// array of bytes if retrieved successfully.


func main(){

	game := client.CreateGame()
	fmt.Println(game)

	game = client.GetGame(game.ID)
	fmt.Println(game)

	turn := client.GetTurn(game.ID)
	fmt.Println(turn)

	// turn := client.RollTurn(game.ID, turn.Number, turn.PlayerTurn)
	// fmt.Println



	content, _ := client.GetContent( client.BaseUrl() + "/v1/games")
		// if errs != nil {
		// 	// An error occurred while fetching the JSON
		// 	panic(errs)
		// }
		// Fill the record with the data from the JSON

		var data client.Data
		err := json.Unmarshal([]byte(content), &data)
		if err != nil {
			// An error occurred while converting our JSON to an object
			panic(err)
		}

		// for _, g :=range data.Games {
		// 	fmt.Println(g.Name)
		// }
}
