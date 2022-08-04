package main

import (
	"fmt"
	"go/ast"
)

// Dump output a text representation of thr pkgs content.
func Dump(pkgs map[string]*ast.Package) {

	for k, p := range pkgs {
		fmt.Println(k, p.Files, p.Imports, p.Name)
		for f, ff := range p.Files {
			fmt.Println("\t", f)
			fmt.Println("\t\t")
			for _, imp := range ff.Imports {
				fmt.Printf("\t\timport %v\n", imp.Path.Value)
			}
			fmt.Println("\t\t")
			for _, obj := range ff.Scope.Objects {
				fmt.Printf("\t\t%v %v\n", obj.Kind, obj.Name)
				switch o := obj.Decl.(type) {
				default:
					fmt.Printf("\t\t\t%T : %v\n", o, o)

				}
			}
			fmt.Println("\t\t")

		}
		fmt.Println()

	}

}
