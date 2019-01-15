package main

import "fmt"

func main(){
    var n int 
    var f float32
    n = 10
    //f = n / 3
    f = float32(n / 3)
    fmt.Println(n, f)
    n = int(f * 10)
    fmt.Println(n, f)
}
