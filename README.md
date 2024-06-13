# wordgen

`wordgen` is a Go package that generates random words from a given wordlist.

## Installation

To install the package, use `go get`:

```
go get github.com/wordgen/wordgen
```

## Usage

Import the package in your Go code:

```go
import "github.com/wordgen/wordgen"
```

Use the `WordGenerator` function to generate random words:

```go
package main

import (
    "fmt"
    "github.com/wordgen/wordgen"
)

func main() {
    wordlist := []string{"apple", "banana", "cherry", "date"}
    result, err := wordgen.WordGenerator(wordlist, 5, "title", "-")
	
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	
    fmt.Println("Generated words:", result)
}
```

## WordGenerator Function

`WordGenerator` generates a specified number of random words from the given
wordlist, applies the specified case transformation ("upper", "title", "lower"),
and joins them with the specified separator. If the casing is an empty string,
it writes the word directly to the string builder.

**Parameters:**

- `list []string`: The wordlist to generate words from.
- `count int`: The number of words to generate.
- `casing string`: The case transformation to apply
  ("upper", "title", "lower", or empty for no transformation).
- `separator string`: The separator to use between words.

**Returns:**

- A string of generated words and nil, or an empty string and an error if the
  wordlist is empty or if random number generation fails.

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
[wordlists package]: https://github.com/wordgen/wordlists
[command-line tool]: https://github.com/wordgen/cli
[Conventional Commits]: https://conventionalcommits.org
[LICENSE]: LICENSE
