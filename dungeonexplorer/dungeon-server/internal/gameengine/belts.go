package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothBelt struct {
	id int64
}

func (helm *ClothBelt) Name() string {
	return "Cloth Belt"
}

func (helm *ClothBelt) Description() string {
	return "A raggedy cloth belt"
}

func (helm *ClothBelt) GoldValue() int64 {
	return 3
}

func (helm *ClothBelt) Class() int64 {
	return 1
}

func (helm *ClothBelt) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_BELT
}
