import unittest

from phonebook import Phonebook


class PhonebookTest(unittest.TestCase):

    def setUp(self):
        self.phonebook = Phonebook()

    # def tearDown(self):
        # pass

    def test_lookup_entry_by_name(self):
        self.phonebook.add("Bob", "12345")
        self.assertEqual("12345", self.phonebook.lookup("Bob"))

    def test_missing_entry_raises_KeyError(self):
        with self.assertRaises(KeyError):
            self.phonebook.lookup("missing")

    # @unittest.skip("WIP")
    def test_empty_phonebook_is_consistent(self):
        self.assertTrue(self.phonebook.is_consistent())

    # bad way of writing tests
    @unittest.skip("bad test")
    def test_is_consistent(self):
        self.assertTrue(self.phonebook.is_consistent())
        self.phonebook.add("Bob", "12345")
        self.assertTrue(self.phonebook.is_consistent())

        self.phonebook.add("Marry", "012345")
        self.assertTrue(self.phonebook.is_consistent())

        self.phonebook.add("Sue", "12345")  # identical to Bob
        self.assertFalse(self.phonebook.is_consistent())

        self.phonebook.add("Sue", "123")  # prefix to Bob
        self.assertFalse(self.phonebook.is_consistent())

    def test_phonebook_with_normal_entries_is_consistent(self):
        self.assertTrue(self.phonebook.is_consistent())
        self.phonebook.add("Bob", "12345")
        self.assertTrue(self.phonebook.is_consistent())

    def test_phonebook_with_duplicate_entries_is_inconsistent(self):
        self.phonebook.add("Bob", "12345")
        self.phonebook.add("Marry", "012345")
        self.assertFalse(self.phonebook.is_consistent())

    def test_phonebook_with_numbers_that_prefix_one_another_is_inconsistent(self):
        self.phonebook.add("Bob", "12345")
        self.phonebook.add("Sue", "123")  # prefix to Bob
        self.assertFalse(self.phonebook.is_consistent())

