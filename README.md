[![Go Reference](https://pkg.go.dev/badge/github.com/xavier268/goviz.svg)](https://pkg.go.dev/github.com/xavier268/goviz)
[![Go Report Card](https://goreportcard.com/badge/github.com/xavier268/goviz)](https://goreportcard.com/report/github.com/xavier268/goviz)

# goviz

Create graphs of package dependencies, generating a graph file in the **graphviz** (.dot) format.

## Typical use


To analyse a project and all the packages it contains :

```
go run . -i "../path/to/my/project/" -f -e | dot -Tsvg | inkscape -p -g
```

## Getting help 

Use the -h option to get help :

```
go run . -h


Graphical dependency analysis for Golang packages
(c) 2022 Xavier Gandilot (aka xavier268)
Version     :   0.4.1
dot version :   dot - graphviz version 2.48.0 (0)
Typical use :   go run . -i "../path/to/my/project/" -f -e | dot -Tsvg | inkscape -p -g
  -V    Display version information and exit.
  -e    Show external packages.
  -f    Show program file names.
  -h    Show this help instructions and exit.
  -i string
        Top level input directory to analyse. (default "/home/xavier/Desktop/goviz")
  -t    Include test files (*_test.go).
  -v    Print verbose debugging information.
```

