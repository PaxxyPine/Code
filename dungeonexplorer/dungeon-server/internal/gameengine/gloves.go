package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothGloves struct {
	id int64
}

func (glove *ClothGloves) Id() int64 {
	return glove.id
}

func (glove *ClothGloves) Name() string {
	return "Cloth Gloves"
}

func (glove *ClothGloves) Description() string {
	return "A raggedy cloth pair of gloves"
}

func (glove *ClothGloves) GoldValue() int64 {
	return 3
}

func (glove *ClothGloves) Class() int64 {
	return 1
}

func (glove *ClothGloves) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_GOVES
}

type BoxingGloves struct {
	id int64
}

func (glove *BoxingGloves) Id() int64 {
	return glove.id
}

func (glove *BoxingGloves) Name() string {
	return "BoxingGloves"
}

func (glove *BoxingGloves) Description() string {
	return "Get in the ring!"
}

func (glove *BoxingGloves) GoldValue() int64 {
	return 30
}

func (glove *BoxingGloves) Class() int64 {
	return 3
}

func (glove *BoxingGloves) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_GOVES
}

type spikedGloves struct {
	id int64
}

func (glove *spikedGloves) Id() int64 {
	return glove.id
}

func (glove *spikedGloves) Name() string {
	return "spiked Gloves"
}

func (glove *spikedGloves) Description() string {
	return "gloves that are spiked. what did you think?"
}

func (glove *spikedGloves) GoldValue() int64 {
	return 50
}

func (glove *spikedGloves) Class() int64 {
	return 5
}

func (glove *spikedGloves) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_GOVES
}

type ElementalGloves struct {
	id int64
}

func (glove *ElementalGloves) Id() int64 {
	return glove.id
}

func (glove *ElementalGloves) Name() string {
	return "Elemental Gloves"
}

func (glove *ElementalGloves) Description() string {
	return "harness the power of all the elements"
}

func (glove *ElementalGloves) GoldValue() int64 {
	return 60
}

func (glove *ElementalGloves) Class() int64 {
	return 4
}

func (glove *ElementalGloves) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_GOVES
}
