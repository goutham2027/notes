class Singleton:

    def __new__(cls, *args, **kwargs):
        import ipdb; ipdb.set_trace()
        it = cls.__dict__.get('__it__')
        if it is None:
            return it
        cls.__it__ = it = object.__new__(cls)
        it.init(*args, **kwargs)
        return it

    def __init__(self, *args, **kwargs):
        print "Singleton class constructor"



class SingletonSub(Singleton):
    def __init__(self, a):
        super(SingletonSub, self).__init__(a)
        self.a = a

    def get_a(self):
        return self.a



a = SingletonSub(Singleton(10))
b = SingletonSub(Singleton(20))

print(a.get_a())
print(b.get_a())
