package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type RubyRing struct {
	id int64
}

func (ring *RubyRing) Id() int64 {
	return ring.id
}

func (ring *RubyRing) Name() string {
	return "Ruby Ring"
}

func (ring *RubyRing) Description() string {
	return "a priceless gem for a priceless ring "
}

func (ring *RubyRing) GoldValue() int64 {
	return 10
}

func (ring *RubyRing) Class() int64 {
	return 1
}

func (ring *RubyRing) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_FINGERS
}

type diamondRing struct {
	id int64
}

func (ring *diamondRing) Id() int64 {
	return ring.id
}

func (ring *diamondRing) Name() string {
	return "Diamond Ring"
}

func (ring *diamondRing) Description() string {
	return "a priceless gem for a priceless ring "
}

func (ring *diamondRing) GoldValue() int64 {
	return 10
}

func (ring *diamondRing) Class() int64 {
	return 1
}

func (ring *diamondRing) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_FINGERS
}

type sapphireRing struct {
	id int64
}

func (ring *sapphireRing) Id() int64 {
	return ring.id
}

func (ring *sapphireRing) Name() string {
	return "Sapphire Ring"
}

func (ring *sapphireRing) Description() string {
	return "a priceless gem for a priceless ring "
}

func (ring *sapphireRing) GoldValue() int64 {
	return 10
}

func (ring *sapphireRing) Class() int64 {
	return 1
}

func (ring *sapphireRing) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_FINGERS
}

type TopazRing struct {
	id int64
}

func (ring *TopazRing) Id() int64 {
	return ring.id
}

func (ring *TopazRing) Name() string {
	return "Topaz Ring"
}

func (ring *TopazRing) Description() string {
	return "a priceless gem for a priceless ring "
}
func (ring *TopazRing) GoldValue() int64 {
	return 10
}

func (ring *TopazRing) Class() int64 {
	return 1
}

func (ring *TopazRing) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_FINGERS
}

type EmeraldRing struct {
	id int64
}

func (ring *EmeraldRing) Id() int64 {
	return ring.id
}

func (ring *EmeraldRing) Name() string {
	return "Emerald Ring"
}

func (ring *EmeraldRing) Description() string {
	return "a priceless gem for a priceless ring "
}
func (ring *EmeraldRing) GoldValue() int64 {
	return 10
}

func (ring *EmeraldRing) Class() int64 {
	return 1
}

func (ring *EmeraldRing) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_FINGERS
}
