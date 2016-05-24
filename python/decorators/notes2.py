# Python's Decorator Syntax
# Python makes creating and using decorators a bit cleaner and nicer for the
# programmer through some syntactic sugar.
# To decorate get_text we don't have to get_text = p_decorate(get_text)
# There is a neat shortcut for that, which is to mention the name of the
# decorating function before the function to be decorated.
# The name of the decorator should be prepended with an @ symbol.

def p_decorate(func):
    def func_wrapper(name):
        return "<p>{0}</p>".format(func(name))

    return func_wrapper

@p_decorate
def get_text(name):
    return "lorem ipsum, {0} dolor sit amet".format(name)

print get_text("Goutham")

# Now let's consider we wanted to decorate our get_text function by 2 other
# functions to wrap a div and strong tag around the string output.
def p_decorate(func):
    def func_wrapper(name):
        return "<p>{0}</p>".format(func(name))

    return func_wrapper

def strong_decorate(func):
    def func_wrapper(name):
        return "<strong>{0}</strong>".format(func(name))

    return func_wrapper

def div_decorate(func):
    def func_wrapper(name):
        return "<div>{0}</div>".format(func(name))

    return func_wrapper

# With the basic approach, decorating get_text would be along the lines of
# get_text = div_decorate(p_decorate(strong_decorate(get_text)))

# With Python's decorator syntax, same thing can be achieved with much more
# expressive power.

@div_decorate
@p_decorate
@strong_decorate
def get_text(name):
    return "lorem ipsum, {0} dolor sit amet".format(name)

print get_text("Goutham")
