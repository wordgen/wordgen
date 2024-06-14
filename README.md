# wordgen

[![badge][badge-url]][pkg-url]

wordgen is a Go package that generates random words from a given wordlist. It
allows you to specify the number of words, their casing, and the separator
between them.

## Installation

To install the package, use `go get`:

```shell
go get github.com/wordgen/wordgen
```

## Usage

Import the package in your Go code:

```go
import "github.com/wordgen/wordgen"
```

Use the `Generator` struct and its `Generate` method to generate random words:

```go
package main

import (
    "fmt"
    "github.com/wordgen/wordgen"
)

func main() {
    wordlist := []string{"apple", "banana", "cherry", "date"}

    gen := wordgen.NewGenerator()
    gen.Words = wordlist
    gen.Count = 5
    gen.Casing = "title"
    gen.Separator = "-"

    result, err := gen.Generate()

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Generated words:", result)
}
```

## Wordlists Package and CLI Tool

The wordgen GitHub organization also includes a [wordlists package] containing
various wordlists converted to Go string slices, and a wordgen [command-line tool]
for generating random words.

You can use these components together to easily generate random words in your Go
applications or from the terminal.

## Contributing

When submitting a pull request, please ensure they are directed to the `dev`
branch of the repository.

Ensure your commit messages and pull request titles follow the
[Conventional Commits] specification.

## License

All files in this repository are licensed under the GNU Affero General Public
License v3.0 or later - see the [LICENSE] file for details.

<!-- links --->
[badge-url]: https://pkg.go.dev/badge/github.com/wordgen/wordgen.svg
[pkg-url]: https://pkg.go.dev/github.com/wordgen/wordgen
[wordlists package]: https://github.com/wordgen/wordlists
[command-line tool]: https://github.com/wordgen/cli
[Conventional Commits]: https://conventionalcommits.org
[LICENSE]: LICENSE
