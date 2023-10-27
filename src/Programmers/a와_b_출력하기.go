package main

import "fmt"

func main() {
    var a, b int
    fmt.Scan(&a, &b)

    // '=' 다음에 자동으로 띄어쓰기가 들어감.
    // 예시 출력: a = 3
    fmt.Println("a =", a)
    fmt.Println("b =", b)

    // 따라서 아래와 같이 코드를 쓰면 띄어쓰기가 두 번 들어감.
    // 예시 출력: a =  3
    // fmt.Println("a = ", a)
    // fmt.Println("b = ", b)
}
