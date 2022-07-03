package gameengine

import (
	"testing"

	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
	"github.com/stretchr/testify/require"
)

func TestClothChest_DefaultsOk(t *testing.T) {
	chest := ClothChest{id: 1}
	itemInterfaceTest(&chest)
	require.Less(t, chest.Class(), int64(100))
	require.Equal(t, Code.ArmorType_ArmorType_CHEST, chest.Type())
}
