package main

import (
	"flag"
	"fmt"
	"strings"
)

// 定义string类型的flag
// 命令行参数 默认值 说明
// 返回的是字符串指针 先声明个地址给你，用的时候使用*使用，所以必须用指针
var sep = flag.String("s", " ", "分隔符")

var line = flag.Bool("n", false, "是否换行")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if *line {
		fmt.Println()
	}
}
