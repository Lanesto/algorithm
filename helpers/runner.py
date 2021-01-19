#!/usr/bin/python3
"""
Provides functions to redirect stdin and stdout with files or io.StringIO
to mock ordinary I/Os repeatedly used to run algorithm test cases.
"""
import sys
import io
import re
import pathlib
import typing as T
from contextlib import redirect_stdout
from difflib import unified_diff

# Entry of test case
Readline = T.Callable[[T.Any], T.AnyStr]
Entry = T.Callable[[Readline], None]

# regex pattern for data file
# [in]
# ...
# [out]
# ...
input_block = re.compile(r'^\[in\]([\d\w\s]*)\[out\]')
output_block = re.compile(r'\[out\]([\s\S]*)$')


def _find_data(workdir: pathlib.Path = None) -> pathlib.Path:
    """
    Finds data file(.dat) in :cwd used as input to run test cases.
    It returns path of file in string.

    It first finds with extension and name of current working directory,
    for example, with directory structure like:

    /workspace
        /algospot
            /ZIMBABWE
                ZIMBABWE.ipynb
                ZIMBABWE.dat

    On ZIMBABWE.ipynb, it first tries to find 'ZIMBABWE.dat'. If not found,
    it will try to find any file with extension .dat

    If none of data file exists, it will emit FileNotFoundError.

    >>> from pathlib import Path
    >>> p = Path('/workspace/template')
    >>> _find_data(p).name  # look for template.dat
    'PROBLEM.dat'
    >>> _find_data()  # run for current directory: /workspace
    Traceback (most recent call last):
        ...
    FileNotFoundError: data file is not found
    """
    if workdir is None:
        workdir = pathlib.Path.cwd()

    p = workdir / f'{workdir.name}.dat'
    if p.exists() and p.is_file():
        return p
    else:
        for child in workdir.iterdir():
            if child.is_file() and child.suffix == '.dat':
                return child

    raise FileNotFoundError('data file is not found')


def _parse_data(content: str) -> T.Tuple[str, str]:
    """
    Parse data file and return in, out blocks.

    :fp should be path to file or pathlib.Path object for file.

    >>> from io import StringIO
    >>> _parse_data('[in]\\n1\\n2\\n3\\n[out]\\n4\\n5')
    ('1\\n2\\n3', '4\\n5')
    """
    input = input_block.findall(content)[0].strip()
    output = output_block.findall(content)[0].strip()

    return (input, output)


def run(entry: Entry):
    """
    Run entry(readline) with given input stream and compare
    output with expectation

    >>> from tempfile import NamedTemporaryFile
    >>> def main(readline):
    ...     for line in iter(readline, ''):
    ...         print(line)
    >>>
    >>> with NamedTemporaryFile(mode='w+t', dir='.', suffix='.dat') as fp:
    ...     _ = fp.write('[in]\\n1\\n2\\n3\\n[out]\\n4\\n5')
    ...     _ = fp.seek(0)
    ...     run(main)
    --- expected
    +++ output
    @@ -1,2 +1,5 @@
    -4
    -5
    +1
    +
    +2
    +
    +3
    >>>
    """
    data_file = _find_data()
    with data_file.open(mode='rt') as fp:
        content = fp.read()
        input, expected = _parse_data(content)

    ifs = io.StringIO(input)
    with io.StringIO() as ofs, redirect_stdout(ofs):
        entry(ifs.readline)
        output = ofs.getvalue()

    sys.stdout.writelines(
        unified_diff(
            (expected.strip() + '\n').splitlines(keepends=True),
            (output.strip() + '\n').splitlines(keepends=True),
            fromfile='expected',
            tofile='output',
        )
    )


if __name__ == '__main__':
    import doctest
    doctest.testmod()
