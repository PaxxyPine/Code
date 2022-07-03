package gameengine

import (
	"fmt"

	"github.com/PaxxyPine/Code/proto/gitub.com/PaxxyPine/Code"
)

type Stats struct {
	Dexterity        int64
	maxHP            int64
	currentHP        int64
	damage           int64
	speed            int64
	critRate         int64
	critDamage       int64
	elementalMastery Code.Elements
}

type EquippedGear struct {
	armor  []Armor
	weapon Weapon
}

func (gear *EquippedGear) GetArmor(slot Code.ArmorSlot) Armor {
	return gear.armor[slot]
}

func (gear *EquippedGear) TotalArmor() int64 {
	var total int64 = 0
	for _, a := range gear.armor {
		total += a.Class()
	}
	return total
}

func (gear *EquippedGear) GetWeapon() Weapon {
	return gear.weapon
}

func (gear *EquippedGear) EquipArmor(slot Code.ArmorSlot, armor Armor) error {
	switch slot {
	case Code.ArmorSlot_ArmorSlot_NOT_SET:
		return fmt.Errorf("unable to equip armor to NOT_SET slot")
	case Code.ArmorSlot_ArmorSlot_HEAD:
		if armor.Type() != Code.ArmorType_ArmorType_HELM {
			return fmt.Errorf("ony helmets may be worn on your head [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_SHOULDERS:
		if armor.Type() != Code.ArmorType_ArmorType_SHOULDERS {
			return fmt.Errorf("ony shoulders may be worn on your shoulders [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_CHEST:
		if armor.Type() != Code.ArmorType_ArmorType_CHEST {
			return fmt.Errorf("ony chest may be worn on your chest [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_GLOVES:
		if armor.Type() != Code.ArmorType_ArmorType_GOVES {
			return fmt.Errorf("ony gloves may be worn on your hands [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_WAIST:
		if armor.Type() != Code.ArmorType_ArmorType_BELT {
			return fmt.Errorf("ony belts may be worn on your waist [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_LEGS:
		if armor.Type() != Code.ArmorType_ArmorType_LEGS {
			return fmt.Errorf("ony legs may be worn on your legs [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_FEET:
		if armor.Type() != Code.ArmorType_ArmorType_BOOTS {
			return fmt.Errorf("ony boots may be worn on your feet [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_LEFT_HAND:
		if armor.Type() != Code.ArmorType_ArmorType_HAND_HELD {
			return fmt.Errorf("ony shields may be carried in your left hand [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_NECK:
		if armor.Type() != Code.ArmorType_ArmorType_NECK {
			return fmt.Errorf("ony amulets may be worn on your neck [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_LEFT_FINGERS:
		if armor.Type() != Code.ArmorType_ArmorType_FINGERS {
			return fmt.Errorf("ony rings may be worn on your fingers [%s]", armor.Type().String())
		}
	case Code.ArmorSlot_ArmorSlot_RIGHT_FINGERS:
		if armor.Type() != Code.ArmorType_ArmorType_FINGERS {
			return fmt.Errorf("ony rings may be worn on your fingers [%s]", armor.Type().String())
		}
	default:
		return fmt.Errorf("unknown armor type [%s]", armor.Type().String())
	}
	gear.armor[slot] = armor
	return nil
}

func (gear *EquippedGear) EquipWeapon(weapon Weapon) error {
	gear.weapon = weapon
	return nil
}

type Character struct {
	id        int64
	name      string
	gear      EquippedGear
	inventory map[int64]Item // map of item id to items limit 12
	stats     Stats
}

func (character *Character) EquipArmorFromInventory(armor Armor, slot Code.ArmorSlot) error {
	item := character.inventory[armor.Id()]
	if item == nil {
		return fmt.Errorf("cannot equip unknown item %d", armor.Id())
	}
	character.gear.EquipArmor(slot, armor)
	delete(character.inventory, item.Id())
	return nil
}
