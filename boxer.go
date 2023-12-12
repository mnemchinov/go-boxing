package main

import (
	"math/rand"
)

type Boxer struct {
	Name             string   `json:"name"`
	Health           int64    `json:"health"`
	AttackedBodyPart BodyPart `json:"attacked_body_part"`
	BlockedBodyPart  BodyPart `json:"blocked_body_part"`
}

func getTagetBodyPart() BodyPart {
	bodyPart := BodyParts[rand.Intn(3)]
	return bodyPart
}

func (boxer Boxer) Attack(attackedBoxer *Boxer) (BodyPart, BodyPart) {
	attackedBodyPart, attackedPower := boxer.Hit()
	blockedBodyPart, blockedPower := attackedBoxer.Block()
	powerDiff := max(attackedPower-blockedPower, 0)

	if attackedBodyPart != blockedBodyPart {
		attackedBoxer.Health = attackedBoxer.Health - attackedPower
	} else {
		attackedBoxer.Health = attackedBoxer.Health - powerDiff
	}
	return attackedBodyPart, blockedBodyPart
}

func (boxer Boxer) Hit() (BodyPart, int64) {
	bodyPart := getTagetBodyPart()
	return bodyPart, 1
}

func (boxer Boxer) Block() (BodyPart, int64) {
	bodyPart := getTagetBodyPart()
	return bodyPart, 2
}

func CreateBoxer(name string, health int64) Boxer {
	newBoxer := Boxer{Name: name, Health: health}
	return newBoxer
}
