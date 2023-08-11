package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    if n%2 == 0 {
        fmt.Println(n, "is even")
    } else {
        fmt.Println(n, "is odd")
    }
}
