package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "removing whitespace")
var sep = flag.String("s", " ", "Separate strings using custom seperator")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

}
