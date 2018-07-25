class C(object):
    def __init__(self):
        self._x = None

    @property
    def x(self):
        return self._x

    @x.setter
    def x(self, value):
        self._x = value

    @x.deleter
    def x(self):
        del self._x

c1 = C()
print(c1.x)
c1.x = 20
print(c1.x)
del c1.x
print(c1.x)
