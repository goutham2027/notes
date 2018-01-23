class TypedAttribute:
    def __init__(self, name, type, default=None):
        self.name = "_" + name
        self.type = type
        self.default = default if default else type()

    def __get__(self, instance, cls):
        return getattr(instance, self.name, self.default)

    def __set__(self, instance, value):
        if not isinstance(value, self.type):
            raise TypeError("Must be a %s " % self.type)
        setattr(instance, self.name, value)

    def __delete__(self, instance):
        raise AttributeError("Can't delete attribtute")


class Account:
    name = TypedAttribute("name", str)
    balance = TypedAttribute("balance", int, 42)


acct = Account()
acct.name = "Goutham"
acct.balance = 1234

print(acct.balance)
print(acct.name)

# should throw an error
acct.balance = "1234"
print(acct.balance)
