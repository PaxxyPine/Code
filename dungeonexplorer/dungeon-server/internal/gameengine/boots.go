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

type Runners struct {
	id int64
}

func (boot *Runners) Id() int64 {
	return boot.id
}

func (boot *Runners) Name() string {
	return "Runners"
}

func (boot *Runners) Description() string {
	return "Got the running drip"
}

func (boot *Runners) GoldValue() int64 {
	return 3
}

func (boot *Runners) Class() int64 {
	return 1
}

func (boot *Runners) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_BOOTS
}

type IronBoots struct {
	id int64
}

func (boot *IronBoots) Id() int64 {
	return boot.id
}

func (boot *IronBoots) Name() string {
	return "IronBoots"
}

func (boot *IronBoots) Description() string {
	return "boots made from hard cold steel"
}

func (boot *IronBoots) GoldValue() int64 {
	return 3
}

func (boot *IronBoots) Class() int64 {
	return 1
}

func (boot *IronBoots) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_BOOTS
}
