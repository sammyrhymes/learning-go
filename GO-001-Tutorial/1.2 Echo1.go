package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s = s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	sec := time.Since(start).Seconds()
	fmt.Printf("%.5f s\n\n", sec)
}
