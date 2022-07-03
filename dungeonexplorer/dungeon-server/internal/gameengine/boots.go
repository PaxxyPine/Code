package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothBoots struct {
	id int64
}

func (boot *ClothBoots) Id() int64 {
	return boot.id
}

func (boot *ClothBoots) Name() string {
	return "Cloth Boots"
}

func (boot *ClothBoots) Description() string {
	return "A raggedy pair of cloth boots"
}

func (boot *ClothBoots) GoldValue() int64 {
	return 3
}

func (boot *ClothBoots) Class() int64 {
	return 1
}

func (boot *ClothBoots) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_BOOTS
}
