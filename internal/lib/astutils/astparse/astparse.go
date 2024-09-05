// Package astparse contains utility functions to parse go source files
package astparse

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"strings"
)

func Reader(reader io.Reader) (*ast.File, error) {
	buffer := bytes.Buffer{}
	_, err := io.Copy(&buffer, reader)
	fileSet := token.NewFileSet()
	parsedFile, err := parser.ParseFile(fileSet, "", buffer.String(), 0)
	if err != nil {
		return nil, err
	}
	return parsedFile, nil
}

func String(s string) (*ast.File, error) {
	return Reader(strings.NewReader(s))
}

func Lines(lines []string) (*ast.File, error) {
	return String(strings.Join(lines, "\n"))
}
