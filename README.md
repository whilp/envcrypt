# `envcrypt`

`envcrypt` runs commands in a special process environment loaded from encrypted files.

```
envcrypt aws.asc /bin/sh -c 'echo $AWS_ACCESS_KEY_ID'
XXXXXXXX
```

## Install

```
python setup.py install
```

## Test

```
python setup.py test
```
