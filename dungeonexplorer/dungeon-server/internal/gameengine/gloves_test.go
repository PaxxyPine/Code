package gameengine

import (
	"testing"

	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
	"github.com/stretchr/testify/require"
)

func TestClothGloves_DefaultsOk(t *testing.T) {
	gloves := ClothGloves{id: 1}
	itemInterfaceTest(&gloves)
	require.Less(t, gloves.Class(), int64(100))
	require.Equal(t, Code.ArmorType_ArmorType_GOVES, gloves.Type())
}
