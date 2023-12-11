package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Boxer struct {
	Name   string
	Health int64
}

type BodyParts int64

const (
	Head BodyParts = 0
	Body BodyParts = 1
	Legs BodyParts = 2
)

func (cs BodyParts) String() string {
	switch cs {
	case Head:
		return "Head"
	case Body:
		return "Body"
	case Legs:
		return "Legs"
	}
	return "Unknown"
}

var listBodyParts = [3]BodyParts{Head, Body, Legs}

func (b Boxer) Attack(attackedBoxer *Boxer) {
	attackedBodyPart := listBodyParts[rand.Intn(2)]
	blockedBodyPart := listBodyParts[rand.Intn(2)]
	fmt.Println(b.Name, "attack:", attackedBodyPart, "|", attackedBoxer.Name, "block:", blockedBodyPart)

	if attackedBodyPart != blockedBodyPart {
		attackedBoxer.Health = attackedBoxer.Health - 1
	}
}

func getHealthString(health int64, maxHeath int64) string {
	filled := strings.Repeat("█", int(health))
	unfilled := strings.Repeat("░", int(max(maxHeath-health, 0)))
	return filled + unfilled
}

func main() {
	var maxHeath int64 = 3
	boxer1 := Boxer{Name: "Boxer1", Health: maxHeath}
	boxer2 := Boxer{Name: "Boxer2", Health: maxHeath}
	maxStep := 100
	step := 0
	fmt.Println("Let's get ready to rumble!")
	fmt.Println("---------- BOX!!! ----------")
	for boxer1.Health > 0 && boxer2.Health > 0 && step < maxStep {
		fmt.Printf("%s: %s ", boxer1.Name, getHealthString(boxer1.Health, maxHeath))
		fmt.Printf("%s: %s ", boxer2.Name, getHealthString(boxer2.Health, maxHeath))
		fmt.Println()
		step++
		boxer1.Attack(&boxer2)
		boxer2.Attack(&boxer1)
		fmt.Printf("***** STEP %v *****", step)
		fmt.Println()
	}
	if boxer1.Health == boxer2.Health {
		fmt.Println("Draw!")
	}
	if boxer1.Health > boxer2.Health {
		fmt.Printf("%s win!!!", boxer1.Name)
	}
	if boxer1.Health < boxer2.Health {
		fmt.Printf("%s win!!!", boxer2.Name)
	}

}
