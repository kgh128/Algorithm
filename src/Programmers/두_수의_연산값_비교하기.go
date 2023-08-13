import "strconv"

func solution(a int, b int) int {
    concat, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
    multiply := 2 * a * b

    if concat > multiply {
        return concat
    } else {
        return multiply
    }
}
