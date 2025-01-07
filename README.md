# ghc - The Git Hooks Captain

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