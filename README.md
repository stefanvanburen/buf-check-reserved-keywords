# buf-check-reserved-keywords

This repository contains a [`buf` check plugin][buf-check-plugin] that checks for use of language reserved keywords in protobuf files.

It is published to the BSR at [svanburenorg/reserved-keywords][bsr-module].

## Usage

You can find usage instructions on the BSR at [svanburenorg/reserved-keywords][bsr-module],
but to summarize:

1. Add the following to your `buf.yaml`'s `plugins:` stanza:

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


## Why?

While it's considered best practice to [avoid using language reserved keywords for protobuf types][best-practice],
the default `buf` linter does not contain any checks to prevent their usage,
as this would require `buf` to know about the reserved keywords of various languages.

This plugin serves as a place to add these reserved keywords,
and prevent their usage in protobuf files.

## Supported Languages

* [Java][]
* [Go][]
* [Python][]
* [JavaScript][]
* [Dart][]

[best-practice]: https://buf.build/docs/best-practices/style-guide/#recommendations
[bsr-module]: https://buf.build/svanburenorg/reserved-keywords
[buf-check-plugin]: https://github.com/bufbuild/bufplugin
[java]: https://docs.oracle.com/javase/tutorial/java/nutsandbolts/_keywords.html
[go]: https://go.dev/ref/spec#Keywords
[python]: https://docs.python.org/3/reference/lexical_analysis.html#keywords
[javascript]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#keywords
[dart]: https://dart.dev/language/keywords
