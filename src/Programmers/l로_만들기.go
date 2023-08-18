func solution(myString string) string {
    result := []rune{}

    for _, char := range myString {
        if char < 'l' {
            result = append(result, 'l')
        } else {
            result = append(result, char)
        }
    }

    return string(result)
}
