# buf-check-reserved-keywords

This repository contains a [`buf` check plugin][buf-check-plugin] that checks for use of language reserved keywords in protobuf files.

It is published to the BSR at [svanburenorg/reserved-keywords][bsr-module].

## Usage

You can find usage instructions on the BSR at [svanburenorg/reserved-keywords][bsr-module],
but to summarize:

1. Add the following to your [`buf.yaml`'s `plugins:` stanza][buf-yaml-plugins]:

```yaml
- plugin: buf.build/svanburenorg/reserved-keywords:main
```

2. Run the following command to download the plugin to your local environment:

```console
$ buf plugin update
```

3. If you have any `lint.use` rules specified
   (you probably have `DEFAULT` already added from the default `buf config init`),
   you'll need to explicitly add the lint rule to your `lint.use` stanza:

```diff
lint:
  use:
    - DEFAULT
+   - PLUGIN_PACKAGE_NO_LANGUAGE_RESERVED_KEYWORDS
```

## Options

`buf-check-reserved-keywords` currently supports a single option, `enabled_languages`.
If not specified, the plugin checks for keywords for all supported languages.
If specified, only the specified languages are checked.

For example, the following enables just checking for keywords for `go` and `python`.

```yaml
- plugin: buf.build/svanburenorg/reserved-keywords:main
  options:
    enabled_languages:
      - go
      - python
```


## Why?

While it's considered best practice to [avoid using language reserved keywords for protobuf types][best-practice],
the default `buf` linter does not contain any checks to prevent their usage,
as this would require `buf` to know about the reserved keywords of various languages.

This plugin serves as a place to add these reserved keywords,
and prevent their usage in protobuf files.

## Supported Languages

* [C][]
* [C++][]
* [Dart][]
* [Go][]
* [Java][]
* [JavaScript][]
* [Python][]
* [Rust][]

[best-practice]: https://buf.build/docs/best-practices/style-guide/#recommendations
[buf-yaml-plugins]: https://buf.build/docs/configuration/v2/buf-yaml/#plugins
[bsr-module]: https://buf.build/svanburenorg/reserved-keywords
[buf-check-plugin]: https://github.com/bufbuild/bufplugin
[c]: https://en.cppreference.com/w/c/keyword.html
[c++]: https://en.cppreference.com/w/cpp/keyword.html
[dart]: https://dart.dev/language/keywords
[go]: https://go.dev/ref/spec#Keywords
[java]: https://docs.oracle.com/javase/tutorial/java/nutsandbolts/_keywords.html
[javascript]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#keywords
[python]: https://docs.python.org/3/reference/lexical_analysis.html#keywords
[rust]: https://doc.rust-lang.org/reference/keywords.html
