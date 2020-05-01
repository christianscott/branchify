# branchify

Command line utility for quickly creating branch names.

## Usage

```
Usage of branchify:
  -ns string
        prefixed to the branch name to namespace it
  -nsSep string
        seperator used to join the namespace and the branch name (default "@")
  -sep string
        seperator used to join branch name parts (default "-")
```

## Examples

Simple usage

```bash
$ branchify foo bar baz
foo-bar-baz
```

Using all the options

```bash
$ branchify -ns me -nsSep ! -sep . foo bar baz
me!foo.bar.baz
```

Stripping out illegal chars (won't catch everything)

```bash
$ branchify foo:bar:baz
foobarbaz
```

As part of a git alias

```bash
$ cat ~/.gitconfig
[alias]
  cbc="!f() { local parts=$@; local branch=$(branchify -ns christianscott \"${parts[@]}\"); git checkout -b $branch; }; f"
$ git cbc foo bar baz
christianscott@foo-bar-baz
```

## Installation

Installable via `go get`

```bash
$ go get christianscott/branchify
```
