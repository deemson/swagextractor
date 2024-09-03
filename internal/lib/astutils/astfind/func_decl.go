package astfind

import "go/ast"

func FuncDecl(node ast.Node, funcName string) *ast.FuncDecl {
	var funcDecl *ast.FuncDecl
	ast.Inspect(node, func(childNode ast.Node) bool {
		if childNode == node {
			return true
		}
		funcDeclNode, ok := childNode.(*ast.FuncDecl)
		if ok {
			if funcDeclNode.Name.Name == funcName {
				funcDecl = funcDeclNode
			}
		}
		return false
	})
	return funcDecl
}
