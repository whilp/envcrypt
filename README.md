# `envcrypt`

`envcrypt` runs commands in a special process environment loaded from encrypted files.

```
echo AWS_ACCESS_KEY_ID=XXXXXXX | gpg -a -e aws.asc
envcrypt aws.asc /bin/sh -c 'echo $AWS_ACCESS_KEY_ID'
XXXXXXXX
```

## Install

```
pip install .
```

## Test

```
python setup.py test
```
