package main

import "fmt"

func main(){
    var x int
    x = 1
    var y int
    y = 2
    fmt.Println(x,y)
    swap(&x,&y)
    fmt.Print(x,y)
}

func swap(p *int, q *int){
    var t int
    t = *p
    *p = *q
    *q = t
}
