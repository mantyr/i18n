# i18n

[![GoDoc](https://godoc.org/github.com/mantyr/i18n?status.png)](http://godoc.org/github.com/mantyr/i18n)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

A template-based internationalization library for Go. Messages are stored per
language and compiled into `text/template`, so any message can itself contain
template actions and be rendered with data.

## Description

```
Plural
  └──Catalog          all languages
     └── Dictionary   one language
         └── key -> *template.Template

Catalog
  └── Dictionary
      └── key -> *template.Template
```

- **Catalog** holds one dictionary per language and is the main entry point.
- **Dictionary** holds the compiled templates for a single language.
- **Extractor** decides where the language sits in a nested key path.
- **Source** provides raw message data (e.g. a YAML file).
- **Extension** injects custom template functions into every message template.
- **Plural** determines the number type and declines it according to the language catalog.

## Author

[Oleg Shevelev][mantyr]

[mantyr]: https://github.com/mantyr
