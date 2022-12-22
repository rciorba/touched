# touched

touched is porcelain around git to get the list of modified files.

## Why

This tool lets you get the list of modified files from git.

Esenially it's porcelain around the following commands:

```
BASE="$(git merge-base origin/master HEAD)"
git diff $BASE...HEAD --name-only
git diff --name-only
```

## Installation

You can download a binary for your OS and CPU architecture from
https://github.com/rciorba/touched/releases


## Usage

Get list of files changed on your branch, which is based on the main branch:

```
touched origin/main
```

Get list of files changed on your branch, which is based on the main branch, but ignore
uncommited changes:

```
touched origin/main --iu
```


Get list of changes not yet commited:

```
touched
```

## License

Distributed under the terms of the `MIT`_ license, "touched" is free
and open source software

## Issues

If you encounter any problems, please `file an issue`_ along with a detailed description.
