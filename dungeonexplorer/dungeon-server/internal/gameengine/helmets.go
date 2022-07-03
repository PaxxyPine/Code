package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothHat struct {
	id int64
}

func (helm *ClothHat) Name() string {
	return "Cloth Hat"
}

func (helm *ClothHat) Description() string {
	return "A raggedy cloth hat"
}

func (helm *ClothHat) GoldValue() int64 {
	return 3
}

func (helm *ClothHat) Class() int64 {
	return 1
}

func (helm *ClothHat) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_HELM
}
