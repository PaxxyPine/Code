from enum import Enum

class Elemnts(Enum):
    FIRE = 1,
    WIND = 2,
    WATER = 3,
    WOOD = 4

class ArmorType(Enum):
    HELM = 1,
    SHOULDERS = 2,
    CHEST = 3,
    GOVES = 4,
    BELT = 5,
    LEGS = 6,
    BOOTS = 7,
    HELD = 8,
    NECK = 9,
    FINGERS = 10

class Item:
    def __init__(self, name, description, value, weight):
        self.name = name
        self.description = description
        self.value = value
        self.weight = weight

class Weapon(Item):
    def __init__(self, name, description, value, weight):
        super().__init__(name, description, value, weight)
        self.damage = 0
        self.element = Elemnts.FIRE
        self.range = 0

class Armor(Item):
    def __init__(self, name, description, value, weight):
        super().__init__(name, description, value, weight)
        self.armor = 0
        self.element = Elemnts.FIRE
        self.type = ArmorType.BOOTS