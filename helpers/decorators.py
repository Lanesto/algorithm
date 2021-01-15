#!/usr/bin/python3
import time
import functools
import signal
import typing as T


# Custom error class represents function timed out
class TimeoutError(Exception):
    pass


def clock(
    func: T.Callable = None,
    *,
    fmt: str = '[{elapsed:0.8f}] {fname}({params:.40}) → {result:.80}'
) -> T.Callable:
    """
    Decorator factory measuring function runtime.

    As a way of backward compatibility, because once it is used as a simple
    decorator, it acts like just a decorator if a non-keyword argument is
    given.

    All controls should be given as the keyword arguments due to the reason
    explained above.

    Old usage:
    >>> @clock
    ... def a_legacy_usage(num, subject='legacy'):
    ...     print(f'Hello, {num} {subject}!')
    ...     return num
    >>> a_legacy_usage(1, subject='legacy')  # doctest: +ELLIPSIS
    Hello, 1 legacy!
    [...] a_legacy_usage(1, subject='legacy') → 1
    1

    New usage:
    >>> @clock(fmt='ran {fname}({params}) in {elapsed:0.3f} seconds: {result}')
    ... def a_new_usage(num, subject='new'):
    ...     print(f'Hello, {subject}!')
    ...     return num
    >>> a_new_usage(2)  # doctest: +ELLIPSIS
    Hello, new!
    ran a_new_usage(2) in ... seconds: 2
    2
    """
    def decorator(func: T.Callable) -> T.Callable:

        @functools.wraps(func)
        def wrapper(*args, **kwargs) -> T.Any:
            fname = func.__name__
            params = '{}'.format(
                ', '.join([repr(arg) for arg in args] +
                          [f'{k}={v!r}' for k, v in kwargs.items()])
            )
            t0 = time.perf_counter()
            result = func(*args, **kwargs)
            elapsed = time.perf_counter() - t0

            print(fmt.format_map({
                'fname': fname,
                'params': params,
                'elapsed': elapsed,
                'result': repr(result),
            }))
            return result

        return wrapper

    # For backward compatibility.
    # If used as a decorator, just act like that (manually decorates it).
    if func is not None:
        return decorator(func)

    return decorator


def timeout(seconds: float = 10) -> T.Callable:
    """
    Decorator factory used to limit function execution time.

    >>> @timeout(seconds=3)
    ... def lazy_friend():
    ...     from time import sleep
    ...     sleep(5)
    >>> try:
    ...     lazy_friend()
    ... except TimeoutError as err:
    ...     print(err)
    function exceeded timeout (3s)
    """
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


# Run doctest by default. No output means no errors.
# Use -v option to see outputs in verbose form.
if __name__ == '__main__':
    import doctest
    doctest.testmod()
