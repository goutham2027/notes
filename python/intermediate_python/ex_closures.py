class Log:
    def __init__(self, level):
        self._level = level

    def __call__(self, message):
        print("{}: {}".format(self._level, message))


log_info = Log("info")
log_error = Log("error")

# the same functionality can be implemented with function closures


def make_log(level):
    def _(message):
        print("{}: {}".format(level, message))
    return _

log_info = make_log('info')
log_error = make_log('error')
