package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothBelt struct {
	id int64
}

func (belt *ClothBelt) Id() int64 {
	return belt.id
}

func (belt *ClothBelt) Name() string {
	return "Cloth Belt"
}

func (belt *ClothBelt) Description() string {
	return "A raggedy cloth belt"
}

func (belt *ClothBelt) GoldValue() int64 {
	return 3
}

func (belt *ClothBelt) Class() int64 {
	return 1
}

func (belt *ClothBelt) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_BELT
}

type IronBelt struct {
	id int64
}

func (belt *IronBelt) Name() string {
	return "Iron Belt"
}

func (belt *IronBelt) Description() string {
	return "A rich iron belt"
}

func (belt *IronBelt) GoldValue() int64 {
	return 20
}

func (belt *IronBelt) Class() int64 {
	return 2
}

func (belt *IronBelt) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_BELT
}

type WitchBelt struct {
	id int64
}

func (belt *WitchBelt) Name() string {
	return "Witch Belt"
}

func (belt *WitchBelt) Description() string {
	return "The witchs old master belt"
}

func (belt *WitchBelt) GoldValue() int64 {
	return 30
}

func (belt *WitchBelt) Class() int64 {
	return 3
}

func (belt *WitchBelt) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_BELT
}
