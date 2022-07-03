package gameengine

import (
	"testing"

	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
	"github.com/stretchr/testify/require"
)

func TestClothBoots_DefaultsOk(t *testing.T) {
	boots := ClothBoots{id: 1}
	itemInterfaceTest(&boots)
	require.Less(t, boots.Class(), int64(100))
	require.Equal(t, Code.ArmorType_ArmorType_BOOTS, boots.Type())
}
