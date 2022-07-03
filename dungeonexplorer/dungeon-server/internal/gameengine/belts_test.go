package gameengine

import (
	"testing"

	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
	"github.com/stretchr/testify/require"
)

func TestClothBelt_DefaultsOk(t *testing.T) {
	belt := ClothBelt{id: 1}
	itemInterfaceTest(&belt)
	require.Less(t, belt.Class(), int64(100))
	require.Equal(t, Code.ArmorType_ArmorType_BELT, belt.Type())
}
