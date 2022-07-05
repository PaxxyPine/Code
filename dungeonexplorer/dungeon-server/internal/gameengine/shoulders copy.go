package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothShoulders struct {
	id int64
}

func (shoulders *ClothShoulders) Id() int64 {
	return shoulders.id
}

func (shoulders *ClothShoulders) Name() string {
	return "Cloth Shoulders"
}

func (shoulders *ClothShoulders) Description() string {
	return "A raggedy cloth shoulder sash"
}

func (shoulders *ClothShoulders) GoldValue() int64 {
	return 3
}

func (shoulders *ClothShoulders) Class() int64 {
	return 1
}

func (shoulders *ClothShoulders) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_SHOULDERS
}

type ShoulderSpikes struct {
	id int64
}

func (shoulders *ShoulderSpikes) Id() int64 {
	return shoulders.id
}

func (shoulders *ShoulderSpikes) Name() string {
	return "Shoulder Spikes"
}

func (shoulders *ShoulderSpikes) Description() string {
	return "a very pointy and fancy plate for the shoulder"
}

func (shoulders *ShoulderSpikes) GoldValue() int64 {
	return 45
}

func (shoulders *ShoulderSpikes) Class() int64 {
	return 3
}

func (shoulders *ShoulderSpikes) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_SHOULDERS
}

type ShoulderPad struct {
	id int64
}

func (shoulders *ShoulderPad) Id() int64 {
	return ShoulderPad.id
}

func (shoulders *ShoulderPad) Name() string {
	return "Shoulder Pad"
}

func (shoulders *ShoulderPad) Description() string {
	return "just some cloth on your shoulders"
}

func (shoulders *ShoulderPad) GoldValue() int64 {
	return 45
}

func (shoulders *ShoulderPad) Class() int64 {
	return 3
}

func (shoulders *ShoulderPad) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_SHOULDERS
}
