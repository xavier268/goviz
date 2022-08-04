package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	FlagHelp       bool   // Flag to show help and version info.
	FlagTest       bool   // Flag to include test files.
	FlagVerbose    bool   // Flag to print verbose output.
	FlagExternal   bool   // Flag to show external packages or not
	FlagFiles      bool   // Flag to show files in packages
	FlagInputDir   string // Define the top-level directory to analyse
	FlagOutputFile string // Define the output file
	FlagVersion    bool   // print version information

	VERSION   = "0.3.5"
	COPYRIGHT = "(c) 2022 Xavier Gandilot (aka xavier268)"
)

func init() {
	wd, _ := os.Getwd()
	flag.BoolVar(&FlagVerbose, "v", false, "Print verbose debugging information.")
	flag.BoolVar(&FlagTest, "t", false, "Include test files (*_test.go).")
	flag.BoolVar(&FlagExternal, "e", false, "Show external packages.")
	flag.BoolVar(&FlagFiles, "f", false, "Show program file names.")
	flag.BoolVar(&FlagHelp, "h", false, "Show this help instructions and exit.")
	flag.BoolVar(&FlagVersion, "V", false, "Display version information and exit.")
	flag.StringVar(&FlagInputDir, "i", wd, "Top level input directory to analyse.")
	flag.StringVar(&FlagOutputFile, "o", "out.dot", "Output file in .dot (graphviz) format.")
}

func welcome() {
	fmt.Printf("\nGraphical dependency analysis for Golang packages\n%s\nVersion     : \t%s\n",
		COPYRIGHT, VERSION)
	cmd := exec.Command("dot", "-V")
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Println("The dot (graphviz) utilities do not seem to be available on your system ?")
	} else {
		fmt.Printf("dot version : \t%s", string(out))
	}
	fmt.Println("Typical use : \tgo run . -e -f -o a.dot && dot -Tsvg a.dot > a.svg && firefox a.svg")
}

func main() {

	flag.Parse()

	if FlagVersion {
		welcome()
		fmt.Println("For help    : \tgo run . -h")
		return
	}

	if FlagHelp {
		welcome()
		flag.PrintDefaults()
		return
	}

	if FlagVerbose {
		fmt.Printf("\nAnalysing : %s\n", FlagInputDir)
	}

	_, pkgs, err := Parse(FlagInputDir)
	if err != nil {
		panic(err)
	}

	if FlagVerbose {
		Dump(pkgs)
	}

	of, err := os.Create(FlagOutputFile)
	if err != nil {
		panic(err)
	}
	defer of.Close()

	DrawPackages(of, AnalysePackages(pkgs))

}
