# goviz

Create graphs of package dependencies, generating a graph file in the **graphviz** (.dot) format.

## Typical use

```
go run . -i "../path/to/a/directory/" -f -e && dot -Tsvg out.dot > out.svg && firefox out.svg
```

## Getting help 

Use the -h option to get help :

```
go run . -h

Graphical dependency analysis for Golang packages
(c) 2022 Xavier Gandilot (aka xavier268)
Version     :   0.3.4
dot version :   dot - graphviz version 2.48.0 (0)
Typical use :   go run . -e -f  && dot -Tsvg out.dot > out.svg && firefox out.svg
  -V    Display version information and exit.
  -e    Show external packages.
  -f    Show program file names.
  -h    Show this help instructions and exit.
  -i string
        Top level input directory to analyse. (default to current working directory)
  -o string
        Output file in .dot (graphviz) format. (default "out.dot")
  -v    Print verbose debugging information.
```

