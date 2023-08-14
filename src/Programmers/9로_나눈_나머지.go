func solution(number string) int {
    sum := 0

    for _, num := range number {
        sum += int(num - '0')
    }

    return sum % 9
}
