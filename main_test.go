package main

import (
	"testing"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checktest"
	"github.com/stretchr/testify/require"
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
					Message: `Field name "for" should not use Java reserved keyword "for".`,
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
					Message: `Package name "private.v1" should not use Java reserved keyword "private".`,
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
					RuleID:  ruleIDFieldNoLanguageReservedKeywords,
					Message: `Field name "for" should not use Go reserved keyword "for".`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "go.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "select.v1" should not use Go reserved keyword "select".`,
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
					Message: `Field name "for" should not use Python reserved keyword "for".`,
					FileLocation: &checktest.ExpectedFileLocation{
						FileName:    "python.proto",
						StartLine:   5,
						StartColumn: 2,
						EndLine:     5,
						EndColumn:   17,
					},
				},
				{
					RuleID:  ruleIDPackageNoLanguageReservedKeywords,
					Message: `Package name "continue.v1" should not use Python reserved keyword "continue".`,
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
					Message: `Field name "for" should not use JavaScript reserved keyword "for".`,
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
					Message: `Package name "typeof.v1" should not use JavaScript reserved keyword "typeof".`,
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
					Message: `Field name "for" should not use Dart reserved keyword "for".`,
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
					Message: `Package name "show.v1" should not use Dart reserved keyword "show".`,
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
					Message: `Field name "for" should not use Rust reserved keyword "for".`,
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
					Message: `Package name "trait.v1" should not use Rust reserved keyword "trait".`,
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
					Message: `Field name "for" should not use C reserved keyword "for".`,
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
					Message: `Package name "switch.v1" should not use C reserved keyword "switch".`,
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
					Message: `Field name "for" should not use C++ reserved keyword "for".`,
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
					Message: `Package name "class.v1" should not use C++ reserved keyword "class".`,
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
					Message: `Field name "for" should not use C# reserved keyword "for".`,
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
					Message: `Package name "namespace.v1" should not use C# reserved keyword "namespace".`,
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
					Message: `Field name "for" should not use Kotlin reserved keyword "for".`,
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
					Message: `Package name "object.v1" should not use Kotlin reserved keyword "object".`,
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
					Message: `Field name "for" should not use PHP reserved keyword "for".`,
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
					Message: `Package name "function.v1" should not use PHP reserved keyword "function".`,
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
					Message: `Field name "for" should not use Ruby reserved keyword "for".`,
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
					Message: `Package name "class.v1" should not use Ruby reserved keyword "class".`,
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
					Message: `Field name "for" should not use Scala reserved keyword "for".`,
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
					Message: `Package name "object.v1" should not use Scala reserved keyword "object".`,
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
					Message: `Field name "for" should not use Swift reserved keyword "for".`,
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
					Message: `Package name "protocol.v1" should not use Swift reserved keyword "protocol".`,
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
					Message: `Field name "for" should not use TypeScript reserved keyword "for".`,
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
					Message: `Package name "namespace.v1" should not use TypeScript reserved keyword "namespace".`,
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
				require.NoError(t, err)
				client, err := check.NewClientForSpec(spec)
				require.NoError(t, err)
				_, err = client.Check(ctx, request)
				require.Error(t, err)
				// Just check the prefix, so this doesn't fail as we add new supported
				// languages.
				require.ErrorContains(t, err, `Failed with code unknown: parsing options: invalid language given "invalid", expected one of:`)
			})
			t.Run("valid", func(t *testing.T) {
				requestSpec := newRequestSpec(
					"testdata/correct",
					[]string{"correct.proto"},
					map[string]any{
						"enabled_languages": []string{"c", "c++", "c#", "dart", "go", "java", "javascript", "kotlin", "php", "python", "ruby", "rust", "scala", "swift", "typescript"},
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
