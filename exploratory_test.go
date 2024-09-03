package swagextractor

import (
	"fmt"
	"github.com/deemson/swagextractor/internal/lib/astutils/astw"
	"github.com/deemson/swagextractor/internal/lib/openapi"
	"github.com/stretchr/testify/require"
	"go/ast"
	"strings"
	"testing"
)

func shallowFind[T ast.Node](node ast.Node) []T {
	var foundNodes []T
	ast.Inspect(node, func(childNode ast.Node) bool {
		if childNode == node {
			return true
		}
		concreteTypeNode, ok := childNode.(T)
		if ok {
			foundNodes = append(foundNodes, concreteTypeNode)
		}
		return false
	})
	return foundNodes
}

func getFuncDecl(node ast.Node, funcName string) *ast.FuncDecl {
	funcDecls := shallowFind[*ast.FuncDecl](node)
	for _, funcDecl := range funcDecls {
		if funcDecl.Name.Name == funcName {
			return funcDecl
		}
	}
	return nil
}

func f(node ast.Node) {
	ast.Inspect(node, func(childNode ast.Node) bool {
		if childNode == node {
			return true
		}
		assignStmt, ok := childNode.(*ast.AssignStmt)
		if ok {
			fmt.Printf("%#v\n", assignStmt)
		}
		return false
	})
}

func TestS(t *testing.T) {
	testCases := map[string]struct {
		source []string
		spec   openapi.Spec
	}{
		"asd": {
			source: []string{
				`package test`,
				``,
				`func f() {`,
				`	r := router()`,
				`	r.Methods("GET", "POST").Path("/path").Handle(nil)`,
				`	return r`,
				`}`,
			},
			spec: openapi.Spec{
				Paths: openapi.Paths{
					"/path": openapi.Path{
						Get:  openapi.Get{},
						Post: openapi.Post{},
					},
				},
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			astFile, err := astw.ParseReader(strings.NewReader(strings.Join(testCase.source, "\n")))
			require.NoError(t, err)
			astFuncDecl := getFuncDecl(astFile, "f")
			require.NotNil(t, astFuncDecl)
			ast.Inspect(astFuncDecl.Body, func(childNode ast.Node) bool {
				if childNode == astFuncDecl.Body {
					return true
				}
				fmt.Printf("%#v\n", childNode)
				return false
			})
		})
	}
}
