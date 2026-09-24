package main

import (
	"testing"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checktest"
	"go.vanburen.xyz/ok"
)

func TestSpec(t *testing.T) {
	t.Parallel()
	checktest.SpecTest(t, spec)
}

func TestRule(t *testing.T) {
	t.Parallel()
	t.Run("invalid", func(t *testing.T) {
		t.Parallel()
		t.Run("java", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/java",
				[]string{"java.proto"},
				map[string]any{
					"enabled_languages": []string{"java"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Java.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "java.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "private.v1" contains "private", a reserved keyword in Java.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "java.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   19,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("go", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/go",
				[]string{"go.proto"},
				map[string]any{
					"enabled_languages": []string{"go"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "select.v1" contains "select", a reserved keyword in Go.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "go.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   18,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("python", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/python",
				[]string{"python.proto"},
				map[string]any{
					"enabled_languages": []string{"python"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Python.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "python.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "class" is a reserved keyword in Python.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "python.proto",
						StartLine:   6,
						StartColumn: 2,
						EndLine:     6,
						EndColumn:   32,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "continue.v1" contains "continue", a reserved keyword in Python.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "python.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   20,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("javascript", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/javascript",
				[]string{"javascript.proto"},
				map[string]any{
					"enabled_languages": []string{"javascript"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in JavaScript.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "javascript.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "typeof.v1" contains "typeof", a reserved keyword in JavaScript.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "javascript.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   18,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("dart", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/dart",
				[]string{"dart.proto"},
				map[string]any{
					"enabled_languages": []string{"dart"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Dart.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "dart.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "show.v1" contains "show", a reserved keyword in Dart.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "dart.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   16,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("rust", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/rust",
				[]string{"rust.proto"},
				map[string]any{
					"enabled_languages": []string{"rust"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Rust.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "rust.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "trait.v1" contains "trait", a reserved keyword in Rust.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "rust.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   17,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("c", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/c",
				[]string{"c.proto"},
				map[string]any{
					"enabled_languages": []string{"c"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in C.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "c.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "switch.v1" contains "switch", a reserved keyword in C.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "c.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   18,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("cpp", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/cpp",
				[]string{"cpp.proto"},
				map[string]any{
					"enabled_languages": []string{"c++"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in C++.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "cpp.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "class.v1" contains "class", a reserved keyword in C++.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "cpp.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   17,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("csharp", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/csharp",
				[]string{"csharp.proto"},
				map[string]any{
					"enabled_languages": []string{"c#"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in C#.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "csharp.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "namespace.v1" contains "namespace", a reserved keyword in C#.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "csharp.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   21,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("kotlin", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/kotlin",
				[]string{"kotlin.proto"},
				map[string]any{
					"enabled_languages": []string{"kotlin"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Kotlin.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "kotlin.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "object.v1" contains "object", a reserved keyword in Kotlin.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "kotlin.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   18,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("php", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/php",
				[]string{"php.proto"},
				map[string]any{
					"enabled_languages": []string{"php"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in PHP.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "php.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "function.v1" contains "function", a reserved keyword in PHP.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "php.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   20,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("ruby", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/ruby",
				[]string{"ruby.proto"},
				map[string]any{
					"enabled_languages": []string{"ruby"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Ruby.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "ruby.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "class.v1" contains "class", a reserved keyword in Ruby.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "ruby.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   17,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("scala", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/scala",
				[]string{"scala.proto"},
				map[string]any{
					"enabled_languages": []string{"scala"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Scala.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "scala.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "object.v1" contains "object", a reserved keyword in Scala.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "scala.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   18,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("swift", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/swift",
				[]string{"swift.proto"},
				map[string]any{
					"enabled_languages": []string{"swift"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Swift.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "swift.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "protocol.v1" contains "protocol", a reserved keyword in Swift.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "swift.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   20,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("typescript", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/typescript",
				[]string{"typescript.proto"},
				map[string]any{
					"enabled_languages": []string{"typescript"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in TypeScript.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "typescript.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "namespace.v1" contains "namespace", a reserved keyword in TypeScript.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "typescript.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   21,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("objc", func(t *testing.T) {
			t.Parallel()
			requestSpec := newRequestSpec(
				"testdata/objc",
				[]string{"objc.proto"},
				map[string]any{
					"enabled_languages": []string{"objective-c"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Objective-C.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "objc.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "protocol.v1" contains "protocol", a reserved keyword in Objective-C.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "objc.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   20,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
		t.Run("multiple", func(t *testing.T) {
			t.Parallel()
			// A field or package component that is a keyword in several languages
			// gets one annotation listing them in sorted order, whatever the option
			// order.
			requestSpec := newRequestSpec(
				"testdata/multiple",
				[]string{"multiple.proto"},
				map[string]any{
					"enabled_languages": []string{"rust", "python", "java"},
				},
			)
			want := []checktest.ExpectedAnnotation{
				{
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" is a reserved keyword in Java, Python, Rust.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "multiple.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "for.v1" contains "for", a reserved keyword in Java, Python, Rust.`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "multiple.proto",
						StartLine:   2,
						StartColumn: 0,
						EndLine:     2,
						EndColumn:   15,
					},
				},
			}
			runCheckTest(t, requestSpec, want...)
		})
	})
	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		requestSpec := newRequestSpec(
			"testdata/correct",
			[]string{"correct.proto"},
			nil,
		)
		runCheckTest(t, requestSpec)
	})
	t.Run("options", func(t *testing.T) {
		t.Parallel()
		t.Run("enabled_languages", func(t *testing.T) {
			t.Run("invalid", func(t *testing.T) {
				requestSpec := newRequestSpec(
					"testdata/correct",
					[]string{"correct.proto"},
					map[string]any{
						"enabled_languages": []string{"invalid"},
					},
				)

				ctx := t.Context()
				request, err := requestSpec.ToRequest(ctx)
				ok.MustNoError(t, err)
				client, err := check.NewClientForSpec(spec)
				ok.MustNoError(t, err)
				_, err = client.Check(ctx, request)
				// Just check the prefix, so this doesn't fail as we add new supported
				// languages. ErrorContains reports a nil error itself, so the
				// separate "expected error" check would only have been a way to
				// panic on err.Error() below.
				const want = `Failed with code unknown: parsing options: invalid language given "invalid", expected one of:`
				ok.ErrorContains(t, err, want)
			})
			t.Run("valid", func(t *testing.T) {
				requestSpec := newRequestSpec(
					"testdata/correct",
					[]string{"correct.proto"},
					map[string]any{
						"enabled_languages": []string{"c", "c++", "c#", "dart", "go", "java", "javascript", "kotlin", "objective-c", "php", "python", "ruby", "rust", "scala", "swift", "typescript"},
					},
				)

				runCheckTest(t, requestSpec)
			})
		})
	})
}

func runCheckTest(t *testing.T, request *checktest.RequestSpec, want ...checktest.ExpectedAnnotation) {
	checktest.CheckTest{
		Spec:                spec,
		Request:             request,
		ExpectedAnnotations: want,
	}.Run(t)
}

func newRequestSpec(dir string, files []string, options map[string]any) *checktest.RequestSpec {
	return &checktest.RequestSpec{
		Files: &checktest.ProtoFileSpec{
			DirPaths:  []string{dir},
			FilePaths: files,
		},
		Options: options,
		// RuleIDs: []string{RuleID}, // The plugin is set to default=true
	}
}
