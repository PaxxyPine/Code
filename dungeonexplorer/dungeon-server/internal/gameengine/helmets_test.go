package gameengine

import (
	"testing"

	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
	"github.com/stretchr/testify/require"
)

func itemInterfaceTest(Item) {
}

func TestClothHat_DefaultsOk(t *testing.T) {
	hat := ClothHat{id: 1}
	itemInterfaceTest(&hat)
	require.Less(t, hat.Class(), int64(100))
	require.Equal(t, Code.ArmorType_ArmorType_HELM, hat.Type())
}
