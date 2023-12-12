package main

import (
	"fmt"
	"strings"
)

func getHealthString(health int64, maxHeath int64) string {
	filled := strings.Repeat("█", int(health))
	unfilled := strings.Repeat("░", int(max(maxHeath-health, 0)))
	return filled + unfilled
}

func main() {
	var maxHeath int64 = 3
	boxer1 := CreateBoxer("Boxer1", maxHeath)
	boxer2 := CreateBoxer("Boxer2", maxHeath)
	game := CreateNewGame(180, &boxer1, &boxer2)
	fmt.Println("Let's get ready to rumble!")
	fmt.Println("---------- BOX!!! ----------")
	for !game.GameOver {
		fmt.Printf("***** STEP %v *****", game.CurrentStep)
		fmt.Println()
		fmt.Printf("%s: %s ", boxer1.Name, getHealthString(game.Boxer1.Health, maxHeath))
		fmt.Printf("%s: %s ", boxer2.Name, getHealthString(game.Boxer2.Health, maxHeath))
		fmt.Println()
		game.Next()
		fmt.Println(game.Boxer1.Name, "attack:", game.Boxer1.AttackedBodyPart, "|", game.Boxer2.Name, "block:", game.Boxer2.BlockedBodyPart)
		fmt.Println(game.Boxer2.Name, "attack:", game.Boxer2.AttackedBodyPart, "|", game.Boxer1.Name, "block:", game.Boxer1.BlockedBodyPart)
	}
	if game.CurrentWinner == nil {
		fmt.Println("Draw!")
	} else {
		fmt.Printf("The %s win!!!", game.CurrentWinner.Name)
	}
}
