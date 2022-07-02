from enum import Enum
from inventory import *

class ArmorSlot(Enum):
    HEAD = 1,
    SHOULDERS = 2,
    CHEST = 3,
    GLOVES = 4,
    WAIST = 5,
    LEGS = 6,
    FEET = 7,
    LEFT_HAND = 8,
    NECK = 9,
    LEFT_FINGERS = 10,
    RIGHT_FINGERS = 11

class EquippedGear:
    def __init__(self):
        self.armor = [Armor() for ii in ArmorSlot]
        self.weapon = Weapon()
    def getArmor(self, slot):
        return self.armor[slot]
    def getWeapon(self):
        return self.weapon
    def equipArmor(self, armor, slot):
        self.armor[slot] = armor
    def totalArmor(self):
        total = 0
        for slot in ArmorSlot:
            total += self.armor[slot].armor
        return total

class Player:
    def __init__(self, hp, armor, speed, cr, cd, atk, elemental):
        self.maxHP = hp
        self.hitPoints = hp
        self.armor = armor
        self.maxArmor = armor
        self.speed = speed
        self.critRate = cr
        self.critDamage = cd
        self.attack = atk
        self.elementalMastery = elemental
        self.equippedItems = {}
        self.inventory = {}

    def pickup(self, item):
        self.inventory.append(item)

    def drop(self, itemId):
        self.inventory.popitem(itemId)

   # def equipFromInventory(self, itemId):
        # 

class Monster:
    def __init__(self, name, description, hp, armor, speed, atk, elemental):
        self.name = name
        self.description = description
        self.maxHP = hp
        self.hitPoints = hp
        self.armor = armor
        self.maxArmor = armor
        self.speed = speed
        self.attack = atk
        self.elementalMastery = elemental