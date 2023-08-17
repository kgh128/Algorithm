import (
    "strings"
    "strconv"
)

func solution(binomial string) int {
    parseArr := strings.Fields(binomial)

    op := parseArr[1]
    a, _ := strconv.Atoi(parseArr[0])
    b, _ := strconv.Atoi(parseArr[2])

    switch op {
        case "+":
            return a + b
        case "-":
            return a - b
        case "*":
            return a * b
    }

    return 0
}
