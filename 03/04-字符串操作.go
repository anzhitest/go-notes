package main

import "fmt"

func main(){
    // 相加
    s1 := "hello" + " world"

    // 取字符
    // 字节 byte uint8
    var c1 byte
    c1 = s1[0]
    fmt.Printf("%d, %c \n", c1, c1)

    // 切片
    s2 := s1[0:3]

    fmt.Println(s1)
    fmt.Println(c1)
    fmt.Println(s2)
}
