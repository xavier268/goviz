package main

import (
	"fmt"
	"go/ast"
)

// Dump output a text representation of thr pkgs content.
func Dump(pkgs map[string]*ast.Package) {

	pps := AnalysePackages(pkgs)
	for p, pp := range pps {
		fmt.Print(p)
		if pp.Local {
			fmt.Println(" (local package)")
		} else {
			fmt.Println(" (external package)")
		}
		if len(pp.Imports) != 0 {
			fmt.Println("\tImports :\t")
			for x := range pp.Imports {
				fmt.Println("\t\t", x)
			}
		}
		if len(pp.Files) != 0 {
			fmt.Println("\tFiles   :\t")
			for x := range pp.Files {
				fmt.Println("\t\t", x)
			}
		}
	}

}
