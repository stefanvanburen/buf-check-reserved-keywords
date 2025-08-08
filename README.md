# buf-check-reserved-keywords

This repository contains a [`buf` check plugin][buf-check-plugin] that checks for use of language reserved keywords in protobuf files.

It is published to the BSR at [svanburenorg/reserved-keywords](https://buf.build/svanburenorg/reserved-keywords).

## Why?

While it's considered best practice to [avoid using language reserved keywords for protobuf types][best-practice],
the default `buf` linter does not contain any checks to prevent their usage,
as this would require `buf` to know about the reserved keywords of various languages.

This plugin serves as a place to add these reserved keywords,
and prevent their usage in protobuf files.

## Supported Languages

* [Java][]
* [Go][]

[best-practice]: https://buf.build/docs/best-practices/style-guide/#recommendations
[buf-check-plugin]: https://github.com/bufbuild/bufplugin
[java]: https://docs.oracle.com/javase/tutorial/java/nutsandbolts/_keywords.html
[go]: https://go.dev/ref/spec#Keywords
