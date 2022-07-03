package gameengine

import (
	"testing"

	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
	"github.com/stretchr/testify/require"
)

func TestWoodenShield_DefaultsOk(t *testing.T) {
	shield := WoodenShield{id: 1}
	itemInterfaceTest(&shield)
	require.Less(t, shield.Class(), int64(100))
	require.Equal(t, Code.ArmorType_ArmorType_HAND_HELD, shield.Type())
}
