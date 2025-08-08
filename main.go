package main

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"strings"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
	"buf.build/go/bufplugin/descriptor"
	"buf.build/go/bufplugin/option"
)

func main() {
	check.Main(spec)
}

const (
	ruleID = "PLUGIN_PACKAGE_NO_LANGUAGE_RESERVED_KEYWORDS"

	// enabledLanguagesOptionKey is the option key to override the default set of enabled
	// languages.
	// By default, all languages are checked.
	enabledLanguagesOptionKey = "enabled_languages"
)

var spec = &check.Spec{
	Rules: []*check.RuleSpec{
		{
			ID:      ruleID,
			Default: true,
			Purpose: "Checks that all package names have no components that are language-reserved keywords.",
			Type:    check.RuleTypeLint,
			Handler: checkutil.NewFileRuleHandler(checkPackageNoLanguageReservedKeywords, checkutil.WithoutImports()),
		},
	},
}

func checkPackageNoLanguageReservedKeywords(
	_ context.Context,
	responseWriter check.ResponseWriter,
	request check.Request,
	fileDescriptor descriptor.FileDescriptor,
) error {
	// Default to all languages being enabled.
	validLanguages := make([]string, 0, len(languageReservedKeywords))
	for language := range maps.Keys(languageReservedKeywords) {
		validLanguages = append(validLanguages, strings.ToLower(language))
	}
	enabledLanguagesOptionKey, err := option.GetStringSliceValue(request.Options(), enabledLanguagesOptionKey)
	if err != nil {
		return err
	}
	if len(enabledLanguagesOptionKey) != 0 {
		for _, optionLanguage := range enabledLanguagesOptionKey {
			if !slices.Contains(validLanguages, optionLanguage) {
				return fmt.Errorf("invalid language given %q, expected one of: %q", optionLanguage, strings.Join(validLanguages, ", "))
			}
		}
		// Use the specified languages instead.
		validLanguages = enabledLanguagesOptionKey
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
					check.WithDescriptor(fileDescriptor.ProtoreflectFileDescriptor()),
				)
			}
		}
	}
	return nil
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
	}
)
