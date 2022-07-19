package main

import (
	"fmt"
	"go/parser"
	"go/token"
)

func main() {

	dir := "../myprolog/pcontext"
	/*
		path, err := filepath.Abs(dir)
		if err != nil {
			panic(err)
		}
	*/

	fmt.Printf("Looking into %s\n", dir)

	fileset := token.NewFileSet()

	pkmap, err := parser.ParseDir(fileset, dir, nil, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(pkmap)
	for k, v := range pkmap {
		fmt.Printf("%v --> %#v\n", k, v)

	}

}
