func solution(a int, b int, c int) int {
    if a != b && b != c && c != a {
        return a + b + c
    } else if a == b && b == c {
        return (a + b + c) * (a*a + b*b + c*c) * (a*a*a + b*b*b + c*c*c)
    } else {
        return (a + b + c) * (a*a + b*b + c*c)
    }
}
