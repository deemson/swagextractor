package openapi_test

import (
	"github.com/deemson/swagextractor/internal/lib/openapi"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	testCases := map[string]struct {
		path openapi.Path
		yml  []string
		err  string
	}{
		"empty": {path: openapi.Path{}, yml: []string{"{}"}},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			yml, err := yaml.Marshal(testCase.path)
			if testCase.err != "" {
				require.Error(t, err)
				require.Equal(t, testCase.err, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, strings.Join(testCase.yml, "\n")+"\n", string(yml))
			}
		})
	}
}

func TestPaths(t *testing.T) {
	testCases := map[string]struct {
		paths openapi.Paths
		yml   []string
		err   string
	}{
		"three empty paths": {
			paths: openapi.Paths{
				"/b": openapi.Path{},
				"/a": openapi.Path{},
				"/c": openapi.Path{},
			},
			yml: []string{
				`/a: {}`,
				`/b: {}`,
				`/c: {}`,
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			yml, err := yaml.Marshal(testCase.paths)
			if testCase.err != "" {
				require.Error(t, err)
				require.Equal(t, testCase.err, err.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, strings.Join(testCase.yml, "\n")+"\n", string(yml))
			}
		})
	}
}
