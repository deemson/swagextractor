package swagextractor

import (
	"fmt"
	"github.com/deemson/swagextractor/internal/lib/astutils"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestS(t *testing.T) {
	source := []string{
		`package test`,
		`func f() *mux.Router {`,
		`	r := mux.NewRouter()`,
		`	r.Methods("GET").Path("/path").Handler(nil)`,
		`	return r`,
		`}`,
	}
	astFile, err := astutils.ParseFile(strings.NewReader(strings.Join(source, "\n")))
	require.NoError(t, err)
	fmt.Println(astutils.FindFunctionDeclaration(astFile, "f"))
}
