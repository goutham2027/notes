def make_counter():
    count = 0
    def counter():
        nonlocal count

        count += 1
        return count
    return counter

x = make_counter()
print(x)
print(x())
print(x())
print(x)
import ipdb; ipdb.set_trace()
