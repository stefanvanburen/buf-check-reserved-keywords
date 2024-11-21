package main

import (
	"context"
	"slices"
	"strings"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
	"buf.build/go/bufplugin/descriptor"
)

func main() {
	check.Main(spec)
}

const (
	ruleID = "PLUGIN_PACKAGE_NO_LANGUAGE_RESERVED_KEYWORDS"
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
	_ check.Request,
	fileDescriptor descriptor.FileDescriptor,
) error {
	packageName := fileDescriptor.FileDescriptorProto().Package
	if packageName == nil {
		return nil
	}
	packageComponents := strings.Split(*packageName, ".")
	for _, packageComponent := range packageComponents {
		for language, reservedKeywords := range languageReservedKeywords {
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
	}
)
