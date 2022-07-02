import sys
sys.path.append("../dungeonexplorer")
from dungeonexplorer.character import *
import unittest

class TestEquppedGear(unittest.TestCase):
    def test_init(self):
        #self.assertEqual
        helmet = Armor()
        gear = EquippedGear()
        gear.equipArmor(helmet, ArmorSlot.HEAD)
        print(gear)


if __name__ == '__main__':
    unittest.main()