

def small_straight(dice):
    """Score the given roll in the 'Small Straight' Yatzy category.

    Args:
        dice: a sorted list of 5 integers indicating the dice rolled
    Returns:
        an integer score

    >>> small_straight([1,2,3,4,5])
    15
    >>> small_straight([1,2,3,4,4])
    0

    This function only recognizes sorted lists or sets or other collections.
    >>> small_straight({1,2,3,4,5})
    15
    >>> small_straight([5,4,3,2,1])
    15
    """
    if sorted(dice) == [1, 2, 3, 4, 5]:
        return sum(dice)
    else:
        return 0


def full_house(dice):
    """Score the given roll in the 'Full House' category.

    >>> full_house([1,1,2,2,2])
    8
    >>> full_house([6,6,6,2,2])
    22
    >>> full_house([1,2,3,4,5])
    0
    >>> full_house([1,2,2,1,3])
    0
    """

    counts = {x: dice.count(x) for x in range(1, 7)}
    if 2 in counts.values() and 3 in counts.values():
        return sum(dice)
    return 0
