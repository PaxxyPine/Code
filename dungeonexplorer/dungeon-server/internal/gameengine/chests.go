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
	return 30
}

func (chest *ClothChest) Class() int64 {
	return 2
}

func (chest *ClothChest) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_CHEST
}
type Chainmailchestplate struct {
	id int64
}

func (chest *Chainmailchestplate) Id() int64 {
	return chest.id
}

func (chest *Chainmailchestplate) Name() string {
	return "Chainmail chestplate"
}

func (chest *Chainmailchestplate-) Description() string {
	return "you could make this out of can caps!"
}

func (chest *Chainmailchestplate) GoldValue() int64 {
	return 40
}

func (chest *Chainmailchestplate) Class() int64 {
	return 4
}

func (chest *Chainmailchestplate) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_CHEST
}
