import unittest
from decimal import Decimal
from main import calculate_pi, get_pi_5_digits


class TestPiCalculation(unittest.TestCase):
    """Test cases for pi calculation functions"""
    
    def test_pi_5_digits(self):
        """Test that pi to 5 digits is calculated correctly"""
        pi = get_pi_5_digits()
        # Pi to 5 decimal places should be 3.14159
        expected = Decimal('3.14159')
        self.assertEqual(pi, expected, 
                        f"Expected {expected}, got {pi}")
    
    def test_pi_value_range(self):
        """Test that calculated pi is in valid range"""
        pi = calculate_pi(10)
        # Pi should be between 3 and 3.15
        self.assertGreater(pi, Decimal('3.1'))
        self.assertLess(pi, Decimal('3.15'))
    
    def test_pi_approximation(self):
        """Test that pi is close to the known value"""
        pi = calculate_pi(10)
        # Known value of pi to 10 decimal places
        known_pi = Decimal('3.1415926536')
        # Check first 5 digits match
        self.assertEqual(str(pi)[:8], str(known_pi)[:8])
    
    def test_pi_first_digits(self):
        """Test the first few digits of pi"""
        pi = get_pi_5_digits()
        pi_str = str(pi)
        # Should start with 3.14159
        self.assertTrue(pi_str.startswith('3.14159'),
                       f"Pi string {pi_str} doesn't start with 3.14159")
    
    def test_calculate_pi_different_precisions(self):
        """Test calculate_pi with different precision levels"""
        pi_5 = calculate_pi(5)
        pi_10 = calculate_pi(10)
        
        # Both should start with 3.14159
        self.assertTrue(str(pi_5).startswith('3.14159'))
        self.assertTrue(str(pi_10).startswith('3.14159'))


if __name__ == '__main__':
    unittest.main()
