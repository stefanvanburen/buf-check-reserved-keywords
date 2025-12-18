package main

import (
	"context"
	_ "embed"
	"fmt"
	"slices"
	"strings"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
	"buf.build/go/bufplugin/descriptor"
	"buf.build/go/bufplugin/info"
	"buf.build/go/bufplugin/option"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	//go:embed README.md
	readmeMarkdown string
)

func main() {
	check.Main(spec)
}

const (
	ruleIDPackageNoLanguageReservedKeywords = "PLUGIN_PACKAGE_NO_LANGUAGE_RESERVED_KEYWORDS"
	ruleIDFieldNoLanguageReservedKeywords   = "PLUGIN_FIELD_NO_LANGUAGE_RESERVED_KEYWORDS"

	// enabledLanguagesOptionKey is the option key to override the default set of enabled
	// languages.
	// By default, all languages are checked.
	enabledLanguagesOptionKey = "enabled_languages"
)

var spec = &check.Spec{
	Rules: []*check.RuleSpec{
		{
			ID:      ruleIDPackageNoLanguageReservedKeywords,
			Default: true,
			Purpose: "Checks that all package names have no components that are language-reserved keywords.",
			Type:    check.RuleTypeLint,
			Handler: checkutil.NewFileRuleHandler(checkPackageNoLanguageReservedKeywords, checkutil.WithoutImports()),
		},
		{
			ID:      ruleIDFieldNoLanguageReservedKeywords,
			Default: true,
			Purpose: "Checks that all field names are not language-reserved keywords.",
			Type:    check.RuleTypeLint,
			Handler: checkutil.NewFieldRuleHandler(checkFieldNoLanguageReservedKeywords, checkutil.WithoutImports()),
		},
	},
	Info: &info.Spec{
		Documentation: readmeMarkdown,
		SPDXLicenseID: "apache-2.0",
		LicenseURL:    "https://github.com/stefanvanburen/buf-check-reserved-keywords/blob/main/LICENSE",
	},
}

func checkPackageNoLanguageReservedKeywords(
	_ context.Context,
	responseWriter check.ResponseWriter,
	request check.Request,
	fileDescriptor descriptor.FileDescriptor,
) error {
	validLanguages, err := getOptions(request)
	if err != nil {
		return fmt.Errorf("parsing options: %w", err)
	}
	packageName := fileDescriptor.FileDescriptorProto().Package
	if packageName == nil {
		return nil
	}
	packageComponents := strings.SplitSeq(*packageName, ".")
	for packageComponent := range packageComponents {
		for language, reservedKeywords := range languageReservedKeywords {
			if !slices.Contains(validLanguages, strings.ToLower(language)) {
				// Skip languages that aren't enabled.
				continue
			}
			if slices.Contains(reservedKeywords, packageComponent) {
				responseWriter.AddAnnotation(
					check.WithMessagef(
						"Package name %q should not use %s reserved keyword %q.",
						*packageName,
						language,
						packageComponent,
					),
					check.WithFileNameAndSourcePath(
						*fileDescriptor.FileDescriptorProto().Name,
						// A well-formed .proto file can only have a single `package` statement;
						// use that location.
						// https://github.com/protocolbuffers/protobuf/blob/6556a4ea26f2273797f559ebad87df42cd540443/src/google/protobuf/descriptor.proto#L109
						[]int32{2},
					),
				)
			}
		}
	}
	return nil
}

func checkFieldNoLanguageReservedKeywords(
	_ context.Context,
	responseWriter check.ResponseWriter,
	request check.Request,
	fieldDescriptor protoreflect.FieldDescriptor,
) error {
	validLanguages, err := getOptions(request)
	if err != nil {
		return fmt.Errorf("parsing options: %w", err)
	}
	for language, reservedKeywords := range languageReservedKeywords {
		if !slices.Contains(validLanguages, strings.ToLower(language)) {
			// Skip languages that aren't enabled.
			continue
		}
		fieldName := string(fieldDescriptor.Name())
		if slices.Contains(reservedKeywords, fieldName) {
			responseWriter.AddAnnotation(
				check.WithMessagef(
					"Field name %q should not use %s reserved keyword %q.",
					fieldName,
					language,
					fieldName,
				),
				check.WithDescriptor(fieldDescriptor),
			)
		}
	}
	return nil
}

func getOptions(request check.Request) (validLanguages []string, err error) {
	// Default to all languages being enabled.
	validLanguages = make([]string, 0, len(languageReservedKeywords))
	for language := range languageReservedKeywords {
		validLanguages = append(validLanguages, strings.ToLower(language))
	}
	enabledLanguagesOptionKey, err := option.GetStringSliceValue(request.Options(), enabledLanguagesOptionKey)
	if err != nil {
		return nil, err
	}
	if len(enabledLanguagesOptionKey) != 0 {
		for _, optionLanguage := range enabledLanguagesOptionKey {
			if !slices.Contains(validLanguages, optionLanguage) {
				return nil, fmt.Errorf("invalid language given %q, expected one of: %q", optionLanguage, strings.Join(validLanguages, ", "))
			}
		}
		// Use the specified languages instead.
		validLanguages = enabledLanguagesOptionKey
	}
	return validLanguages, nil
}

