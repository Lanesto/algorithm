#!/usr/bin/python3
import cProfile
from io import StringIO
import traceback
import typing as T
from .decorators import clock, timeout  # type: ignore


def is_equal(a: T.Any, b: T.Any) -> bool:
    return a == b


defaults = {
    # Input for solution
    'input': [],

    # Output expected
    'expected': [],

    # Function to compare expected and actual output
    # If None check with == operation
    'oracle': is_equal,

    # Limit of function execution time
    'timeout': 10,

    # Print verbosity; run profiler if given testcase suceeed
    'verbose': False,
}

supported_keys = defaults.keys()


def _get_unsupported_keys(a: dict, b: dict) -> T.Set:
    """
    Returns the keys that is in dict b but not in dict a

    >>> a_test_case = {
    ...     'input': [1, 2],
    ...     'expected': 55,
    ...     '_unsupported_key': 'error'
    ... }
    >>> _get_unsupported_keys(defaults, a_test_case)
    {'_unsupported_key'}
    """
    return b.keys() - a.keys() & b.keys()


def run(func: T.Callable, test_cases: T.Dict[str, T.Any]):
    """
    Run :func with :test_cases; Will not be evaluated if element of :test_cases
    includes any key not specified in :defaults.

    See supported options and its explanations at :defaults.

    >>> def solution(x, y):
    ...     return x + y
    >>> test_cases = [
    ...     {
    ...         'input': [1, 2],
    ...         'expected': 3,
    ...         'verbose': True,
    ...     },
    ...     {
    ...         'input': [-5, 7],
    ...         'expected': 2,
    ...     },
    ...     {
    ...         'input': [2, 5],
    ...         'expected': 6,
    ...     }
    ... ]
    >>> from io import StringIO
    >>> from contextlib import redirect_stdout
    >>> with StringIO() as buf, redirect_stdout(buf):
    ...     run(solution, test_cases)
    ...     content = buf.getvalue()
    >>> print(len(content))
    842
    """

    # For profiler statement
    old_func = func

    # Timeout is decorated for each test cases to support different timeouts
    func = clock()(func)

    # List of index of passed tests
    passed_cases: list[int] = []

    # List of index of failed tests
    failed_cases: list[int] = []

    # Run tests
    for i, tc in enumerate(test_cases, start=1):
        # Check all keys are supported
        tc_keys = tc.keys()
        unsupported_keys = _get_unsupported_keys(defaults, tc)
        if unsupported_keys:
            message = 'unsupported keys: {}'.format(', '.join(unsupported_keys))  # noqa: E501
            print(f'Case {i} FAIL: {message}')
            continue

        # Merge :tc over copy of :defaults
        temp = defaults.copy()
        temp.update(tc)
        tc = temp

        # Run test
        t_func = timeout(seconds=tc['timeout'])(func)
        input, expected = tc['input'], tc['expected']
        try:
            oracle = tc.get('oracle')
            assert callable(oracle), \
                f'oracle is not callable: {oracle!r}'

            result = t_func(*input)
            assert oracle(result, expected), \
                f'wants {expected!r:.40} but got {result!r:.40}'
        except Exception as e:
            failed_cases.append(i)
            print(f'Case {i} FAIL: {e!s}')
            with StringIO() as buf:
                traceback.print_exc(limit=5, file=buf)
                indented = '\n'.join(
                    '    ' + line for line in buf.getvalue().split('\n')
                )
                print()
                print(indented)

        else:
            passed_cases.append(i)
            print(f'Case {i} PASS')
            if tc['verbose']:
                print()
                cProfile.runctx(
                    'old_func(*input)',
                    globals=globals(),
                    locals=locals()
                )

        print()

    print()

    # Collect statistics on results
    total_passed = len(passed_cases)
    total_failed = len(failed_cases)
    total = len(test_cases)
    if total_passed == total:
        print(f'All {len(test_cases)} test(s) have passed')
    else:
        fmt = '{} out of {} test(s) failed: {}'
        print(fmt.format(
            total_failed,
            total,
            ', '.join(str(fc) for fc in failed_cases)
        ))


if __name__ == '__main__':
    import doctest
    doctest.testmod()
