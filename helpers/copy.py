#!/usr/bin/env python3
"""Copy templates to target directory with placerholders replaced."""
import os
import sys
import shutil

_root_dir = os.path.dirname(__file__)
_template_dir = os.path.join(_root_dir, "template")


def _copy_with_name_format(format_map=None):
    def wrapper(src, dst, *args, **kwargs):
        dst = str(dst).format_map(format_map)
        print(f"coping {src} to {dst}")
        return shutil.copy2(src, dst, *args, **kwargs)

    return wrapper


if __name__ == "__main__":
    dest = sys.argv[1]
    ask = input(f"Copy templates to {dest}? (Y/n):")
    if ask.lower() != "y":
        exit()

    format_map = {"prob_id": os.path.basename(dest)}
    shutil.copytree(
        _template_dir,
        dest,
        copy_function=_copy_with_name_format(format_map=format_map),
    )
