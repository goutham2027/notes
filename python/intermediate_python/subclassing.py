import math

class NextInteger(int):
    def __new__(cls, val):
        return int.__new__(cls, math.ceil(val))

print(NextInteger(5.2))

