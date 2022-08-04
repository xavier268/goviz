package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	FlagExternal bool // Flag to show external packages or not
	FlagFiles    bool // Flag to show files in packages
)

func init() {
	flag.BoolVar(&FlagExternal, "e", false, "Show external packages.")
	flag.BoolVar(&FlagFiles, "f", false, "Show program file names.")
}

func main() {

	flag.Parse()

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

	DrawPackages(of, AnalysePackages(pkgs))

}
