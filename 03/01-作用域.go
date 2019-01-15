package main

import "fmt"

var x = 200

func localFunc(){
    fmt.Println(&x)
    fmt.Println(x)
}

func main(){
    x := 1

    localFunc()
    fmt.Println(&x)
    fmt.Println(x)

    if true{
        x := 100
        fmt.Println(&x)
        fmt.Println(x)
    }

    localFunc()
    fmt.Println(&x)
    fmt.Println(x)
}
