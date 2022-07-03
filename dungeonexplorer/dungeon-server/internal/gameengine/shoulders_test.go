package gameengine

import (
	"testing"

	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
	"github.com/stretchr/testify/require"
)

func TestClothShoulders_DefaultsOk(t *testing.T) {
	shoulders := ClothShoulders{id: 1}
	itemInterfaceTest(&shoulders)
	require.Less(t, shoulders.Class(), int64(100))
	require.Equal(t, Code.ArmorType_ArmorType_SHOULDERS, shoulders.Type())
}
