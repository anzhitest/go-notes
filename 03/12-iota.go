package main

import "fmt"

const(
    A = iota * 10
    B
    C
)

func main(){
    fmt.Println(A, B, C)
}
