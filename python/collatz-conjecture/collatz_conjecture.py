from loguru import logger

def steps(number: int) -> int:
    """
    Using the Collatz Conjecture, find out how many steps are needed to reach 1, given any number.

    Args:
        number: Initial number that started.

    Returns:
        The number of steps required to reach 1.
    """
    if number < 1:
        raise ValueError("Only positive integers are allowed")
    i = 0
    while number != 1:
        number = number/2 if number%2==0 else 3*number+1
        i+=1
    return i
