[![Coverage Status](https://coveralls.io/repos/github/heat/latinify/badge.svg?branch=master)](https://coveralls.io/github/heat/latinify?branch=master)
# latinify

Library to help with Unicode conversion and romanization from other alphabets.

## Installation via Go Modules

Add the package to your `go.mod` with the command:

```bash
go get github.com/heat/latinify
```

Then import it in your code:

```go
import "github.com/heat/latinify"
```

## Usage

### Basic string conversion

Use `String` to normalize accents and other characters to the Latin alphabet:

```go
texto, err := latinify.String("Olá, 世界")
```

### Conversion with custom tables

`StringWithTable` allows applying additional mapping tables:

```go
table := latinify.Table{
    'ü': 'u',
    'ä': 'a',
}
texto, err := latinify.StringWithTable("für ärger", table)
```

### Generate slugs for URLs

`Slugify` creates a URL-friendly version of the given text:

```go
slug, err := latinify.Slugify("Olá, Mundo!")
// slug == "ola-mundo"
```

## Contribution guide

1. Fork the repository and create a branch for your changes.
2. Run `go test ./...` to make sure the tests still pass.
3. Open a Pull Request describing your changes.
