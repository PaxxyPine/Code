package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type RubyNeckless struct {
	id int64
}

func (amulet *RubyNeckless) Id() int64 {
	return amulet.id
}

func (amulet *RubyNeckless) Name() string {
	return "Ruby Neckless"
}

func (amulet *RubyNeckless) Description() string {
	return "a priceless gem for a priceless Neckless"
}

func (amulet *RubyNeckless) GoldValue() int64 {
	return 10
}

func (amulet *RubyNeckless) Class() int64 {
	return 1
}

func (amulet *RubyNeckless) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_NECK
}

type diamondNeckless struct {
	id int64
}

func (amulet *diamondNeckless) Id() int64 {
	return amulet.id
}

func (amulet *diamondNeckless) Name() string {
	return "Diamond Neckless"
}

func (amulet *diamondNeckless) Description() string {
	return "a priceless gem for a priceless Neckless "
}

func (amulet *diamondNeckless) GoldValue() int64 {
	return 10
}

func (amulet *diamondNeckless) Class() int64 {
	return 1
}

func (amulet *diamondNeckless) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_NECK
}

type sapphireNeckless struct {
	id int64
}

func (amulet *sapphireNeckless) Id() int64 {
	return amulet.id
}

func (amulet *sapphireNeckless) Name() string {
	return "Sapphire Neckless"
}

func (amulet *sapphireNeckless) Description() string {
	return "a priceless gem for a priceless Neckless "
}

func (amulet *sapphireNeckless) GoldValue() int64 {
	return 10
}

func (amulet *sapphireNeckless) Class() int64 {
	return 1
}

func (amulet *sapphireNeckless) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_NECK
}

type TopazNeckless struct {
	id int64
}

func (amulet *TopazNeckless) Id() int64 {
	return amulet.id
}

func (amulet *TopazNeckless) Name() string {
	return "Topaz Neckless"
}

func (amulet *TopazNeckless) Description() string {
	return "a priceless gem for a priceless Neckless "
}
func (amulet *TopazNeckless) GoldValue() int64 {
	return 10
}

func (amulet *TopazNeckless) Class() int64 {
	return 1
}

func (amulet *TopazNeckless) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_NECK
}

type EmeraldNeckless struct {
	id int64
}

func (amulet *EmeraldNeckless) Id() int64 {
	return amulet.id
}

func (amulet *EmeraldNeckless) Name() string {
	return "Emerald Neckless"
}

func (amulet *EmeraldNeckless) Description() string {
	return "a priceless gem for a priceless Neckless "
}
func (amulet *EmeraldNeckless) GoldValue() int64 {
	return 10
}

func (amulet *EmeraldNeckless) Class() int64 {
	return 1
}

func (amulet *EmeraldNeckless) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_NECK
}
