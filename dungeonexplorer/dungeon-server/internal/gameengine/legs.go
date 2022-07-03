package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothPants struct {
	id int64
}

func (helm *ClothPants) Name() string {
	return "Cloth Pants"
}

func (helm *ClothPants) Description() string {
	return "A raggedy pair of cloth pants"
}

func (helm *ClothPants) GoldValue() int64 {
	return 3
}

func (helm *ClothPants) Class() int64 {
	return 1
}

func (helm *ClothPants) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_LEGS
}
