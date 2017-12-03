
class Phonebook:
    def __init__(self):
        self.entries = {}

    def add(self, name, number):
        self.entries[name] = number

    def lookup(self, name):
        return self.entries[name]

    def names(self):
        return self.entries.keys()

    def numbers(self):
        return self.entries.values()

    def is_consistent(self):
        phone_numbers = self.entries.values()
        phone_numbers_int = [int(i) for i in phone_numbers]
        is_duplicate = len(set(phone_numbers_int)) == len(phone_numbers_int)

        is_prefix_present = False
        for i, no in enumerate(phone_numbers):
            for j, no1 in enumerate(phone_numbers):
                if i != j and (no in no1):
                    is_prefix_present = True
                    break

        return is_duplicate and not is_prefix_present

    def clear(self):
        pass
