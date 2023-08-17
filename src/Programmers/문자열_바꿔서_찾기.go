import "strings"

func solution(str string, pat string) int {
    changed := []rune{}

    for _, char := range str {
        if char == 'A' {
            changed = append(changed, 'B')
        } else if char == 'B' {
            changed = append(changed, 'A')
        } else {
            changed = append(changed, char)
        }
    }

    if strings.Contains(string(changed), pat) {
        return 1
    } else {
        return 0
    }
}
