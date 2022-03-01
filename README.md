# zample: streaming line filter

# EXAMPLE

```console
$ cd examples

$ zample romeo-and-juliet.txt
  Escalus, Prince of Verona.
  Friar John, Franciscan.
  Three Musicians.
  An Officer.
  Nurse to Juliet.
    In fair Verona, where we lay our scene,
                                                         [Exit.]
```

# ABOUT

`zample` selects random lines from text inputs. This kind of filter has several uses:

* Statistics
* Random name generators
* Text processing
* Telemetry downsampling
* File previews

For example, `head`/`tail` may show the start and end of a document. Whereas `zample` shows a more representative view of the overall document body. In this way, `zample` behaves akin to `less`/`more`, but in a compact, lossy form.

# LICENSE

FreeBSD

# DOWNLOAD

https://github.com/mcandre/zample/releases

# API DOCUMENTATION

https://pkg.go.dev/github.com/mcandre/zample

# USAGE

By default, the selection rate of each text line is a `0.10` (10%) chance per line. That is, about 10% of text lines may become output, with the remaining 90% being stripped away from the output.

The sampling rate can be customized with a `-rate <value>` flag, using values in the range `[0.0, ... 1.0]`. For example, to select `0.05` (5%) of stellar constellations:

```console
$ zample -rate 0.05 constellations.txt
Aquarius
Pisces
```

`zample` supports multiple file paths.

```console
$ zample constellations.txt cities.txt colors.txt
Pisces
Yokohama
Red
```

`zample` is optimized for large data sets, and does not support robust entry reordering. Any apparent reorderding is an accidental artifact. This is a consequence of how `zample` optimizes for large data sets, for example by deciding each text line catch/release chance in a streaming fashion. Each chance resolves independently of the other. For deliberate shuffling of your data, you may pipe `zample` to additional tools such as `shuf`.

Likewise, different runs of the same `zample` experiments may produce different result sizes. For best effect, generate more input data, increase the selection rate, or try the `-skip` option.

`-skip <n>` deterministically skips every nth text line. This disables probabalistic rate behavior.

```console
$ zample -skip 2 cities.txt
Amsterdam
Casablanca
Edison
Gallipoli
Italia
Kilogramme
Madagascar
Oslo
Quebec
Santiago
Upsala
Washington
Yokohama
```

By default, `zample` reads from stdin.

See `zample -help` for more information.

# CONTRIBUTING

See [DEVELOPMENT.md](DEVELOPMENT.md).

# RESOURCES

* [awk](https://en.wikipedia.org/wiki/AWK), a complex line processor
* [head](https://linux.die.net/man/1/head)/[tail](https://linux.die.net/man/1/tail), basic text truncators
* [hellcat](https://github.com/mcandre/hellcat), a portable hex dumper
* [less](https://linux.die.net/man/1/less), an interactive text file reader
* [more](https://en.wikipedia.org/wiki/More_(command)), a limited interactive text file reader
* [od](https://linux.die.net/man/1/od), a classic hex dumper
* [perl](https://www.perl.org/), a very complex text processor
* [sed](https://en.wikipedia.org/wiki/Sed), a simple line processor
* [shuf](https://linux.die.net/man/1/shuf), a line shuffler
* [uniq](https://linux.die.net/man/1/uniq), a text filter for uniqueness
* [wc](https://linux.die.net/man/1/wc) provides basic text file metrics

🧪
