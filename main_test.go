package main

import (
	"testing"

	"buf.build/go/bufplugin/check/checktest"
)

func TestRule(t *testing.T) {
	t.Parallel()

	t.Run("invalid", func(t *testing.T) {
		request := newRequest(
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
		runCheckTest(t, request, want)
	})

	t.Run("valid", func(t *testing.T) {
		request := newRequest(
			"testdata/correct",
			[]string{"correct.proto"},
			nil,
		)
		runCheckTest(t, request)
	})
}

func runCheckTest(t *testing.T, request *checktest.RequestSpec, want ...checktest.ExpectedAnnotation) {
	checktest.CheckTest{
		Spec:                spec,
		Request:             request,
		ExpectedAnnotations: want,
	}.Run(t)
}

func newRequest(dir string, files []string, options map[string]any) *checktest.RequestSpec {
	return &checktest.RequestSpec{
		Files: &checktest.ProtoFileSpec{
			DirPaths:  []string{dir},
			FilePaths: files,
		},
		Options: options,
		// RuleIDs: []string{RuleID}, // The plugin is set to default=true
	}
}
