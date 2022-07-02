from enum import Enum

class ArmorSlot(Enum):
    HEAD = 1,
    SHOULDERS = 2,
    CHEST = 3,
    GLOVES = 4,
    WAIST = 5,
    LEGS = 6,
    FEET = 7,
    LEFT_HAND = 8,
    RIGHT_HAND = 9,
    NECK = 10,
    LEFT_FINGERS = 11,
    RIGHT_FINGERS = 12

class EquippedGear:
    def __init__(self):
        self.gear = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    def get(self, slot):
        return self.gear[slot]
    #def totalArmor(self):
        #total = 0
        #for slot in ArmorSlot:
            #self.gear[slot].


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