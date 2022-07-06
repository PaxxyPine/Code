package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type WoodenShield struct {
	id int64
}

func (shield *WoodenShield) Id() int64 {
	return shield.id
}

func (shield *WoodenShield) Name() string {
	return "Wooden Shield"
}

func (shield *WoodenShield) Description() string {
	return "A small wooden shield"
}

func (shield *WoodenShield) GoldValue() int64 {
	return 35
}

func (shield *WoodenShield) Class() int64 {
	return 3
}

func (shield *WoodenShield) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_HAND_HELD
}

type SpikedSheild struct {
	id int64
}

func (shield *SpikedSheild) Id() int64 {
	return shield.id
}

func (shield *SpikedSheild) Name() string {
	return "Spiked Sheild"
}

func (shield *SpikedSheild) Description() string {
	return "ouch"
}

func (shield *SpikedSheild) GoldValue() int64 {
	return 45
}

func (shield *SpikedSheild) Class() int64 {
	return 5
}

func (shield *SpikedSheild) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_HAND_HELD
}
