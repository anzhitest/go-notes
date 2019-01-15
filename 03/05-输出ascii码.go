package main

import "fmt"

func main(){
    var b byte
    for b=0;b<=177;b++{
        fmt.Printf("%d,%c\n", b, b)
    }
    
    // 输出16进制
    fmt.Println(0Xa)
}
