#!/usr/bin/env python
"""Usage: envcrypt PATH COMMAND...

Set environment variables defined in encrypted file PATH and run COMMAND.

Arguments:
    PATH        path to a gpg-encrypted file that can be read with eg
                `gpg -d PATH`
    COMMAND     command to be invoked in the context of the
                environment defined in PATH
"""

import collections
import os
import subprocess
import sys

from docopt import docopt


def main(argv):
    # Set options_first=True in case argv is eg:
    #  [ "/bin/sh", "-c", "echo sorry" ]
    opts = docopt(__doc__, argv=argv[1:], options_first=True)
    path = opts["PATH"]
    command = opts["COMMAND"]

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
