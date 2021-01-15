import typing as T


# Truncate given item's reprentation if it is longer than given 'n'
def truncate(obj: T.Any, n: int) -> str:
    """
    Truncate repr(:obj) with ellipsis(...) if size exceeds :n

    >>> truncate('a' * 30, 10)
    'aaaaaaaaaa...'
    >>> truncate(2**30, 5)
    '10737...'
    """
    s = repr(obj) if not isinstance(obj, str) else obj
    if n < 0:
        return s

    return s[:n] + '...' if len(s) > n else s


# Run doctest by default. No output means no errors.
# Use -v option to see outputs in verbose form.
if __name__ == '__main__':
    import doctest
    doctest.testmod()
