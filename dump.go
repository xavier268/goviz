package main

import (
	"fmt"
	"go/ast"
	"os"
)

// Dump output a text representation of thr pkgs content.
func Dump(pkgs map[string]*ast.Package) {

	pps := AnalysePackages(pkgs)
	for p, pp := range pps {
		fmt.Fprint(os.Stderr, p)
		if pp.Local {
			fmt.Fprintln(os.Stderr, " (local package)")
		} else {
			fmt.Fprintln(os.Stderr, " (external package)")
		}
		if len(pp.Imports) != 0 {
			fmt.Fprintln(os.Stderr, "\tImports :\t")
			for x := range pp.Imports {
				fmt.Fprintln(os.Stderr, "\t\t", x)
			}
		}
		if len(pp.Files) != 0 {
			fmt.Fprintln(os.Stderr, "\tFiles   :\t")
			for x := range pp.Files {
				fmt.Fprintln(os.Stderr, "\t\t", x)
			}
		}
	}

}
