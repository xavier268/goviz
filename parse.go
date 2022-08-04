package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Recursively parse provided dir and sub-dirs for go files.
func Parse(dir string) (fset *token.FileSet, pkgs map[string]*ast.Package, err error) {

	fset = token.NewFileSet()
	pkgs = make(map[string]*ast.Package)
	err = filepath.WalkDir(dir,
		func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				if d.Name()[0] == '.' {
					return filepath.SkipDir
				}
				pkmap, err := parser.ParseDir(fset, path, skipTests, 0)
				if err != nil {
					return err
				}
				// merge with previous information ...
				for k, v := range pkmap {
					pkgs[k] = v
				}
				return nil
			}
			return nil
		})

	return fset, pkgs, err
}

func skipTests(ff os.FileInfo) bool {
	return !strings.HasSuffix(ff.Name(), "_test.go")
}

// Analyse a parsed ast tree into a map of local packages.
func AnalysePackages(pkgs map[string]*ast.Package) map[string]Package {

	pps := make(map[string]Package)
	for pn, p := range pkgs {
		np := Package{
			Name:    pn,
			Local:   true,
			Imports: map[string]bool{},
			Files:   map[string]bool{},
		}

		// Add files & imports
		for f, ff := range p.Files {
			np.Files[f] = true
			for _, imp := range ff.Imports {
				base := filepath.Base(imp.Path.Value)
				base = strings.Trim(base, "\"")
				np.Imports[base] = true
			}
		}
		pps[pn] = np
	}

	// Add external imported packages to the global map.
	// The notion of "external" is relative to what directory/subdirectory was parsed.
	// Not assumption should be made on the package path.
	for _, p := range pps {
		for ip := range p.Imports {
			if _, found := pps[ip]; !found {
				// add to map an empty external package
				pps[ip] = Package{
					Name:    ip,
					Local:   false,
					Imports: map[string]bool{},
					Files:   map[string]bool{},
				}
			}
		}
	}
	return pps
}
