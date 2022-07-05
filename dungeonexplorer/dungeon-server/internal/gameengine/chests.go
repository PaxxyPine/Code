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
	return
}

func (chest *ClothChest) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_CHEST
}

type BlackHawkPlate struct {
	id int64
}

func (chest *BlackHawkPlate) Id() int64 {
	return chest.id
}

func (chest *BlackHawkPlate) Name() string {
	return "Black Hawk chestplate"
}

func (chest *BlackHawkPlate) Description() string {
	return "pure evil held in a plate that belongs on your chest"
}

func (chest *BlackHawkPlate) GoldValue() int64 {
	return 60
}

func (chest *BlackHawkPlate) Class() int64 {
	return 6
}

func (chest *BlackHawkPlate) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_CHEST
}

type InronChestPlate struct {
	id int64
}

func (chest *InronChestPlate) Id() int64 {
	return chest.id
}

func (chest *InronChestPlate) Name() string {
	return "Inron chestplate"
}

func (chest *InronChestPlate) Description() string {
	return "wow 1000000 pounds is pretty light!!"
}
func (chest *InronChestPlate) GoldValue() int64 {
	return 50
}

func (chest *InronChestPlate) Class() int64 {
	return 5
}

func (chest *InronChestPlate) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_CHEST
}
