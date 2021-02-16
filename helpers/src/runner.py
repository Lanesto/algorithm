#!/usr/bin/env python3
"""
Module provides running entry with parsed data and reporting differences

"""
import sys
import io
import re
import typing as T
from pathlib import Path
from contextlib import redirect_stdout
from difflib import unified_diff


def _find_data(path: str = ".") -> str:
    """
    Return path for data file(.dat) in dir for input to run test cases.

    It first finds with extension and name of current working directory,
    for example, with directory structure like:

    /workspace
        /algospot
            /ZIMBABWE
                ZIMBABWE.ipynb
                ZIMBABWE.dat

    On 'ZIMBABWE.ipynb', it first tries to find 'ZIMBABWE.dat'. If not found,
    it will try to find any file with extension .dat

    If none of data file exists, it will emit FileNotFoundError.

    Doctest:

        >>> _find_data('template')
        '/workspace/helpers/template/PROBLEM.dat'
        >>> _find_data()
        Traceback (most recent call last):
            ...
        FileNotFoundError: '*.dat' file not found

    """
    # Look for dir.dat first
    dir = Path(path)
    p = dir / f"{dir.name}.dat"
    ret: Path = None
    if p.exists() and p.is_file():
        ret = p
    else:
        # Look for *.dat
        for child in dir.iterdir():
            if child.is_file() and child.suffix.lower() == ".dat":
                ret = child
                break

    # Found any
    if ret:
        return ret.resolve().absolute().as_posix()

    # Not found
    raise FileNotFoundError("'*.dat' file not found")


_re_block = re.compile(r"^\[in\]\n?(.*)\n\[out\]\n?(.*)$", flags=re.DOTALL)


def _parse_data(body: str) -> T.Tuple[str, str]:
    """
    Parse data file and return [in], [out] block contents.

    Expected data content is like:

        [in]
        1
        2
        3
        [out]
        4
        5

    Doctest:

        >>> _parse_data('[in]\\n1\\n2\\n3\\n[out]\\n4\\n5')
        ('1\\n2\\n3', '4\\n5')

    """
    b_in, b_out = None, None
    m = _re_block.match(body)
    if m:
        b_in, b_out, *_ = m.groups()

    return (b_in, b_out)


# Type aliases
_readline = T.Callable[[T.Any], T.AnyStr]
_entry = T.Callable[[_readline], None]


def run(entry: _entry):
    """
    Run entry(readline) with given input stream and report differences between output and expectation.

    Doctest:

    Prepare .dat file

        >>> from tempfile import NamedTemporaryFile
        >>> def main(readline):
        ...     for line in iter(readline, ''):
        ...         print(line)

        >>> with NamedTemporaryFile(mode='w+t', dir='.', suffix='.dat') as fp:
        ...     _ = fp.write('[in]\\n1\\n2\\n3\\n[out]\\n4\\n5')
        ...     _ = fp.seek(0)
        ...     run(main)
        --- Expect
        +++ Got
        @@ -1,2 +1,5 @@
        -4
        -5
        +1
        +
        +2
        +
        +3

    """  # noqa: E501
    # Look for data file
    try:
        data_file = _find_data()
    except FileNotFoundError:
        raise

    # Parse data file
    with open(data_file, encoding="utf-8", mode="rt") as f:
        content = f.read()
    input, expected = _parse_data(content)

    # Run entry() and catch standard output
    with io.StringIO(input) as ifs, io.StringIO() as ofs, redirect_stdout(ofs):
        entry(ifs.readline)
        output = ofs.getvalue()

    # Report
    sys.stdout.writelines(
        unified_diff(
            (expected.strip() + "\n").splitlines(keepends=True),
            (output.strip() + "\n").splitlines(keepends=True),
            fromfile="Expect",
            tofile="Got",
        )
    )
