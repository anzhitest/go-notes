package main

import(
    "fmt"
    "strconv"
)

func main(){
    var n int
    var s string

    n = 10
    s = strconv.Itoa(n)
    m, err := strconv.Atoi("abc123")
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(n, s, m)
}
