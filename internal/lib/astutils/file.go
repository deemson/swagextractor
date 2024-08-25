package astutils

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
)

func ParseFile(reader io.Reader) (*ast.File, error) {
	buffer := bytes.Buffer{}
	_, err := io.Copy(&buffer, reader)
	fileSet := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fileSet, "", buffer.String(), 0)
	if err != nil {
		return nil, err
	}
	return parsedFile, nil
}
