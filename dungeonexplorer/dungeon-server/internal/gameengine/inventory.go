package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type Item interface {
	Id() int64
	Name() string
	Description() string
	GoldValue() int64
}

type Weapon interface {
	Id() int64
	Damage() int64
	Element() Code.Elements
	Speed() int64
}

type Armor interface {
	Id() int64
	Class() int64
	Type() Code.ArmorType
}
