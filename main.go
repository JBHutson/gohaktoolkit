package main

import (
	"flag"

	"github.com/JBHutson/gohaktoolkit/portscanner"
)

var ps = flag.Bool("p", false, "run a port scan on a single target")

func main() {
	flag.Parse()
	if *ps {
		portscanner.Scan(flag.Arg(0))
	}
}
