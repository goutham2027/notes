class Account(object):
    num_accounts = 0

    def __init__(self, name, balance):
        self.name = name
        self.balance = balance
        Account.num_accounts += 1

    def del_account(self):
        Account.num_accounts -= 1

    def __getattr__(self, name):
        return "No attribute: {}".format(name)



a1 = Account('gp', 1000)

print a1.get_balance

