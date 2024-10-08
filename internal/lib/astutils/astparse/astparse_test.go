package astparse_test

import (
	"github.com/deemson/swagextractor/internal/lib/astutils/astparse"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReader(t *testing.T) {
	testCases := map[string]struct {
		lines []string
		err   string
	}{
		"empty": {
			lines: []string{},
			err:   "1:1: expected 'package', found 'EOF'",
		},
		"bad package": {
			lines: []string{
				`asdf`,
			},
			err: "1:1: expected 'package', found asdf",
		},
		"just package": {
			lines: []string{
				`package test`,
			},
		},
		"bad content": {
			lines: []string{
				`package test`,
				`asdf`,
			},
			err: "2:1: expected declaration, found asdf",
		},
		"ok empty func": {
			lines: []string{
				`package test`,
				`func f() {`,
				`}`,
			},
		},
		"func no closing bracket": {
			lines: []string{
				`package test`,
				`func f() {`,
			},
			err: "2:11: expected '}', found 'EOF'",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			_, err := astparse.Lines(testCase.lines)
			if testCase.err == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Equal(t, testCase.err, err.Error())
			}
		})
	}
}
