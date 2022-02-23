# slice: text sampler

slice samples random lines from STDIN.

# ABOUT

slice extracts random lines from text. This is useful for a variety of applications.

* Statistics
* Name generation
* File previewing

For example, `head`/`tail` only show the very start and end of a document. Where slice shows a more representative interlace of the overall content. Like `less`/`more`, but in a compact, intentionally lossy form.

# EXAMPLES

```console
$ cd examples

$ slice romeo-and-juliet.txt
  Escalus, Prince of Verona.
  Friar John, Franciscan.
  Three Musicians.
  An Officer.
  Nurse to Juliet.
    In fair Verona, where we lay our scene,
                                                         [Exit.]

$ slice -rate 0.05 constellations.txt
Ara
Bootes
Canis
```

By default, `slice` reads from STDIN. See `slice -help` for more information.

# DOWNLOAD

https://github.com/mcandre/slice/releases

# DOCUMENTATION

https://pkg.go.dev/github.com/mcandre/slice

# CONTRIBUTING

See [DEVELOPMENT.md](DEVELOPMENT.md).

# LICENSE

FreeBSD

# SEE ALSO

* [awk](https://en.wikipedia.org/wiki/AWK), a complex line processor
* [head](https://linux.die.net/man/1/head)/[tail](https://linux.die.net/man/1/tail), basic text truncators
* [less](https://linux.die.net/man/1/less), an interactive text file reader
* [more](https://en.wikipedia.org/wiki/More_(command)), a limited interactive text file reader
* [perl](https://www.perl.org/), a very complex text processor
* [sed](https://en.wikipedia.org/wiki/Sed), a simple line processor
* [uniq](https://linux.die.net/man/1/uniq), a text filter for uniqueness

🔪
