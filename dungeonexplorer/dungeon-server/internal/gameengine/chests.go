package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothChest struct {
	id int64
}

func (chest *ClothChest) Id() int64 {
	return chest.id
}

func (chest *ClothChest) Name() string {
	return "Cloth Chest"
}

func (chest *ClothChest) Description() string {
	return "A raggedy cloth chest"
}

func (chest *ClothChest) GoldValue() int64 {
	return 3
}

func (chest *ClothChest) Class() int64 {
	return 1
}

func (chest *ClothChest) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_CHEST
}
