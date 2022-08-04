package main

import (
	"fmt"
	"os"
)

func main() {

	dir := "../myprolog"
	fmt.Printf("Looking into %s\n", dir)

	_, pkgs, err := Parse(dir)
	if err != nil {
		panic(err)
	}

	Dump(pkgs)
	of, err := os.Create("pack.dot")
	if err != nil {
		panic(err)
	}
	defer of.Close()

	DrawPackages(of, pkgs)

}
