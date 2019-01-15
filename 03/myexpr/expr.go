package main

import(
    "fmt"
    "os"
    "strconv"
)

func main(){
    a, err := strconv.Atoi(os.Args[1])
    
    if err != nil{
        fmt.Println(err)
    }
    
    b, err := strconv.Atoi(os.Args[2])

    if err != nil{
        fmt.Println(err)
    }

    fmt.Println(a,b)
}
