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
