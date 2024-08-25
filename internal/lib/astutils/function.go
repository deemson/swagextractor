package astutils

import "go/ast"

func FindFunctionDeclaration(node ast.Node, functionName string) *ast.FuncDecl {
	var functionDeclaration *ast.FuncDecl
	ast.Inspect(node, func(childNode ast.Node) bool {
		if childNode == node {
			return true
		}
		funcDeclNode, ok := childNode.(*ast.FuncDecl)
		if ok {
			if funcDeclNode.Name.Name == functionName {
				functionDeclaration = funcDeclNode
			}
		}
		return false
	})
	return functionDeclaration
}

type FunctionReturnType struct {
	Package string
	Name    string
}

func DetermineFunctionReturnType(function *ast.FuncDecl) FunctionReturnType {
	if function == nil {
		panic("function is nil")
	}
	if function.Type == nil {
		panic("function.Type is nil")
	}
	if function.Type.Results == nil {
		panic("function.Type.Results is nil")
	}
	if function.Type.Results.List == nil {
		panic("function.Type.Results.List is nil")
	}
	if len(function.Type.Results.List) != 1 {
		panic("len(function.Type.Results.List) != 1")
	}
	field := function.Type.Results.List[0]
	if field == nil {
		panic("field is nil")
	}
	if field.Type == nil {
		panic("field.Type is nil")
	}
	starExpr := field.Type.(*ast.StarExpr)
	selectorExpr := starExpr.X.(*ast.SelectorExpr)
	ident := selectorExpr.X.(*ast.Ident)
	return FunctionReturnType{
		Package: ident.Name,
		Name:    selectorExpr.Sel.Name,
	}
}
