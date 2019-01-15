package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "分隔符")
var line = flag.Bool("n", false, "是否换行")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *line {
		fmt.Println()
	}
}
