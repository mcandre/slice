# slice: text sampler

slice samples random lines from your texts.

# EXAMPLE

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
```

# ABOUT

slice extracts random lines from text. This is useful for a variety of applications.

* Statistics
* Random name generator
* Text processing
* File previewing

For example, `head`/`tail` show only the very start and end of a document. Whereas slice shows a more representative interlace of the overall content. Akin to `less`/`more`, but in a compact, intentionally lossy form.

## Usage

By default, the preservation rate of each line is `0.1` (10%).

This probability can be customized with a `-rate` flag, as a value in `[0.0, 1.0]`. For example, to sample 5% of the stellar constellations:

```console
$ slice -rate 0.05 constellations.txt
Ara
Bootes
Canis
```

`slice` supports iterating over multiple text files.

```console
$ slice constellations.txt cities.txt colors.txt
Auriga
Bootes
Canis
Cepheus
Corona
Delphinus
Piscis
Sculptor
Telescopium
Triangulum
Amsterdam
Italia
Tripoli
Valencia
Orange
Blue
```

`slice` is not primarily a reordering tool. Neither for line reordering nor file path reordering. Any apparent shuffling is a natural consequence of the inputs. If deliberate shuffling is desired, then pipe slice with additional tools like `shuf`.

For small data sets, `slice` can produce very short output, or even no output. This artifact diminishes as the rate and/or input line count grows. In order to optimize the sampling algorithm for large data sets, we evaluate the chance of preservation once per line, at the time that line is processed. In other words, different runs at the same rating, may produce different sample output _line counts_, as well as different output contents. For best effect, generate more input data, or try the `-skip` option.

`slice` can deterministically skip every nth line of source text with a `-skip` flag. This disables probabalistic rate behavior.

```console
$ slice -skip 2 cities.txt
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

By default, `slice` reads from stdin.

See `slice -help` for more information.

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
* [shuf](https://linux.die.net/man/1/shuf), a line shuffler
* [uniq](https://linux.die.net/man/1/uniq), a text filter for uniqueness
* [wc](https://linux.die.net/man/1/wc) provides basic text file metrics

🔪
