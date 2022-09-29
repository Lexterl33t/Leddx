package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"reflect"
)

func main() {
	fs := token.NewFileSet()
	functionsNumber := map[string]int{}

	variableCountName := &ast.Ident{
		Name: "bryton",
	}

	p, err := parser.ParseFile(fs, "test/main.go", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	//adding loop and count instructions
	ast.Inspect(p, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.FuncDecl:

			if _, ok := functionsNumber[x.Name.String()]; !ok {
				for _, d := range x.Body.List {
					if reflect.TypeOf(d).String() != "*ast.ReturnStmt" {
						functionsNumber[x.Name.String()]++
					}
				}

				functionsNumber[fmt.Sprintf("%vLevel", x.Name.String())] = 0
			}

			assignSmt := &ast.AssignStmt{
				Lhs: []ast.Expr{
					variableCountName,
				},
				Rhs: []ast.Expr{
					&ast.BasicLit{
						Kind:  token.INT,
						Value: "0",
					},
				},
				Tok: token.DEFINE,
			}

			var clausesSmts []ast.Stmt
			block := &ast.BlockStmt{}
			for i, k := range x.Body.List {
				if reflect.TypeOf(k).String() == "*ast.ReturnStmt" {
					block.List = append(block.List, k)
					continue
				}
				clauseSmt := &ast.CaseClause{
					List: []ast.Expr{
						&ast.BasicLit{
							Kind:  token.INT,
							Value: fmt.Sprintf("%v", i),
						},
					},
					Body: []ast.Stmt{
						k,
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								variableCountName,
							},
							Rhs: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.INT,
									Value: fmt.Sprintf("%v", i+1),
								},
							},
							Tok: token.ASSIGN,
						},
					},
				}

				clausesSmts = append(clausesSmts, clauseSmt)
			}

			forSmt := &ast.ForStmt{
				Cond: &ast.BinaryExpr{
					Op: token.NEQ,
					X:  variableCountName,
					Y: &ast.BasicLit{
						Kind:  token.INT,
						Value: fmt.Sprintf("%v", functionsNumber[x.Name.String()]),
					},
				},

				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.SwitchStmt{
							Tag: variableCountName,
							Body: &ast.BlockStmt{
								List: clausesSmts,
							},
						},
					},
				},
			}

			block.List = append([]ast.Stmt{assignSmt, forSmt}, block.List...)
			x.Body = block

		}

		return true
	})

	printer.Fprint(os.Stdout, fs, p)
}
