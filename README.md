# 🚣 ghc - The git hook captain tool

[![Test & Build](https://github.com/zeiss/ghc/actions/workflows/main.yml/badge.svg)](https://github.com/zeiss/ghc/actions/workflows/main.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/zeiss/ghc)](https://goreportcard.com/report/github.com/zeiss/ghc)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

`ghc` is a teeny tiny tool to manage [git](https://git-scm.com) hooks via a configuration file.

### Homebrew

```bash
brew install zeiss/ghc-tap/ghc
```

### Go

```bash
go install github.com/zeiss/ghc@latest
```

## Specification

This is an example configuration file.

```yaml
# .ghc.yaml
version: 1
name: example
description: This is an example of a `ghc` configuration file
hooks:
  pre-commit:
    - make test
    - make lint
```

## License

[Apache 2.0](/LICENSE)
