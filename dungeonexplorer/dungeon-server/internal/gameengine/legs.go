package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothPants struct {
	id int64
}

func (leg *ClothPants) Id() int64 {
	return leg.id
}

func (leg *ClothPants) Name() string {
	return "Cloth Pants"
}

func (leg *ClothPants) Description() string {
	return "A raggedy pair of cloth pants"
}

func (leg *ClothPants) GoldValue() int64 {
	return 3
}

func (leg *ClothPants) Class() int64 {
	return 1
}

func (leg *ClothPants) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_LEGS
}
