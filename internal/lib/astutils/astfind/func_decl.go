package astfind

import (
	"fmt"
	"go/ast"
)

func FuncDecl(node ast.Node, funcName string) (*ast.FuncDecl, error) {
	var funcDecls []*ast.FuncDecl
	ast.Inspect(node, func(childNode ast.Node) bool {
		if childNode == node {
			return true
		}
		funcDeclNode, ok := childNode.(*ast.FuncDecl)
		if ok {
			if funcDeclNode.Name.Name == funcName {
				funcDecls = append(funcDecls, funcDeclNode)
			}
		}
		return false
	})
	switch n := len(funcDecls); n {
	case 0:
		return nil, nil
	case 1:
		return funcDecls[0], nil
	default:
		return nil, fmt.Errorf(`function '%s' declared %d times`, funcName, n)
	}
}
