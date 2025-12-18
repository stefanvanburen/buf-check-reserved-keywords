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
						FileName: "java.proto",
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
						FileName: "go.proto",
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
						FileName: "python.proto",
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
						FileName: "javascript.proto",
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
						FileName: "dart.proto",
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
						"enabled_languages": []string{"java", "go", "python", "javascript", "dart"},
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
