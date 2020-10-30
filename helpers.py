# helper scripts for algorithm debugging
import typing
import time
import functools

DEBUG = True

# Wrapper for print only executed when DEBUG is set True
def debug_msg(*args: list, **kwargs: dict) -> None:
    if DEBUG:
        print(*args, **kwargs)

# Decorator to time function speed
def clock(func: typing.Callable) -> typing.Callable:
    @functools.wraps(func)
    def inner(*args, **kwargs):
        msg = "{}({})".format(
            func.__name__,
            ", ".join([repr(arg) for arg in args] + [f"{k}={v}" for k, v in kwargs.items()])
        )
        t0 = time.time()
        result = func(*args, **kwargs)
        elapsed = time.time() - t0
        debug_msg("[{:0.8f}] {} → {}".format(elapsed, msg, result))
        return result
        
    return inner

# Truncate given item's reprentation if it is longer than given 'n'
def truncate(obj: typing.Any, n: int) -> str:
    s = repr(obj)
    return s[:n] + "..." if len(s) > n else s
