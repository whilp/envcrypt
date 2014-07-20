# `envcrypt`

`envcrypt` runs commands in a special process environment loaded from encrypted files.

```
$ envcrypt -h
usage: envcrypt PATH COMMAND [ARGS...]

Set environment variables defined in encrypted file PATH and run COMMAND.

Arguments:
  PATH     path to a gpg-encrypted file that can be read with eg `gpg -d PATH`
  COMMAND  command to be invoked in the context of the environment defined in PATH
```

The contents of `PATH` should be zero or more lines of string keys and values, separated by `=` and encrypted with GPG:

```
$ echo AWS_ACCESS_KEY_ID=XXXXXXX | gpg -a -e aws.asc
$ envcrypt aws.asc /bin/sh -c 'echo $AWS_ACCESS_KEY_ID'
XXXXXXXX
```

## Install

```
go get github.com/whilp/envcrypt
```

## Test

```
make test
```

## License

BSD; see `LICENSE`.
