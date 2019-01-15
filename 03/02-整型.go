package main

import "fmt"
import "unsafe"

func main(){
    var (
        x0 int8
        x1 int
        x2 int32
        x3 int64
        x4 uint
        x5 uint32
        x6 uint64
    )
    fmt.Println(x1, x2, x3, x4, x5, x6)
    fmt.Println(unsafe.Sizeof(x1), unsafe.Sizeof(x2), unsafe.Sizeof(x3), unsafe.Sizeof(x4), unsafe.Sizeof(x5), unsafe.Sizeof(x6))
    x0 = 127
    x0 = x0 +1
    fmt.Println(x0)
}
