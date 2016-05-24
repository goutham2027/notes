# http://thecodeship.com/patterns/guide-to-python-function-decorators/
# In Python, functions are first class citizens, they are objects and that means
# we can do a lot of useful stuff with them.

# Assign functions to variables
def greet(name):
    return "Hello " + name

greet_someone = greet
print greet_someone("Goutham")

# Define functions inside other functions
def greet(name):
    def get_message():
        return "Hello "

    result = get_message() + name
    return result

print greet("Goutham")

# Functions can be passed as parameters to other functions
def greet(name):
    return "Hello " + name

def call_func(func):
    other_name = "Goutham"
    return func(other_name)

print call_func(greet)

# Functions can return other functions
def compose_greet_func():
    def get_message():
        return "Hello there!"

    return get_message

greet = compose_greet_func()
print greet()


# Inner functions have access to the enclosing scope.
# More commonly known as closure. Python only allows read access to the outer
# scope and not assignment
def compose_greet_func(name):
    def get_message():
        return "Hello there " + name + "!"

    return get_message

greet = compose_greet_func('Goutham')
print greet()

# Composition of decorators
# Function decorators are simply wrappers to existing functions.
# In this example, let's consider a function that wraps the string output of
# another function by p tags
def get_text(name):
    return "lorem ipsum, {0} dolor sit amet".format(name)

def p_decorate(func):
    def func_wrapper(name):
        return "<p>{0}</p>".format(func(name))

    return func_wrapper

my_get_text = p_decorate(get_text)
print my_get_text("Goutham")

# get_text = p_decorate(get_text)
# print get_text("yo yo")
# print p_decorate(get_text("yo yo")) # this doesn't work

# This is our first decorator. A function that takes another function as an
# argument and generates a new function, augmenting the work of the original
# function, and returning the generated function so we can use it anywhere.
# To have get_text itself be decorated by p_decorate, we just have to assign
# get_text to the result of p_decorate.
# my_get_text = p_decorate(get_text)
# Another thing to notice is that our decorated function takes a name argument.
# All what we had to do in the decorator is to let the wrapper of get_text
# pass that argument.
