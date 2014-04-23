#!/usr/bin/env python
# envcrypt

import collections
import os
import subprocess
import sys


def main(argv):
    path = argv[1]
    command = argv[2:]

    data = decrypt(path)
    env = dict(parse(data))
    os.execvpe(command[0], command, env)


def decrypt(path):
    return subprocess.check_output(["gpg", "-q", "--batch", "-d", path])


def parse(envdata, sep="="):
    envlines = envdata.splitlines()

    for line in envlines:
        key, value = line.split(sep, 1)
        yield Variable(key, value)


def run():
    return main(sys.argv)

Variable = collections.namedtuple("Variable", "key value")

if __name__ == "__main__":
    sys.exit(run())
