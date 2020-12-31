# helper scripts for algorithm debugging
import time
import functools
import signal
import typing as T

DEBUG = True


# Custom error class represents function timed out
class TimeoutError(Exception):
    pass


# Wrapper for print only executed when DEBUG is set True
def debug_msg(*args, **kwargs) -> None:
    if DEBUG:
        print(*args, **kwargs)


# Decorator to time function speed
def clock(func: T.Callable) -> T.Callable:
    @functools.wraps(func)
    def wrapper(*args, **kwargs) -> T.Any:
        params = '{}({})'.format(
            func.__name__,
            ', '.join([repr(arg) for arg in args] +
                      [f'{k}={v}' for k, v in kwargs.items()])
        )
        params = truncate(params, 120)
        t0 = time.perf_counter()
        result = func(*args, **kwargs)
        elapsed = time.perf_counter() - t0
        debug_msg('[{:0.8f}] {} → {}'.format(
            elapsed, params, truncate(repr(result), 60)))
        return result

    return wrapper


# Decorator to limit function running time
# Returns None if timeout occurred
def timeout(seconds: float = 10) -> T.Callable:
    def decorator(func: T.Callable) -> T.Any:
        def handler(signum, frame):
            fmt = 'function exceeded timeout ({}s)'
            raise TimeoutError(fmt.format(seconds))

        @functools.wraps(func)
        def wrapper(*args, **kwargs):
            signal.signal(signal.SIGALRM, handler)
            signal.setitimer(signal.ITIMER_REAL, seconds)
            try:
                result = func(*args, **kwargs)
            finally:
                signal.alarm(0)

            return result

        return wrapper

    return decorator


# Truncate given item's reprentation if it is longer than given 'n'
def truncate(obj: T.Any, n: int) -> str:
    s = repr(obj)
    if n < 0:
        return s

    return s[:n] + '...' if len(s) > n else s
