package main

import (
	"fmt"
)

type Player struct {
	space int
	score int
}

func TwentyOne() {
	player1 := Player{
		3,
		0,
	}
	player2 := Player{
		7,
		0,
	}
	players := []Player{player1, player2}
	wins := make(map[int]int)
	DiracDice(players, 0, wins, 1)

	fmt.Printf("Wins: %v\n", wins)

}

func DiracDice(players []Player, playerTurn int, wins map[int]int, multiplier int) {
	if players[0].score >= 21 {
		wins[0] += multiplier
		return
	}
	if players[1].score >= 21 {
		wins[1] += multiplier
		return
	}
	player := players[playerTurn]

	diracValues := make(map[int]int)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				diracValues[i+j+k+3]++
			}
		}
	}
	for diceValue, diceCount := range diracValues {
		newSpace := (player.space+diceValue-1)%10 + 1
		newPlayer := Player{
			newSpace,
			player.score + newSpace,
		}
		if playerTurn == 0 {
			DiracDice([]Player{newPlayer, players[1]}, (playerTurn+1)%2, wins, multiplier*diceCount)
		} else {
			DiracDice([]Player{players[0], newPlayer}, (playerTurn+1)%2, wins, multiplier*diceCount)
		}
	}
}
