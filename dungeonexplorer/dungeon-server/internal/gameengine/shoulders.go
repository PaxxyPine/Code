package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothShoulders struct {
	id int64
}

func (shoulders *ClothShoulders) Id() int64 {
	return shoulders.id
}

func (shoulders *ClothShoulders) Name() string {
	return "Cloth Shoulders"
}

func (shoulders *ClothShoulders) Description() string {
	return "A raggedy cloth shoulder sash"
}

func (shoulders *ClothShoulders) GoldValue() int64 {
	return 3
}

func (shoulders *ClothShoulders) Class() int64 {
	return 1
}

func (shoulders *ClothShoulders) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_SHOULDERS
}
