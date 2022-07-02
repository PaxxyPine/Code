from character import *
import unittest

class TestEquppedGear(unittest.TestCase):
    def test_init(self):
        #self.assertEqual
        helmet = Armor("TestArmor", "Just a test helmet", 100, 10)
        gear = EquippedGear()
        gear.equipArmor(helmet, ArmorSlot.HEAD)
        print(gear)


if __name__ == '__main__':
    unittest.main()