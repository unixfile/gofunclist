package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	fset := token.NewFileSet()

	pkgs, err := parser.ParseDir(fset, ".", nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, pkg := range pkgs {
		ast.Inspect(pkg, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.FuncDecl:
				if x.Name.IsExported() {
					fmt.Println(funcSig(x))
				}
			}
			return true
		})
	}
}

func funcSig(f *ast.FuncDecl) string {
	receiver := ""
	if f.Recv != nil && len(f.Recv.List) > 0 {
		var recvNames []string
		for _, name := range f.Recv.List[0].Names {
			recvNames = append(recvNames, name.Name)
		}
		recvType := exprString(f.Recv.List[0].Type)
		receiver = fmt.Sprintf("(%s %s)", strings.Join(recvNames, ", "), recvType)
	}

	params := parameters(f.Type.Params)
	results := parameters(f.Type.Results)

	if receiver != "" {
		return fmt.Sprintf("func (%s) %s(%s) %s", receiver, f.Name.Name, params, results)
	} else {
		return fmt.Sprintf("func %s(%s) %s", f.Name.Name, params, results)
	}
}

func parameters(fieldList *ast.FieldList) string {
	if fieldList == nil {
		return ""
	}

	var params []string
	for _, field := range fieldList.List {
		var names []string
		for _, name := range field.Names {
			names = append(names, name.Name)
		}
		params = append(params, fmt.Sprintf("%s %s", strings.Join(names, ", "), exprString(field.Type)))
	}

	return strings.Join(params, ", ")
}

func exprString(expr ast.Expr) string {
	switch x := expr.(type) {
	case *ast.StarExpr:
		return "*" + exprString(x.X)
	case *ast.Ident:
		return x.Name
	case *ast.SelectorExpr:
		return exprString(x.X) + "." + x.Sel.Name
	case *ast.ArrayType:
		if x.Len == nil {
			return "[]" + exprString(x.Elt)
		} else {
			return "[" + exprString(x.Len) + "]" + exprString(x.Elt)
		}
	case *ast.Ellipsis:
		return "..." + exprString(x.Elt)
	default:
		return fmt.Sprintf("%T", x)
	}
}
