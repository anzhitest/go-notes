package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var filename string
	filename = os.Args[1]
	//    fmt.Printf(filename)
	printFile(filename)
}

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
