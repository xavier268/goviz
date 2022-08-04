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
