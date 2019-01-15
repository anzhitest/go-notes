package main

import "fmt"

func main(){
    //fmt.Printf("%c\n",'a'+('H'-'h'))
    var s string
    s = "hello world"
    s = toupper(s)
    fmt.Println(s)
}

func toupper(s string) string{
    arr := []byte(s)
    for i:=0;i<len(s);i++{
        arr[i] = arr[i] - 32
    }
    s = string(arr)
    return s
}
