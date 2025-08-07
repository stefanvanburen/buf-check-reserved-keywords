package main

import (
	"strings"
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
		requestSpec := newRequestSpec(
			"testdata/java",
			[]string{"java.proto"},
			nil,
		)
		want := checktest.ExpectedAnnotation{
			RuleID:  ruleID,
			Message: `Package name "private" should not use Java reserved keyword "private".`,
			FileLocation: &checktest.ExpectedFileLocation{
				FileName: "java.proto",
			},
		}
		runCheckTest(t, requestSpec, want)
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
						"enabled_languages": []string{"go"},
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
				require.True(t, strings.HasPrefix(err.Error(), `Failed with code unknown: invalid language given "go", expected one of:`))
			})
			t.Run("valid", func(t *testing.T) {
				requestSpec := newRequestSpec(
					"testdata/correct",
					[]string{"correct.proto"},
					map[string]any{
						"enabled_languages": []string{"java"},
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
