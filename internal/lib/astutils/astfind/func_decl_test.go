package astfind_test

import (
	"github.com/deemson/swagextractor/internal/lib/astutils/astfind"
	"github.com/deemson/swagextractor/internal/lib/astutils/astparse"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFuncDecl(t *testing.T) {
	testCases := map[string]struct {
		lines []string
		err   string
	}{
		"ok": {
			lines: []string{
				`package test`,
				`func f() {`,
				`}`,
			},
		},
		"declared 3 times": {
			lines: []string{
				`package test`,
				`func f() {`,
				`}`,
				`func f() {`,
				`}`,
				`func f() {`,
				`}`,
			},
			err: `function 'f' declared 3 times`,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			astFile, err := astparse.Lines(testCase.lines)
			require.NoError(t, err)
			astFuncDecl, err := astfind.FuncDecl(astFile, "f")
			if testCase.err == "" {
				require.NoError(t, err)
				require.NotNil(t, astFuncDecl)
			} else {
				require.Error(t, err)
				require.Equal(t, testCase.err, err.Error())
			}
		})
	}
}
