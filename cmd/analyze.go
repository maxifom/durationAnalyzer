package cmd

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

type visitor struct{}

func (v visitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	if n, ok := node.(*ast.BinaryExpr); ok {
		if !n.Op.IsOperator() {
			return v
		}
		if _, ok := n.Y.(*ast.BasicLit); !ok {
			return v
		}
		y := n.Y.(*ast.BasicLit)
		if y.Kind != token.INT {
			return v
		}
		if _, ok := n.X.(*ast.SelectorExpr); !ok {
			return v
		}
		x := n.X.(*ast.SelectorExpr)
		if _, ok := x.X.(*ast.Ident); !ok {
			return v
		}
		packageName := x.X.(*ast.Ident).Name
		name := x.Sel.Name
		if packageName == "time" && (name == "Nanosecond" || name == "Microsecond" || name == "Millisecond" || name == "Second" || name == "Minute" || name == "Hour") {
			i, _ := strconv.ParseInt(y.Value, 10, 64)
			fmt.Printf("Incorrect duration order: %s.%s %s %d.\nSuggested: %d %s %s.%s.\nPos: %d-%d\n", packageName, name, n.Op.String(), i, i, n.Op.String(), packageName, name, n.Pos(), n.End())
		}
	}
	return v
}

func analyze(filename string) error {
	fmt.Printf("File %s\n", filename)
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, filename, nil, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("failed to parse file %s: %s", filename, err)
	}
	var v visitor
	ast.Walk(v, file)
	return nil
}
