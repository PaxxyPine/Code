package gameengine

import (
	"testing"

	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
	"github.com/stretchr/testify/require"
)

func TestClothPants_DefaultsOk(t *testing.T) {
	pants := ClothPants{id: 1}
	itemInterfaceTest(&pants)
	require.Less(t, pants.Class(), int64(100))
	require.Equal(t, Code.ArmorType_ArmorType_LEGS, pants.Type())
}
