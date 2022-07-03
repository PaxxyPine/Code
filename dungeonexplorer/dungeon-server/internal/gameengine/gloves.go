package gameengine

import (
	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type ClothGloves struct {
	id int64
}

func (helm *ClothGloves) Name() string {
	return "Cloth Gloves"
}

func (helm *ClothGloves) Description() string {
	return "A raggedy cloth pair of gloves"
}

func (helm *ClothGloves) GoldValue() int64 {
	return 3
}

func (helm *ClothGloves) Class() int64 {
	return 1
}

func (helm *ClothGloves) Type() Code.ArmorType {
	return Code.ArmorType_ArmorType_GOVES
}
