package astfind

import (
	"fmt"
	"go/ast"
)

func DeclStmt(node ast.Node, varName string) {
	ast.Inspect(node, func(childNode ast.Node) bool {
		if childNode == node {
			return true
		}
		declStmtNode, ok := childNode.(*ast.DeclStmt)
		if ok {
			fmt.Printf("%#v\n", declStmtNode.Decl)
		}
		return false
	})
}
