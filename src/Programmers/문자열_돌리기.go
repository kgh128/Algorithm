package main

import "fmt"

func main() {
    var str string
    fmt.Scan(&str)

    for _, char := range str {
        fmt.Println(string(char))
    }
}
