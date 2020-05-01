# branchify

Attempts to create a valid branch name from arguments.

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
