class Account(object):
    def __init__(self):
        self._acct_num = None

    def get_acct_num(self):
        return self._acct_num

    def set_acct_num(self, value):
        self._acct_num = value

    def del_acct_num(self):
        del self._acct_num


    acct_num = property(get_acct_num, set_acct_num, del_acct_num, "Account number property")