// TODO: Support more languages.
var (
	languageReservedKeywords = map[string][]string{
		// Ref: https://docs.oracle.com/javase/tutorial/java/nutsandbolts/_keywords.html
		"Java": {
			"abstract",
			"assert",
			"boolean",
			"break",
			"byte",
			"case",
			"catch",
			"char",
			"class",
			"const",
			"continue",
			"default",
			"do",
			"double",
			"else",
			"enum",
			"extends",
			"final",
			"finally",
			"float",
			"for",
			"goto",
			"if",
			"implements",
			"import",
			"instanceof",
			"int",
			"interface",
			"long",
			"native",
			"new",
			"package",
			"private",
			"protected",
			"public",
			"return",
			"short",
			"static",
			"strictfp",
			"super",
			"switch",
			"synchronized",
			"this",
			"throw",
			"throws",
			"transient",
			"try",
			"void",
			"volatile",
			"while",
		},
		// https://go.dev/ref/spec#Keywords
		"Go": {
			"break",
			"default",
			"func",
			"interface",
			"select",
			"case",
			"defer",
			"go",
			"map",
			"struct",
			"chan",
			"else",
			"goto",
			"package",
			"switch",
			"const",
			"fallthrough",
			"if",
			"range",
			"type",
			"continue",
			"for",
			"import",
			"return",
			"var",
		},
		"Python": {
			// https://docs.python.org/3/reference/lexical_analysis.html#keywords
			"False",
			"await",
			"else",
			"import",
			"pass",
			"None",
			"break",
			"except",
			"in",
			"raise",
			"True",
			"class",
			"finally",
			"is",
			"return",
			"and",
			"continue",
			"for",
			"lambda",
			"try",
			"as",
			"def",
			"from",
			"nonlocal",
			"while",
			"assert",
			"del",
			"global",
			"not",
			"with",
			"async",
			"elif",
			"if",
			"or",
			"yield",
			// https://docs.python.org/3/reference/lexical_analysis.html#soft-keywords
			"match",
			"case",
			"type",
			"_",
		},
		// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#keywords
		"JavaScript": {
			// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#reserved_words
			"break",
			"case",
			"catch",
			"class",
			"const",
			"continue",
			"debugger",
			"default",
			"delete",
			"do",
			"else",
			"export",
			"extends",
			"false",
			"finally",
			"for",
			"function",
			"if",
			"import",
			"in",
			"instanceof",
			"new",
			"null",
			"return",
			"super",
			"switch",
			"this",
			"throw",
			"true",
			"try",
			"typeof",
			"var",
			"void",
			"while",
			"with",

			"let",
			"static",
			"yield",

			"await",
			// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Lexical_grammar#future_reserved_words
			"enum",
			"implements",
			"interface",
			"package",
			"private",
			"protected",
			"public",
			"abstract",
			"boolean",
			"byte",
			"char",
			"double",
			"final",
			"float",
			"goto",
			"int",
			"long",
			"native",
			"short",
			"synchronized",
			"throws",
			"transient",
			"volatile",
		},
		// https://dart.dev/language/keywords
		"Dart": {
			"abstract",
			"as",
			"assert",
			"async",
			"await",
			"base",
			"break",
			"case",
			"catch",
			"class",
			"const",
			"continue",
			"covariant",
			"default",
			"deferred",
			"do",
			"dynamic",
			"else",
			"enum",
			"export",
			"extends",
			"extension",
			"external",
			"factory",
			"false",
			"final",
			"finally",
			"for",
			"Function",
			"get",
			"hide",
			"if",
			"implements",
			"import",
			"in",
			"interface",
			"is",
			"late",
			"library",
			"mixin",
			"new",
			"null",
			"of",
			"on",
			"operator",
			"part",
			"required",
			"rethrow",
			"return",
			"sealed",
			"set",
			"show",
			"static",
			"super",
			"switch",
			"sync",
			"this",
			"throw",
			"true",
			"try",
			"type",
			"typedef",
			"var",
			"void",
			"when",
			"with",
			"while",
			"yield",
		},
		// https://doc.rust-lang.org/reference/keywords.html
		"Rust": {
			// Strict keywords
			"_",
			"as",
			"async",
			"await",
			"break",
			"const",
			"continue",
			"crate",
			"dyn",
			"else",
			"enum",
			"extern",
			"false",
			"fn",
			"for",
			"if",
			"impl",
			"in",
			"let",
			"loop",
			"match",
			"mod",
			"move",
			"mut",
			"pub",
			"ref",
			"return",
			"self",
			"Self",
			"static",
			"struct",
			"super",
			"trait",
			"true",
			"type",
			"unsafe",
			"use",
			"where",
			"while",
			// Reserved keywords
			"abstract",
			"become",
			"box",
			"do",
			"final",
			"gen",
			"macro",
			"override",
			"priv",
			"try",
			"typeof",
			"unsized",
			"virtual",
			"yield",
		},
	}
)
