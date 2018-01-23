class Account(object):
    num_accounts = 0

    def __init__(self, name, balance):
        self.name = name
        self.balance = balance

        Account.num_accounts += 1

    def __call__(self, arg):
        return "I was called with {}".format(arg)

    def del_account(self):
        Account.num_accounts -= 1


acct = Account('gp', 1000)
print acct("Testing function call on instance object")
