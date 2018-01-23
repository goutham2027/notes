class Account(object):
    num_accounts = 0

    @staticmethod
    def test_staticmethod():
        return "Static method"

    def num_accounts(self):
        return self.num_accounts

    @classmethod
    def test_classmethod(cls):
        return "class method"

print Account.test_staticmethod
a = Account()
print a.num_accounts
print Account.test_classmethod
