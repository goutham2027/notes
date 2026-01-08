from decimal import Decimal, getcontext

def calculate_pi(digits):
    """
    Calculate pi to the specified number of decimal digits.
    Uses the Machin formula: pi/4 = 4*arctan(1/5) - arctan(1/239)
    
    Args:
        digits: Number of decimal places to calculate
    
    Returns:
        Decimal: Pi value with the specified precision
    """
    # Set precision higher than needed to avoid rounding errors
    getcontext().prec = digits + 10
    
    def arctan(x, num_terms):
        """Calculate arctan(x) using Taylor series"""
        getcontext().prec = digits + 10
        x = Decimal(x)
        power = x
        result = power
        for n in range(1, num_terms):
            power *= -x * x
            result += power / (2 * n + 1)
        return result
    
    # Machin formula: pi = 16*arctan(1/5) - 4*arctan(1/239)
    # Number of terms for convergence
    num_terms = digits + 10
    
    pi = 4 * (4 * arctan(Decimal(1) / Decimal(5), num_terms) - 
              arctan(Decimal(1) / Decimal(239), num_terms))
    
    # Set precision to requested digits
    getcontext().prec = digits + 1
    return +pi  # The unary plus rounds to current precision

def get_pi_5_digits():
    """
    Calculate pi to the 5th decimal digit.
    Returns pi as 3.14159
    
    Returns:
        Decimal: Pi rounded to 5 decimal places
    """
    pi = calculate_pi(5)
    return pi

if __name__ == "__main__":
    pi_value = get_pi_5_digits()
    print(f"Pi to 5 digits: {pi_value}")
