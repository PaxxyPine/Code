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

type ChainMailPants struct {
	id int64
}

func (leg *ChainMailPants) Id() int64 {
	return leg.id
}

func (leg *ChainMailPants) Name() string {
	return "Cloth Pants"
}

func (leg *ChainMailPants) Description() string {
	return "you can also make these with bottle caps!"
}

func (leg *ChainMailPants) GoldValue() int64 {
	return 35
}

func (leg *ChainMailPants) Class() int64 {
	return 3
}

func (leg *ChainMailPants) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_LEGS
}

type IronPants struct {
	id int64
}

func (leg *IronPants) Id() int64 {
	return leg.id
}

func (leg *IronPants) Name() string {
	return "Iron Pants"
}

func (leg *IronPants) Description() string {
	return "you can also make these with bottle caps!"
}

func (leg *IronPants) GoldValue() int64 {
	return 45
}

func (leg *IronPants) Class() int64 {
	return 4
}

func (leg *IronPants) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_LEGS
}
