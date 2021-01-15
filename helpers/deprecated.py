import warnings


def debug_msg(*args, **kwargs) -> None:
    """
    DEPRECATED: Just run print for given input
    """
    warnings.warn('debug_msg is deprecated; use print()', DeprecationWarning)
    return print(*args, **kwargs)
