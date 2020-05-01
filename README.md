# branchify

Attempts to create a valid branch name from arguments.

# Usage

```
Usage of branchify:
  -ns string
        prefixed to the branch name to namespace it
  -nsSep string
        seperator used to join the namespace and the branch name (default "@")
  -sep string
        seperator used to join branch name parts (default "-")
```

Uses args to create a nice branch name. Can be namespaced for uniqueness

```bash
$ branchify -ns me -nsSep @ -sep - foo bar baz
me#foo.bar.baz
```

It will strip out some illegal chars

```bash
$ branchify foo:bar:baz
foobarbaz
```

I use it as part of a git alias to quickly create branches namespaced with my github username

```ini
[alias]
  cbc="!f() { local parts=$@; local branch=$(branchify -ns christianscott \"${parts[@]}\"); git checkout -b $branch; }; f"
```

You can install it using `go get`

```bash
$ go get christianscott/branchify
```
