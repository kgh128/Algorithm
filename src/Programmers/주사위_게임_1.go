func solution(a int, b int) int {
    if a % 2 == 1 && b % 2 == 1 {
        return a * a + b * b
    } else if a % 2 == 0 && b % 2 == 0 {
        if (a > b) {
            return a - b
        } else {
            return b - a
        }
    } else {
        return 2 * (a + b)
    }
}
