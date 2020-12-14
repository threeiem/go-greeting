# Contributing

By participating to this project, you agree to abide our [code of
conduct](/CODE_OF_CONDUCT.md).

## Setup your machine

Written in [Go](https://golang.org/).

Prerequisites:

- `make`
- [Go 1.15+](https://golang.org/doc/install)
- [snapcraft](https://snapcraft.io/)
- [Docker](https://www.docker.com/)
- `gpg` (probably already installed on your system)

Clone `go-greeting` anywhere:

```sh
git clone git@github.com:threeiem/go-greeting.git
```

Install the build and lint dependencies:

```sh
make setup
```

Make sure everything is all right is running the test suite:

```sh
make test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

```sh
make build
```

Then run:

```sh
make ci
```

Runs all the linters and tests.

## Create a commit

Commit messages should be well formatted, and to make that "standardized", we are using Conventional Commits.

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Submit a pull request

Push your branch to your `go-greeting` fork and open a pull request against the main branch.
