package astfind_test

import (
	"github.com/deemson/swagextractor/internal/lib/astutils/astfind"
	"github.com/deemson/swagextractor/internal/lib/astutils/astparse"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeclStmt(t *testing.T) {
	testCases := map[string]struct {
		lines []string
		err   string
	}{
		"ok": {
			lines: []string{
				`package test`,
				`func f() {`,
				`	var v string`,
				`}`,
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			astFile, err := astparse.Lines(testCase.lines)
			require.NoError(t, err)
			astFuncDecl, err := astfind.FuncDecl(astFile, "f")
			require.NoError(t, err)
			//astfind.DeclStmt(astFuncDecl.Body, "v")
			require.NotNil(t, astFuncDecl)
			if testCase.err == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Equal(t, testCase.err, err.Error())
			}
		})
	}
}
