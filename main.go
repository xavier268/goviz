package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var (
	FlagHelp     bool   // Flag to show help and version info.
	FlagTest     bool   // Flag to include test files.
	FlagVerbose  bool   // Flag to print verbose output.
	FlagExternal bool   // Flag to show external packages or not
	FlagFiles    bool   // Flag to show files in packages
	FlagInputDir string // Define the top-level directory to analyse
	FlagVersion  bool   // print version information

	VERSION   = "0.4.1"
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
}

func welcome() {
	fmt.Fprintf(os.Stderr, "\nGraphical dependency analysis for Golang packages\n%s\nVersion     : \t%s\n",
		COPYRIGHT, VERSION)
	cmd := exec.Command("dot", "-V")
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Fprintln(os.Stderr, "The dot (graphviz) utilities do not seem to be available on your system ?")
	} else {
		fmt.Fprintf(os.Stderr, "dot version : \t%s", string(out))
	}
	fmt.Fprintln(os.Stderr, "Typical use : \tgo run . -i \"../path/to/my/project/\" -f -e | dot -Tsvg | inkscape -p -g")
}

func main() {

	flag.Parse()

	if FlagVersion {
		welcome()
		fmt.Fprintln(os.Stderr, "For help    : \tgo run . -h")
		return
	}

	if FlagHelp {
		welcome()
		flag.PrintDefaults()
		return
	}

	if FlagVerbose {
		fmt.Printf("\n// Analysing : %s\n", FlagInputDir)
	}

	_, pkgs, err := Parse(FlagInputDir)
	if err != nil {
		panic(err)
	}

	if FlagVerbose {
		Dump(pkgs)
	}

	DrawPackages(os.Stdout, AnalysePackages(pkgs))

}
