package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothBoots struct {
	id int64
}

func (helm *ClothBoots) Name() string {
	return "Cloth Boots"
}

func (helm *ClothBoots) Description() string {
	return "A raggedy pair of cloth boots"
}

func (helm *ClothBoots) GoldValue() int64 {
	return 3
}

func (helm *ClothBoots) Class() int64 {
	return 1
}

func (helm *ClothBoots) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_BOOTS
}
