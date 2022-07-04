package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothHat struct {
	id int64
}

func (helm *ClothHat) Id() int64 {
	return helm.id
}

func (helm *ClothHat) Name() string {
	return "Cloth Hat"
}

func (helm *ClothHat) Description() string {
	return "A raggedy cloth hat"
}

func (helm *ClothHat) GoldValue() int64 {
	return 30
}

func (helm *ClothHat) Class() int64 {
	return 2
}

func (helm *ClothHat) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_HELM
}

type VikingHat struct {
	id int64
}

func (helm *VikingHat) Id() int64 {
	return helm.id
}

func (helm *VikingHat) Name() string {
	return "Viking Hat"
}

func (helm *VikingHat) Description() string {
	return "lost item from the past"
}

func (helm *VikingHat) GoldValue() int64 {
	return 50
}

func (helm *VikingHat) Class() int64 {
	return 6
}

func (helm *VikingHat) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_HELM
}

type UmbrellaHat struct {
	id int64
}

func (helm *UmbrellaHat) Id() int64 {
	return helm.id
}

func (helm *UmbrellaHat) Name() string {
	return "Umbrella Hat "
}

func (helm *UmbrellaHat) Description() string {
	return "keep the drip away also keeping your drip"
}

func (helm *UmbrellaHat) GoldValue() int64 {
	return 69
}

func (helm *UmbrellaHat) Class() int64 {
	return 3
}

func (helm *UmbrellaHat) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_HELM
}

type TopHat struct {
	id int64
}

func (helm *TopHat) Id() int64 {
	return helm.id
}

func (helm *TopHat) Name() string {
	return "Top Hat"
}

func (helm *TopHat) Description() string {
	return "FELXXX"
}

func (helm *TopHat) GoldValue() int64 {
	return 300
}

func (helm *TopHat) Class() int64 {
	return 0
}

func (helm *TopHat) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_HELM
}
