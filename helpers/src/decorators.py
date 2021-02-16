#!/usr/bin/env python3
"""Module for helpful decorators for testing."""
import time
import functools
import signal
import typing as T


class TimeoutError(Exception):
    """Represents a error for timeout event."""

    pass


def clock(
    func: T.Callable = None,
    *,
    timer: T.Callable[[], float] = time.perf_counter,
    fmt: str = "[{elapsed:0.8f}] {fname}({params:.40}) → {result:.80}",
) -> T.Callable:
    """
    Return decorator measures function runtime.

    It works as decorator, not decorator factory if argument :func is None.

    Arguments:
    - timer is function called to measure time. It is called before and after the function.
    - fmt is output format with given keys: fname, params, elapsed, result

    Usage:
    - As decorator:

        >>> @clock
        ... def a_legacy_usage(num, subject='legacy'):
        ...     print(f'Hello, {num} {subject}!')
        ...     return num

        >>> a_legacy_usage(1, subject='legacy')  # doctest: +ELLIPSIS
        Hello, 1 legacy!
        [...] a_legacy_usage(1, subject='legacy') → 1
        1

    - As decorator factory:

        >>> @clock(fmt='ran {fname}({params}) in {elapsed:0.3f} seconds: {result}')
        ... def a_new_usage(num, subject='new'):
        ...     print(f'Hello, {subject}!')
        ...     return num

        >>> a_new_usage(2)  # doctest: +ELLIPSIS
        Hello, new!
        ran a_new_usage(2) in ... seconds: 2
        2

    """  # noqa: E501

    def decorator(func: T.Callable) -> T.Callable:
        @functools.wraps(func)
        def wrapper(*args, **kwargs) -> T.Any:
            # Reconstruct function call signature string
            fname = func.__name__
            params = "{}".format(
                ", ".join(
                    [repr(arg) for arg in args]
                    + [f"{k}={v!r}" for k, v in kwargs.items()]
                )
            )

            # Time function
            start = timer()
            result = func(*args, **kwargs)
            elapsed = timer() - start

            # Output
            print(
                fmt.format_map(
                    {
                        "fname": fname,
                        "params": params,
                        "elapsed": elapsed,
                        "result": repr(result),
                    }
                ),
            )
            return result

        return wrapper

    # Work as simple decorator
    if func is not None:
        return decorator(func)

    return decorator


def timeout(*, seconds: float = 10) -> T.Callable:
    """
    Return a decorator that raises TimeoutError when function running time exceeds given limit.

    Usage:

        >>> from time import sleep
        >>> @timeout(seconds=3)
        ... def lazy_friend():
        ...     sleep(5)

        >>> try:
        ...     lazy_friend()
        ... except TimeoutError as err:
        ...     print(err)
        function call exceeded timeout (3s)

    """  # noqa: E501

    def decorator(func: T.Callable) -> T.Callable:
        def handler(signum, frame):
            fmt = "function call exceeded timeout ({}s)"
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
