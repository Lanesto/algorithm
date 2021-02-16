#!/usr/bin/env python3
"""Module for testing for test cases."""
import sys
import cProfile
import pstats
import traceback
import textwrap
import operator
import typing as T
from collections import abc
from contextlib import contextmanager
from io import StringIO
from .decorators import clock, timeout

# Type alias for oracle functions.
Oracle = T.Callable[[T.Any, T.Any], bool]

# Enough small value for comparison between floating point numbers.
EPSILON = 1e-8


def epsilon(a: float, b: float) -> bool:
    """Oracle function for comparing floats."""
    return (b - a) < EPSILON


# TODO: inherit cProfile.Profile class and implement context manager.
@contextmanager
def _profile(
    prof: cProfile.Profile = None, *, outfile: T.TextIO = sys.stdout
) -> cProfile.Profile:
    """
    Implement context manager protocol for cProfile.Profile.

    If using python version 3.8 or higher, use cProfile.Profile to context manager directly
    """  # noqa: E501
    if prof is None:
        prof = cProfile.Profile()

    try:
        prof.enable()
        yield prof
    finally:
        prof.disable()
        sort_by = pstats.SortKey.CUMULATIVE
        ps = pstats.Stats(prof, stream=outfile).sort_stats(sort_by)
        ps.print_stats()


class TestResult(T.NamedTuple):
    """Test result data tuple."""

    ok: bool
    got: T.Any
    profile_report: str


class TestCase:
    """Define a single test case spec."""

    def __init__(
        self,
        input: T.Any = None,
        expected: T.Any = None,
        *,
        oracle: Oracle = operator.eq,
        timeout: float = 10,
        verbose: bool = False,
    ):
        """Initialize a TestCase instance."""
        # keyword arguments for input are currently not supported.
        self.input = input
        self.expected = expected

        # Oracle must be a callable with two parameters
        assert callable(oracle), "oracle must be a callable with two parameters"
        self.oracle = oracle

        # Timeout must be >= 1
        timeout = float(timeout)
        msg = ""
        if timeout < 0:
            msg = "invalid timeout value: {timeout!r}"
        elif timeout < 1:
            msg = "timeout is too short: {timeout!r}"

        if msg:
            raise ValueError(f"timeout ({timeout!r}) is too short")

        self.timeout = timeout

        # Report verbosity. If set, run profiler instead Normal function call.
        self.verbose = bool(verbose)

    def run(self, func: T.Callable) -> TestResult:
        """Run current test case for func."""
        # Apply timer
        func = clock(func)
        func = timeout(seconds=self.timeout)(func)

        # Call function
        got = None
        profile_report = ""
        if self.verbose:
            dump = StringIO()
            with _profile(outfile=dump) as _:
                got = func(*self.input)
            profile_report = dump.getvalue()
        else:
            got = func(*self.input)

        ok = self.oracle(got, self.expected)
        return TestResult(
            ok=ok,
            got=got,
            profile_report=profile_report,
        )


# TODO: move run() to TestRunner method
class TestRunner:
    """Run test(s) for given function and report results."""

    pass


def run(
    func: T.Callable,
    items: T.Iterable[T.Union[T.Dict[str, T.Any], TestCase]],
):
    """
    Run given func for test_cases.

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
        >>> len(content)
        2928


    """
    # Preprocess
    test_cases = []
    for item in items:
        tc: TestCase = None
        if isinstance(item, abc.Mapping):
            # Create TestCase instance using given mapping as arguments
            tc = TestCase(**item)
        elif not isinstance(item, TestCase):
            raise TypeError("test case must be mapping or TestCase instance")

        test_cases.append(tc)

    # Run all tests
    num_tests = len(test_cases)
    if num_tests == 0:
        print("no tests run.")
        return

    passed_tests = []
    failed_tests = []
    for i, tc in enumerate(test_cases, 1):
        try:
            result: TestResult = tc.run(func)
            assert result.ok, f"Expect {tc.expected} but got {result.got}"
        except Exception:
            failed_tests.append(i)
            print(f"Test {i} failed")
            print()
            with StringIO() as buf:
                traceback.print_exc(limit=5, file=buf)
                tb = buf.getvalue()

            print()
            print(textwrap.indent(tb, " " * 4))
            print()
        else:
            passed_tests.append(i)
            print(f"Test {i} passed")
            print()
            if tc.verbose:
                print()
                print(textwrap.indent(result.profile_report, " " * 4))

    # Report
    num_passed_tests = len(passed_tests)
    num_failed_tests = len(failed_tests)

    print()
    if num_passed_tests == 0:
        print(f"All {num_tests} tests failed.")
    elif num_failed_tests == 0:
        print(f"All {num_tests} tests passed.")
    else:
        print(
            f"Test {num_passed_tests} out of {num_tests} tests passed: "
            "failed {}".format(", ".join(map(str, failed_tests)))
        )
