package main

import (
	"Observer/observer"
)

func main() {
	gameRoom := &observer.Subject{}

	player1 := &observer.Player{Name: "Player 1"}
	player2 := &observer.Player{Name: "Player 2"}
	player3 := &observer.Player{Name: "Player 3"}

	gameRoom.Register(player1)
	gameRoom.Register(player2)
	gameRoom.Register(player3)

	gameRoom.Notify("Game is starting...")
	gameRoom.Notify("Player 1 made a move.")
	gameRoom.Notify("Player 2 made a move.")

	gameRoom.Deregister(player3)

	gameRoom.Notify("Player 3 left the game.")
	gameRoom.Notify("Player 1 made another move.")
}
