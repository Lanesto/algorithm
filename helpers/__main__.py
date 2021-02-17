#!/usr/bin/env python3
"""A Tool dispatcher script."""
import sys
import os

_root_dir = os.path.dirname(__file__)
_tool_mapper = {"copy": "copy.py", "test": "test.py"}

program = sys.argv[1]
args = sys.argv[2:]
for k, v in _tool_mapper.items():
    if program == k:
        command = "{} {}".format(os.path.join(_root_dir, v), " ".join(args))
        exit(os.system(command))
