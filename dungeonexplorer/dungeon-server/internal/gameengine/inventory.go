package gameengine

import "github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"

type Item struct {
	name        string
	description string
	value       int
	weight      int
}

type Weapon struct {
	damage  int
	element Code.Elements
	radius  int
}

type Armor struct {
	aClass  int
	element Code.Elements
	aType   Code.ArmorType
}
