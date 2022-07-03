package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothChest struct {
	id int64
}

func (helm *ClothChest) Name() string {
	return "Cloth Chest"
}

func (helm *ClothChest) Description() string {
	return "A raggedy cloth chest"
}

func (helm *ClothChest) GoldValue() int64 {
	return 3
}

func (helm *ClothChest) Class() int64 {
	return 1
}

func (helm *ClothChest) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_CHEST
}
