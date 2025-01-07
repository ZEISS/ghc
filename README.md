# ghc - The Git Hooks Captain

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