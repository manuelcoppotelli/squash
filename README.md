squash
=======

[![pkg.go.dev](https://pkg.go.dev/badge/github.com/manuelcoppotelli/squash)](https://pkg.go.dev/github.com/manuelcoppotelli/squash)

[![Tests](https://github.com/manuelcoppotelli/squash/workflows/Tests/badge.svg)](https://github.com/manuelcoppotelli/squash/actions)

squash is a simple GoLang package to turn an array of arbitrarily nested arrays
into a flat array.


## Example

Input `[[1, 2, [3]], 4]` will become `[1 2 3 4]`

An example of the usage can be seen (and run) in the
[package official documentation](https://pkg.go.dev/github.com/manuelcoppotelli/squash#example-Integers)


## Tests

The package soruce code is tested automatically via
[Github Action](https://github.com/manuelcoppotelli/squash/actions/workflows/tests.yaml)
or can be tested manually typing the `go test` command
