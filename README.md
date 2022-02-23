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

`slice` supports sampling multiple text files concurrently.

```console
$ slice constellations.txt cities.txt colors.txt
Ara
Yellow
Aries
Gallipoli
Cassiopeia
Coma
Washington
Zurich
Equuleus
Leo
Lupus
Phoenix
Vulpecula
```

`slice` does not feature robust reordering, either by line order or by file path order. Any apparent shuffling is a natural consequence of the input text supplied to `slice`, and any inter-threading timings experienced during processing of multiple text files.

If deliberate shuffling is desired, then slice output may be piped further into additional tools like `shuf`.

For small data sets, `slice` may produce very short output, or no output. This artifact diminishes as the rate and/or input line count grows. Rather, we have optimized the sampling algorithm to scale well over large data sets. The chance of preservation is evaluated once per line. In other words, different runs at the same slice rating may produce different sample output _line counts_, as well as different sample line contents. For best effect, grow your input data size, or try using the `-skip` option.

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
