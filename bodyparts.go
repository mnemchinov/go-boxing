package main

type BodyPart int64

const (
	Head BodyPart = 0
	Body BodyPart = 1
	Legs BodyPart = 2
)

func (cs BodyPart) String() string {
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

var BodyParts = [3]BodyPart{Head, Body, Legs}
