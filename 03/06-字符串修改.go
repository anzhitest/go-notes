package main

import "fmt"

func main(){
    // go里面的字符串本身不可以修改，只能通过[]byte互相转换修改
    s1 := "hello world"

    //s1[0] = 'H'

    arr := []byte(s1)
    fmt.Println(arr)
    arr[0] = 'H'
    s1 = string(arr)
    fmt.Println(s1)
}
